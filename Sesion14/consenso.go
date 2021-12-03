package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

const localAddr = "localhost:8003" // su propia IP aquí
const (
	cnum = iota // iota genera valores en secuencia y se reinicia en cada bloque const
	opa
	opb
)

type Info struct {
	Code int
	Addr string
	Op   int
}

// Las IP de los demás participantes acá, todos deberían usar el puerto 8000
var addrs = []string{"localhost:8000",
	"localhost:8001",
	"localhost:8002"}

var chInfo chan map[string]int

func main() {
	chInfo = make(chan map[string]int)
	go func() {
		chInfo <- map[string]int{} // mapa vacío
	}()
	go server() // habilitar el servidor (modo escucha)
	time.Sleep(time.Millisecond * 100)
	var op int
	for {
		fmt.Print("Your option: ")
		fmt.Scanf("%d\n", &op)
		msg := Info{cnum, localAddr, op}
		for _, addr := range addrs {
			send(addr, msg)
		}
	}
}

func server() {
	if ln, err := net.Listen("tcp", localAddr); err != nil {
		log.Panicln("Can't start listener on", localAddr)
	} else {
		defer ln.Close()
		fmt.Println("Listeing on", localAddr)
		for {
			if conn, err := ln.Accept(); err != nil {
				log.Println("Can't accept", conn.RemoteAddr())
			} else {
				go handle(conn)
			}
		}
	}
}

func handle(conn net.Conn) { // lógica
	defer conn.Close()
	dec := json.NewDecoder(conn)
	var msg Info
	if err := dec.Decode(&msg); err != nil {
		log.Println("Can't decode from", conn.RemoteAddr())
	} else {
		fmt.Println(msg)
		switch msg.Code {
		case cnum:
			concensus(conn, msg)
		}
	}
}

func concensus(conn net.Conn, msg Info) {
	// sincronizar
	info := <-chInfo
	// guardamos en el mapa las votaciones de cada nodo
	info[msg.Addr] = msg.Op
	// validar si ya recibimos las votaciones de todos los nodos
	if len(info) == len(addrs) { // si aún no votan todos, espera a que voten
		// aplica la lógica
		ca, cb := 0, 0
		for _, op := range info {
			if op == opa {
				ca++
			} else {
				cb++
			}
		}
		if ca > cb {
			fmt.Println("Ganó A!")
		} else {
			fmt.Println("Ganó B!")
		}
		// inicializamos la info (cero votos)
		info = map[string]int{}
	}
	// enviar para una prox sincronización
	go func() {
		chInfo <- info
	}()
}

func send(remoteAddr string, msg Info) {
	if conn, err := net.Dial("tcp", remoteAddr); err != nil {
		log.Println("Can't dial", remoteAddr)
	} else {
		defer conn.Close()
		fmt.Println("Sending to", remoteAddr)
		enc := json.NewEncoder(conn)
		enc.Encode(msg)
	}
}

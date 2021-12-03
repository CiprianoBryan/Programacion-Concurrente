package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

var addrs []string

type Info struct {
	Codigo int
	Addr   string
	Opcion int
}

const (
	codNum = iota // genera valores secuenciales
	OpA
	OpB
)

var direccion string
var chanInfo chan map[string]int

func main() {
	chanInfo = make(chan map[string]int)

	var n int

	fmt.Print("Ingrese la dirección del nodo: ")
	fmt.Scanf("%s\n", &direccion)

	// 1. Solicitar las direcciones de los demas nodos
	fmt.Print("Ingrese la cantidad de nodos: ")
	fmt.Scanf("%d\n", &n)

	// guardar en la bitacora
	addrs = make([]string, n)
	for i := range addrs {
		fmt.Printf("Host %d = ", i+1)
		fmt.Scanf("%s\n", &(addrs[i]))
	}

	go func() {
		chanInfo <- map[string]int{}
	}()

	go habilitarServicio() // modo servidor

	// 	Envio

}

func habilitarServicio() {
	listen, _ := net.Listen("tcp", direccion)
	defer listen.Close()
	for {
		con, _ := listen.Accept()
		go manejadorServicio(con)
	}
}

func manejadorServicio(con net.Conn) {
	// logica
	defer con.Close()
	bufferIn := bufio.NewReader(con)
	bInfo, _ := bufferIn.ReadString('\n')

	var info Info
	json.Unmarshal([]byte(bInfo), &info)

	fmt.Println(info)

	switch info.Codigo {
	case codNum:
		procesarConsenso(con, info)
	}
}

func procesarConsenso(con net.Conn, info Info) {
	// sincronizar
	mapInfo := <-chanInfo
	mapInfo[info.Addr] = info.Opcion

	if len(mapInfo) == len(addrs) {
		valA, valB := 0, 0
		for _, op := range mapInfo {
			if op == OpA {
				valA++
			} else {
				valB++
			}
		}
	}
}

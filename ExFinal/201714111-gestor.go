package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
)

type Info struct {
	Tipo       string
	Nombre     string
	Encontrado bool
}

type InfoServer struct {
	Letra string
}

var info InfoServer
var bitacoraAddr = []string{
	"localhost:9101",
	"localhost:9102",
	"localhost:9103"}

const AddrServer = "localhost:8101"

var nodosClasificaron []string

func main() {

	for {
		go escucha()
		////////////////// PAUSA //////////////////
		fmt.Println("Presione enter para iniciar: ")
		bufferIn := bufio.NewReader(os.Stdin)
		bufferIn.ReadString('\n')
		//////////////////////////////////////////

		// De forma aleatoria se va a seleccionar un nodo de la red
		remotehost := bitacoraAddr[rand.Intn(len(bitacoraAddr))]
		// De forma aleatoria se va a seleccionar una letra del abecedario
		abc := "abcdefghijklmnoprstuvwxyz"
		info.Letra = string(abc[rand.Intn(len(abc))])

		///////// Se envía la letra a un nodo /////////
		fmt.Printf("Enviando la letra: %s, al nodo: %s\n", info.Letra, remotehost)

		con, _ := net.Dial("tcp", remotehost)
		defer con.Close()
		jsonBytes, _ := json.Marshal(info)
		fmt.Fprintln(con, string(jsonBytes))
		/////////////////////////////////////////////
	}
}

func escucha() {
	listen, _ := net.Listen("tcp", AddrServer)
	defer listen.Close()

	for {
		con, _ := listen.Accept()
		go atenderNotificacion(con)
	}
}

func atenderNotificacion(con net.Conn) {
	///////// Se recibe la respuesta del nodo /////////
	bufferIn := bufio.NewReader(con)
	nodoJsonBytes, _ := bufferIn.ReadString('\n')

	var infoNodo Info
	json.Unmarshal([]byte(nodoJsonBytes), &infoNodo)

	if infoNodo.Encontrado {
		fmt.Printf("Se encontró el nombre: %s\n", infoNodo.Nombre)
		nodosClasificaron = append(nodosClasificaron)
	}
}

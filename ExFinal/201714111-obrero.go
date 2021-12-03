package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

var direccion string // localhost:9101 / localhost:9102 / localhost:9103

const AddrServer = "localhost:8101"

var Nombres = []string{
	"JUAN",
	"INES",
	"CARMEN"}

// estructura que se comparte entre los nodos
type Info struct {
	Tipo       string
	Nombre     string
	Encontrado bool
}

type InfoServer struct {
	Letra string
}

var chIniciar chan bool // esperar que el nodo inicie su trabajo en Sección critica
var chInfo chan Info

func main() {
	fmt.Print("Ingrese la direccion del nodo: ")
	fmt.Scanf("%s\n", &direccion)

	chIniciar = make(chan bool)

	go func() {
		////////////////// PAUSA //////////////////
		//fmt.Println("Presione enter para iniciar: ")
		//bufferIn := bufio.NewReader(os.Stdin)
		//bufferIn.ReadString('\n')
		//////////////////////////////////////////

		// crear la info a enviar
		info := Info{"ENVIO", "", false}
		// notificar a todos los nodos de la bitácora
		go enviar(AddrServer, info)
	}()

	// 4.- Publicar el servicio / ROL SERVIDOR
	serviciosSC()
}

func enviar(addr string, info Info) {
	con, _ := net.Dial("tcp", addr)
	defer con.Close()
	// codificar el mensaje a enviar al servidor
	binfo, _ := json.Marshal(info)
	fmt.Fprintln(con, string(binfo))
}

func serviciosSC() {
	// expone el puerto y se coloca en modo escucha
	listen, _ := net.Listen("tcp", direccion)
	defer listen.Close()
	for {
		con, _ := listen.Accept()
		go manejadorConexiones(con)
	}
}

func manejadorConexiones(con net.Conn) {
	// logica del servicio
	defer con.Close()
	bufferIn := bufio.NewReader(con)
	bInfo, _ := bufferIn.ReadString('\n')
	// decodificar
	var infoServer InfoServer // Obtención de la letra generada por el servidor
	json.Unmarshal([]byte(bInfo), &infoServer)
	fmt.Printf("Se obtuvo del servidor la letra: %s\n", infoServer.Letra)

	//recuperar del canal la info del nodo
	myInfo := <-chInfo
	nombreParcial := myInfo.Nombre + infoServer.Letra

	for _, nombre := range Nombres {
		if len(nombreParcial) <= len(nombre) && nombreParcial == nombre[:len(nombreParcial)] {
			myInfo.Nombre = nombreParcial
			if nombreParcial == nombre {
				fmt.Printf("Se clasificó el nombre: %s\n", nombreParcial)
				myInfo.Encontrado = true
			}
			break
		}
	}

	// retornar por el canal la info actualizada
	go func() {
		chInfo <- myInfo
	}()
}

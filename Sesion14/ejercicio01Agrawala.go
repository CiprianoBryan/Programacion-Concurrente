package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

var addrs []string // libreta de direcciones de los nodos de la red
var ticket int
var direccion string // localhost:9001

// estructura que se comparte entre los nodos
type Info struct {
	Tipo     string
	NumNodo  int
	AddrNodo string
}

var chIniciar chan bool // esperar que el nodo inicie su trabajo en Sección critica
var chMyInfo chan MyInfo

func main() {
	// configuracion de la red
	var n int
	fmt.Print("Ingrese la direccion del nodo: ")
	fmt.Scanf("%s\n", &direccion) // direccionamos el almacenamiento en memoria

	// 1. Solicitar las direcciones de los nodos
	fmt.Print("Ingrese la cantidad de nodos de la red: ")
	fmt.Scanf("%d\n", &n)

	// guardar las direcciones en la bitácora
	addrs = make([]string, n)

	for i, _ := range addrs {
		fmt.Printf("Nodo %d = ", i+1)
		fmt.Scanf("%s\n", &addrs[i]) // guardar directamente al arreglo
	}

	fmt.Println("La bitacora es: ", addrs)

	///////////////
	// 2.- Generar el ticket
	rand.Seed(time.Now().UTC().UnixNano()) // semilla para valores aleatoreos
	ticket = rand.Intn(1000000)
	fmt.Println(ticket)

	// inicializar / crear canales
	chIniciar = make(chan bool)
	chMyInfo = make(chan MyInfo)

	// enviar un mensaje inicial
	go func() {
		chMyInfo <- MyInfo{0, true, 1000001, ""}
	}()

	// 3.- Inicio del proceso
	go func() {
		fmt.Println("Presione enter para iniciar: ")
		// pausa
		bufferIn := bufio.NewReader(os.Stdin)
		bufferIn.ReadString('\n') // pausa

		// crear la info a enviar
		info := Info{"ENVIOTICKET", ticket, direccion}
		// notificar a todos los nodos de la bitácora
		for _, addr := range addrs {
			go enviar(addr, info)
		}
	}()

	// 4.- Publicar el servicio / ROL SERVIDOR
	serviciosSC()
}

func enviar(addr string, info Info) {
	con, _ := net.Dial("tcp", addr)
	defer con.Close()
	// codificar el mensaje a enviar a los nodos
	binfo, _ := json.Marshal(info)
	// enviar
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
	var info Info
	json.Unmarshal([]byte(bInfo), &info)
	fmt.Println(info)
	// evaluación segun el ticket
	switch info.Tipo {
	case "ENVIOTICKET":
		//recuperar del canal la info del nodo
		myInfo := <-chMyInfo
		if info.NumNodo < ticket {
			myInfo.Primero = false
		} else if info.NumNodo < myInfo.ProxNum {
			myInfo.ProxNum = info.NumNodo
			myInfo.ProxAddr = info.AddrNodo
		}

		//actualiza en uno el contador
		myInfo.ContadorMsg++

		// retornar por el canal la info actualizada
		go func() {
			chMyInfo <- myInfo
		}()

		procesarSC()
	case "INICIAR":
		fmt.Print("chIniciar: ")
		<-chIniciar
		procesarSC()
	}
}

func procesarSC() {
	fmt.Println("Inicia las tareas de la SC")
	myInfo := <-chMyInfo

	fmt.Println("Procesando la SC")
	if myInfo.ProxAddr == "" {
		fmt.Println("Soy el último nodo, SC Procesada!")
	} else {
		// notifica al nodo que le continua
		fmt.Println("trabaja concluido, SC finalizada!")
		// enviat la notificacion al próximo
		info := Info{Tipo: "INICIAR"}
		enviar(myInfo.ProxAddr, info)
	}
}

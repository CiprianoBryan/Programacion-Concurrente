package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

var localhostRegist string
var localhostNotify string
var remotehost string     // para comunicar la solicitud de registro
var bitacoraAddr []string // guardar las identificaciones de los nodos de la red localhost:puertoNot

func main() {
	////////////// CONFIGURACIÓN LOCAL //////////////
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingresa el puerto de registro: ") // el puerto por el cual se va a habilitar las comunicaciones para el registro
	port, _ := bufferIn.ReadString('\n')
	port = strings.TrimSpace(port)
	localhostRegist = fmt.Sprintf("localhost:%s", port) // creamos el identificador para la comunicacion de registro

	fmt.Print("Ingresa el puerto de notificación: ") // el puerto por el cual se va a habilitar las notificaciones para la actualizacion de la bitacora
	port, _ = bufferIn.ReadString('\n')
	port = strings.TrimSpace(port)
	localhostNotify = fmt.Sprintf("localhost:%s", port)
	/////////////////////////////////////////////////

	////////////// habilitar el SERVICIO para RECIBIR SOLICITUD DE REGISTRO //////////////
	go habilitarServicioRegistro() // rol de servidor (está escuchando)
	//////////////////////////////////////////////////////////////////////////////////////

	////////////// ENVIAR SOLICITUD DE REGISTRO //////////////
	fmt.Print("Ingrese el puerto de registro del nodo remoto: ") // nodo al cual se quiere comunicar
	port, _ = bufferIn.ReadString('\n')
	port = strings.TrimSpace(port)
	remotehost = fmt.Sprintf("localhost:%s", port) // para comunicarnos con el nodo remoto
	// rol cliente
	if port != "" {
		solicitarRegistro(remotehost)
	}
	//////////////////////////////////////////////////////////

	////////////// RECIBIR NOTIFICACIÓN DE NUEVO NODO A REGISTRAR //////////////
	recibirNotificacion() // recibe las ips de los nuevos nodos que se unen a la red
	////////////////////////////////////////////////////////////////////////////
}

//////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////// ROL SERVIDOR ////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////
func habilitarServicioRegistro() {
	// habilitar un puerto para escuchar las peticiones de registro
	listen, _ := net.Listen("tcp", localhostRegist) // comunicación (RECIBIR)
	defer listen.Close()

	for {
		// aceptar las conexiones
		con, _ := listen.Accept()
		go manejarRegistro(con) // maneja concurrencia (varios)
	}
}

func manejarRegistro(con net.Conn) {
	defer con.Close()
	// 1. recuperar la IP del nodo que solicita el registro en la red
	bufferIn := bufio.NewReader(con)
	ip, _ := bufferIn.ReadString('\n')
	ip = strings.TrimSpace(ip)
	// 2. enviar como respuesta al nodo solicitante la bitácora + la identificación del nodo local
	jsonBytes, _ := json.Marshal(append(bitacoraAddr, localhostNotify))
	fmt.Fprintln(con, string(jsonBytes)) // Aquí se envía la respuesta (bitacora) al nuevo nodo (solicitante)
	//////////////////////////
	// 3. luego este nodo debe informar al resto de nodos de la red, que hay un nuevo nodo integrante con su identificador (ip)
	comunicarTodos(ip)
	//////////////////////////
	// 4. se debe agregar el identificador del nuevo nodo a la bitacora
	bitacoraAddr = append(bitacoraAddr, ip)
	fmt.Println(bitacoraAddr)
}

func comunicarTodos(ip string) {
	// recorrer la bitacora y hacer un dial
	// se notifica a cada nodo de la red mediante su puerto de notificación la ip del nuevo nodo
	for _, addr := range bitacoraAddr {
		notificar(addr, ip) // notificar a esa dirección la ip del nuevo nodo
	}
}

func notificar(addr, ip string) {
	con, _ := net.Dial("tcp", addr) // comunicación (ENVIAR) con el resto de nodos de la red para informar la IP del nuevo nodo
	defer con.Close()

	fmt.Fprintln(con, ip) // transmitimos la IP hacia los nodos de la red q estan en la bitacora por el puerto de notificacion
}

//////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////// ROL CLIENTE /////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////

func solicitarRegistro(remotehost string) {
	con, _ := net.Dial("tcp", remotehost)
	defer con.Close()

	// 1. enviar el identificador de notificacion al nodo remoto
	fmt.Fprintln(con, localhostNotify)
	// 2. recibe la bitácora del nodo remoto
	bufferIn := bufio.NewReader(con)
	bitacoraNodoBytes, _ := bufferIn.ReadString('\n')
	var bitacoraNodo []string
	json.Unmarshal([]byte(bitacoraNodoBytes), &bitacoraNodo) // decodificando
	bitacoraAddr = bitacoraNodo
	fmt.Println(bitacoraAddr) // pintamos la bitacora actual
}

//////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////// RECIBIR NOTIFICACION ////////////////////////////

func recibirNotificacion() {
	listen, _ := net.Listen("tcp", localhostNotify)
	defer listen.Close()

	for {
		con, _ := listen.Accept()
		go atenderNotificacion(con)
	}
}

func atenderNotificacion(con net.Conn) {
	defer con.Close()
	bufferIn := bufio.NewReader(con)
	ip, _ := bufferIn.ReadString('\n')
	ip = strings.TrimSpace(ip)
	// actualizar la bitacora
	bitacoraAddr = append(bitacoraAddr, ip)
	fmt.Println(bitacoraAddr)
}

//////////////////////////////////////////////////////////////////////////////////////

//// README
/*
2. SOLICITAR REGISTRO
1. RECIBE LA SOLICITUD DE REGISTRO
   NOTIFICA AL RESTO DE NODOS DE LA RED LA IP DEL NUEVO NODO
3. CADA NODO DE LA RED RECIBE LA NOTIFICACION DE LA IP DEL NUEVO NODO
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

var localhostRegist string
var localhostNotify string
var localhostHotpotate string
var remotehost string       // para comunicar la solicitud de registro
var bitacoraAddr []string   // guardar las identificaciones de los nodos de la red localhost:puertoNot
var bitacoraAddrHP []string // guardar las identificaciones de los nodos de la red para servicio de HP

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

	////////////// Solicitar el puerto de HP para el procesamiento posterior //////////////
	fmt.Print("Ingrese el puerto de HP: ")
	port, _ = bufferIn.ReadString('\n')
	port = strings.TrimSpace(port)
	localhostHotpotate = fmt.Sprintf("localhost:%s", port)
	/////////////////////////////////////////////////

	////////////// 1: habilitar el SERVICIO para RECIBIR SOLICITUD DE REGISTRO //////////////
	go habilitarServicioRegistro() // rol de servidor (está escuchando)
	// el servicio para procesamiento de HP
	go habilitarServicioHP()
	//////////////////////////////////////////////////////////////////////////////////////

	////////////// 2: ENVIAR SOLICITUD DE REGISTRO //////////////
	fmt.Print("Ingrese el puerto de registro del nodo remoto: ") // nodo al cual se quiere comunicar
	port, _ = bufferIn.ReadString('\n')
	port = strings.TrimSpace(port)
	remotehost = fmt.Sprintf("localhost:%s", port) // para comunicarnos con el nodo remoto
	// rol cliente
	if port != "" {
		solicitarRegistro(remotehost)
	}
	//////////////////////////////////////////////////////////

	////////////// 3: RECIBIR NOTIFICACIÓN DE NUEVO NODO A REGISTRAR //////////////
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
	bufferIn := bufio.NewReader(con)

	/////////////// EL SERVIDOR RECIBE 2 PUERTOS (1: NOTIFICACION, 2: HP)
	/////////////// EL SERVIDOR RESPONDE 2 BITACORAS (1: NOTIFICACION, 2: HP)
	// 1:::: RECEPCION DEL PRIMER IP:PUERTO
	// 1. recuperar la IP del nodo que solicita el registro en la red
	ip, _ := bufferIn.ReadString('\n')
	ip = strings.TrimSpace(ip)
	// 2. enviar como respuesta al nodo solicitante la bitácora + la identificación del nodo local
	jsonBytes, _ := json.Marshal(append(bitacoraAddr, localhostNotify))
	fmt.Fprintln(con, string(jsonBytes)) // Aquí se envía la respuesta (bitacora) al nuevo nodo (solicitante)

	// 2:::: RECEPCION DEL SEGUNDO IP:PUERTO
	ipHP, _ := bufferIn.ReadString('\n')
	ipHP = strings.TrimSpace(ipHP)
	jsonBytes, _ = json.Marshal(append(bitacoraAddrHP, localhostHotpotate))
	fmt.Fprintln(con, string(jsonBytes))

	//////////////////////////
	// 3. luego este nodo debe informar al resto de nodos de la red, que hay un nuevo nodo integrante con su identificador (ip)
	comunicarTodos(ip, ipHP)
	//////////////////////////
	// 4. se debe agregar el identificador del nuevo nodo a la bitacora
	bitacoraAddr = append(bitacoraAddr, ip)
	bitacoraAddrHP = append(bitacoraAddrHP, ipHP)
	fmt.Println(bitacoraAddr)
	fmt.Println(bitacoraAddrHP)
}

func comunicarTodos(ip, ipHP string) {
	// recorrer la bitacora y hacer un dial
	// se notifica a cada nodo de la red mediante su puerto de notificación la ip del nuevo nodo
	for _, addr := range bitacoraAddr {
		notificar(addr, ip, ipHP) // notificar a esa dirección la ip del nuevo nodo
	}
}

func notificar(addr, ip, ipHP string) {
	con, _ := net.Dial("tcp", addr) // comunicación (ENVIAR) con el resto de nodos de la red para informar la IP del nuevo nodo
	defer con.Close()

	fmt.Fprintln(con, ip)   // transmitimos la IP:PUERTO NOTIFICACION hacia los nodos de la red q estan en la bitacora por el puerto de notificacion
	fmt.Fprintln(con, ipHP) // transmitimos el puerto HP hacia los nodos de la red mediante el puerto de notificacion
}

//////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////// SERVER HP /////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////
func habilitarServicioHP() {
	listen, _ := net.Listen("tcp", localhostHotpotate)
	defer listen.Close()
	for {
		con, _ := listen.Accept()
		go manejadorHP(con)
	}
}

func manejadorHP(con net.Conn) {
	defer con.Close()
	// recuperar lo que llega a traves de la conexion
	bufferIn := bufio.NewReader(con)
	strNum, _ := bufferIn.ReadString('\n')
	strNum = strings.TrimSpace(strNum)
	num, _ := strconv.Atoi(strNum)

	fmt.Printf("El número que llegó es %d\n", num)

	// lógica
	if num == 0 {
		fmt.Println("Finaliza el proceso")
	} else {
		enviarProximo(num - 1)
	}
}

func enviarProximo(num int) {
	// de forma aleatoria se va a seleccionar la identificacion del nodo
	indice := rand.Intn(len(bitacoraAddrHP))
	fmt.Printf("Enviando el num %d al nodo %s", num, bitacoraAddrHP[indice])

	// comunicar
	con, _ := net.Dial("tcp", bitacoraAddrHP[indice])
	defer con.Close()
	fmt.Fprintln(con, num)
}

//////////////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////// ROL CLIENTE /////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////

func solicitarRegistro(remotehost string) {
	con, _ := net.Dial("tcp", remotehost)
	defer con.Close()

	///////////////// SE ENVIAN 2 PUERTOS AL REMOTEHOST (SERVIDOR) (1: NOTIFICACION, 2: HP)
	///////////////// SE RECIBEN 2 BITACORAS DEL REMOTEHOST (SERVIFDOR) (1: NOTIFICACION, 2: HP)
	////////////////// 1: REGISTRO NODO
	// 1. enviar el identificador de notificacion al nodo remoto
	fmt.Fprintln(con, localhostNotify) // SE ENVÍA EL IP:PUERTO DEL NODO
	// 2. recibe la bitácora del nodo remoto
	bufferIn := bufio.NewReader(con)
	bitacoraNodoBytes, _ := bufferIn.ReadString('\n')
	var bitacoraNodo []string
	json.Unmarshal([]byte(bitacoraNodoBytes), &bitacoraNodo) // decodificando
	bitacoraAddr = bitacoraNodo
	fmt.Println(bitacoraAddr) // pintamos la bitacora actual

	////////////////// 2: REGISTRO HP
	fmt.Fprintln(con, localhostHotpotate) // SE ENVÍA EL IP:PUERTO DEL HOTPOTATE
	bitacoraNodoBytes, _ = bufferIn.ReadString('\n')
	var bitacoraNodoHP []string
	json.Unmarshal([]byte(bitacoraNodoBytes), &bitacoraNodoHP)
	bitacoraAddrHP = bitacoraNodoHP
	fmt.Println(bitacoraAddrHP)
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
	// RECIBE 2 IP:PUERTOS (1: NOTIFICACION, 2: HP) PARA AGREGAR A LA BITACORA DEL ACTUAL NODO
	defer con.Close()
	bufferIn := bufio.NewReader(con)
	// 1.- Recepción de IP:PUERTO notificación
	ip, _ := bufferIn.ReadString('\n')
	ip = strings.TrimSpace(ip)
	// actualizar la bitacora
	bitacoraAddr = append(bitacoraAddr, ip)
	fmt.Println(bitacoraAddr)

	// 2.- Recepción de IP:PUERTO de HP
	ipHP, _ := bufferIn.ReadString('\n')
	ipHP = strings.TrimSpace(ipHP)
	// actualizar la bitacora de HP
	bitacoraAddrHP = append(bitacoraAddrHP, ipHP)
	fmt.Println(bitacoraAddrHP)
}

//////////////////////////////////////////////////////////////////////////////////////

//// README
/*
2. SOLICITAR REGISTRO
 2. HP SOLICITAR REGISTRO
1. RECIBE LA SOLICITUD DE REGISTRO
   NOTIFICA AL RESTO DE NODOS DE LA RED LA IP DEL NUEVO NODO
3. CADA NODO DE LA RED RECIBE LA NOTIFICACION DE LA IP DEL NUEVO NODO Y ACTUALIZA SU BITACORA AGREGANDOLO
 3. HP RECIBE LA NOTIFICACION Y ACTUALIZA SU BITACORA
*/

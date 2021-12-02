package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Info struct {
	Contador int
	Numero   int
	Suma     int
}

var remotehost string

func main() {
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el puerto HP del nodo remoto: ")
	port, _ := bufferIn.ReadString('\n')
	port = strings.TrimSpace(port)

	remotehost = fmt.Sprintf("localhost:%s", port)

	fmt.Println("Ingrese el valor de N: ")
	strNum, _ := bufferIn.ReadString('\n')
	strNum = strings.TrimSpace(strNum)

	num, _ := strconv.Atoi(strNum)
	var infoJson Info
	infoJson.Contador = 0
	infoJson.Numero = num
	infoJson.Suma = 0
	infoJsonBytes, _ := json.Marshal(infoJson)

	// comunicar al nodo remoto
	con, _ := net.Dial("tcp", remotehost)
	defer con.Close()
	fmt.Fprintln(con, string(infoJsonBytes))
}

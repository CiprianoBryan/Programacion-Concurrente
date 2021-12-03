package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

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

	// comunicar al nodo remoto
	con, _ := net.Dial("tcp", remotehost)
	defer con.Close()
	fmt.Fprintln(con, num)
}





localhost:9101

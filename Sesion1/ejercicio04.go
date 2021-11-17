package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese un número: ")
	ingreso, err := bufferIn.ReadString('\n')
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}

	ingreso = strings.TrimRight(ingreso, "\r\n")

	num, err := strconv.Atoi(ingreso)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	doble := num * 2
	fmt.Println("El doble del número", num, "es", doble)
}

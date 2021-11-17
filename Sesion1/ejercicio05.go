package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese su nombre: ")
	name, _ := bufferIn.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")

	menu :=
		`
		****Carta****
		[1] Pizza
		[2] Empanada
		¿Cuál es tu preferido?
	`

	fmt.Println("Bienvenido ", name)
	fmt.Println(menu)

	option, _ := bufferIn.ReadString('\n')
	option = strings.TrimRight(option, "\r\n")

	switch option {
	case "1":
		fmt.Println("Usted eligió comer Pizza")
	case "2":
		fmt.Println("Usted eligió comer Empanada")
	}
}

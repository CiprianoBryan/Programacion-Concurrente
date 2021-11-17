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
	num_str, _ := bufferIn.ReadString('\n')
	num_str = strings.TrimRight(num_str, "\r\n")
	num, _ := strconv.Atoi(num_str)
	evalua(num)
}

func evalua(num int) {
	if num%2 == 0 {
		fmt.Println("El número ", num, " es divisible por 2")
	} else if num%3 == 0 {
		fmt.Println("El número ", num, " es divisible por 3")
	}
}

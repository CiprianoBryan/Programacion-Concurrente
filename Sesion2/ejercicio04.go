package main

import "fmt"

func saludar(p int) {
	fmt.Println("Holass desde el proceso: ", p)
}

func pausar() {
	var input string
	fmt.Scanln(&input)
}

func main() {
	for i := 0; i < 5; i++ {
		go saludar(i)
	}
	pausar()
}

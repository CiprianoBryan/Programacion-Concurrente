package main

import "fmt"

var turno int = 1

func p() {
	for {
		fmt.Println("Line01 SNC P")
		fmt.Println("Line02 SNC P")
		for turno != 1 {
			// esperar
		}
		fmt.Println("Line01 SC P")
		fmt.Println("Line02 SC P")
		turno = 2
	}
}

func q() {
	for {
		fmt.Println("Line01 SNC Q")
		fmt.Println("Line02 SNC Q")
		for turno != 2 {
			// esperar
		}
		fmt.Println("Line01 SC Q")
		fmt.Println("Line02 SC Q")
		turno = 1
	}
}

func main() {
	go p()
	q()
}

package main

import (
	"fmt"
	"time"
)

func saludar() {
	fmt.Println("Holaass")
}

func main() {
	go saludar() // proceso concurrente
	time.Sleep(time.Second)
}

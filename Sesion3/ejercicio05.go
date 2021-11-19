package main

import "fmt"

func enviar(c chan int) {
	c <- 10
}

func main() {
	c := make(chan int)
	go enviar(c)
	fmt.Print(<-c)
}

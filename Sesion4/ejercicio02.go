package main

import "fmt"

func ping(c chan string) {
	for {
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
	}
}

func escribir(c chan string) {
	for {
		fmt.Println(<-c)
	}
}

func espera() {
	var input string
	fmt.Scanln(&input)
}

func main() {
	c := make(chan string)
	go ping(c)
	go pong(c)
	go escribir(c)

	espera()
}

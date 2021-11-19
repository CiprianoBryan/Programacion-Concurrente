package main

import (
	"fmt"
	"time"
)

func proceso3(c3 chan string) {
	for {
		c3 <- "Datos desde el proceso 3"
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	// funciones anonimas

	go func() {
		for {
			c1 <- "Datos desde el proceso 1"
		}
	}()

	go func() {
		for {
			c2 <- "Datos desde el proceso 2"
		}
	}()

	go proceso3(c3)

	//////////////
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-c3:
				fmt.Println("Proceso 3")
			case <-time.After(time.Second):
				fmt.Println("no llego dato despues de un segundo")
			}
		}
	}()

	//////////////
	var input string
	fmt.Scanln(&input)
}

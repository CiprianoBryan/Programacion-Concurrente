package main

import "fmt"

func publicador(id int, ch chan string) {
	d := 0
	for {
		d++
		ch <- fmt.Sprintf("Mensaje %d del publicador %d", d, id)
	}
}

func receptor(id int, ch chan string) {
	for {
		fmt.Printf("Receptor %d: %s\n", id, <-ch)
	}
}

func main() {
	ch := make(chan string) // sincronos

	for i := 0; i < 4; i++ {
		go publicador(i, ch)
		go receptor(i, ch)
	}

	receptor(4, ch)
}

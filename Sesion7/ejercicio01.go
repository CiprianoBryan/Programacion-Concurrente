package main

import (
	"fmt"
	"sync"
)

var N int = 10

var buffer []int

func productor(m *sync.Mutex) {
	d := 0
	for {
		m.Lock()
		if len(buffer) < N {
			d++
			buffer = append(buffer, d)
			fmt.Println("Depositando mensaje en la cola!!!")
		} else {
			fmt.Println("Cola llena!!!")
		}
		m.Unlock()
	}
}

func consumidor(m *sync.Mutex) {
	var d int
	for {
		m.Lock()
		if len(buffer) > 0 {
			d = buffer[0]
			buffer = buffer[1:]
			fmt.Printf("Recuperando mensaje %d de la cola!!!\n", d)
		} else {
			fmt.Print("Cola vacia!!!")
		}
		m.Unlock()
	}
}

func main() {
	m := new(sync.Mutex)
	go productor(m)
	go consumidor(m)

	// pausa
	var input string
	fmt.Scanln(&input)
}

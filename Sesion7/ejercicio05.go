package main

import (
	"fmt"
	"sync"
)

var horase int = 5
var horasb int = 3

func becario(m *sync.Mutex) {
	for {
		m.Lock()
		if horasb > 0 && horase == 0 {
			horasb--
			fmt.Printf("Becario en sala, hora: %d\n", horasb)
		} else {
			horase = 3
		}
		m.Unlock()
	}
}

func estudiante(m *sync.Mutex) {
	for {
		m.Lock()
		if horase > 0 {
			horase--
			fmt.Printf("Estudiante en sala, hora: %d\n", horase)
		} else {
			horasb = 5
		}
		m.Unlock()
	}
}

func main() {
	m := new(sync.Mutex)
	go becario(m)
	go estudiante(m)

	var input string
	fmt.Scanln(&input)
}

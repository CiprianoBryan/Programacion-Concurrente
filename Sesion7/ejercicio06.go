package main

import (
	"fmt"
	"sync"
)

var estudEnSala [5]bool
var encarEnSala [3]bool
var cantEstud int = 0
var cantEncar int = 0

func estudiante(id int, m *sync.Mutex) {
	for {
		m.Lock()
		if !estudEnSala[id] && cantEstud != 5 {
			fmt.Printf("Estudiante %d ingresó a la sala\n", id)
			estudEnSala[id] = true
			cantEstud++
			if cantEstud == 5 {
				cantEncar = 0
				for i := 0; i < 3; i++ {
					encarEnSala[i] = false
				}
			}
		}
		m.Unlock()
	}
}

func encargado(id int, m *sync.Mutex) {
	for {
		m.Lock()
		if !encarEnSala[id] && cantEstud == 5 {
			fmt.Printf("Encargado %d ingresó a la sala\n", id)
			encarEnSala[id] = true
			cantEncar++
			if cantEncar == 3 {
				cantEstud = 0
				for i := 0; i < 5; i++ {
					estudEnSala[i] = false
				}
			}
		}
		m.Unlock()
	}
}

func main() {
	m := new(sync.Mutex)
	for i := 0; i < 5; i++ {
		go estudiante(i, m)
	}
	for i := 0; i < 3; i++ {
		go encargado(i, m)
	}

	var input string
	fmt.Scanln(&input)
}

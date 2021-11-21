package main

import (
	"fmt"
	"sync"
	"time"
)

var enSala int = 0

func becario(m *sync.Mutex) {
	intento_ingreso := 1
	for {
		m.Lock()
		if enSala == 0 {
			fmt.Println("Becario entrando a la sala")
			enSala = 1
			intento_ingreso = 1
		}
		m.Unlock()
		if enSala == 1 {
			h := 3
			for i := 0; i < h; i++ {
				fmt.Printf("Becario leyendo hora %d/%d\n", i, h)
			}
			m.Lock()
			fmt.Printf("Becario saliendo de la sala\n")
			enSala = 0
			m.Unlock()
			time.Sleep(time.Millisecond)
		} else {
			fmt.Printf("Becario intentando ingresar a la sala, intento nro %d\n", intento_ingreso)
			intento_ingreso++
		}
	}
}

func estudiante(m *sync.Mutex) {
	intento_ingreso := 1
	for {
		m.Lock()
		if enSala == 0 {
			fmt.Println("Estudiante entrando a la sala")
			enSala = 2
			intento_ingreso = 1
		}
		m.Unlock()
		if enSala == 2 {
			h := 2
			for i := 0; i < h; i++ {
				fmt.Printf("Estudiante leyendo hora %d/%d\n", i, h)
			}
			m.Lock()
			fmt.Printf("Estudiante saliendo de la sala\n")
			enSala = 0
			m.Unlock()
			time.Sleep(time.Millisecond)
		} else {
			fmt.Printf("Estudiante intentando ingresar a la sala, intento nro %d\n", intento_ingreso)
			intento_ingreso++
		}
	}
}

func main() {
	m := new(sync.Mutex)
	go becario(m)
	go estudiante(m)

	var input string
	fmt.Scanln(&input)
}

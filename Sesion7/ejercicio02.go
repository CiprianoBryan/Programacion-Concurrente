package main

import (
	"fmt"
	"sync"
)

func filosofo(id int, fork1, fork2 sync.Mutex) {
	for {
		fmt.Printf("Filosofo %d, Pensando!!!\n", id)
		fork1.Lock()
		fork2.Lock()
		fmt.Printf("Filosofo %d, Comiendo!!!\n", id)
		fork1.Unlock()
		fork2.Unlock()
	}
}

func main() {
	fork := make([]sync.Mutex, 5)
	go filosofo(1, fork[0], fork[1])
	go filosofo(2, fork[1], fork[2])
	go filosofo(3, fork[2], fork[3])
	go filosofo(4, fork[3], fork[4])
	filosofo(5, fork[4], fork[0])
}

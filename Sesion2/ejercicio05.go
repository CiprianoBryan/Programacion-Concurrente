package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n int

func esperar() {
	duracion := time.Duration(rand.Intn(251))
	time.Sleep(time.Millisecond * duracion)
}

func pq() {
	var temp int
	for i := 0; i < 10; i++ {
		temp = n
		esperar()
		n = temp + 1
	}
}

func main() {
	go pq() // proceso 1
	go pq() // proceso 2
	time.Sleep(time.Second)
	fmt.Println(n)
}

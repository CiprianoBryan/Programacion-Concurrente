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

func p(k int) {
	for i := 0; i < k; i++ {
		temp := n
		esperar()
		n = temp + 1
	}
}

func q(k int) {
	for i := 0; i < k; i++ {
		temp := n
		esperar()
		n = temp - 1
	}
}

func main() {
	go p(4)
	go q(4)
	time.Sleep(time.Second)
	fmt.Println(n)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n int

func pq() {
	var temp int
	temp = n
	duracion := time.Duration(rand.Intn(251))
	time.Sleep(time.Millisecond * duracion)
	n = temp + 1
}

func main() {
	go pq()
	go pq()
	time.Sleep(time.Second)

	fmt.Println(n)
}

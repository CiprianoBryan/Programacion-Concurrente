package main

import (
	"fmt"
	"time"
)

var n int

func p(k int) {
	var temp int
	for i := 0; i < k; i++ {
		temp = n
		time.Sleep(time.Second)
		n = temp + 1
	}
}

func q(k int) {
	var temp int
	for i := 0; i < k; i++ {
		temp = n
		time.Sleep(time.Second)
		n = temp - 1
	}
}

func main() {
	go p(4)
	go q(4)
	time.Sleep(time.Second)
	fmt.Println(n)
}

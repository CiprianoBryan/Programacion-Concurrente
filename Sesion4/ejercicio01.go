package main

import "fmt"

func p(c chan int, val int) {
	c <- val
}

func main() {
	c := make(chan int)

	go p(c, 1)
	go p(c, 2)
	go p(c, 3)

	num1 := <-c
	num2 := <-c
	num3 := <-c

	fmt.Println(num1, num2, num3)
}

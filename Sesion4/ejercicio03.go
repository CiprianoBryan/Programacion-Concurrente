package main

import "fmt"

func sumar(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func main() {
	arr := []int{7, 2, 8, 1, 9, 6}

	c := make(chan int)

	go sumar(arr[:len(arr)/2], c)
	go sumar(arr[len(arr)/2:], c)

	a, b := <-c, <-c
	fmt.Println("La suma es: ", a+b)
}

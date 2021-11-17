package main

import "fmt"

func main() {
	x := make([]float64, 5)
	fmt.Println(x)

	arr := [5]float64{1, 2, 3, 4, 5}
	h := arr[0:3]
	fmt.Println(h)
}

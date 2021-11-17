package main

import "fmt"

func main() {
	arreglo := [7]int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range arreglo {
		fmt.Printf("El valor de v es %d en la posición #%d\n", v, i)
	}

	i := 0
	c := 0
	for {
		if arreglo[i]%2 == 0 {
			fmt.Println(arreglo[i])
			c++
		}
		i++
		if c == 2 {
			break
		}
	}
	fmt.Println("***************")
	i = 1
	c = 0
	for c < 10 {
		if i%3 == 0 {
			fmt.Println(i)
			c++
		}
		i++
	}
}

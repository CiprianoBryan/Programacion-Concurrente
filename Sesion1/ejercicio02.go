package main

import "fmt"

// variables globales
const IGV float64 = 18.0

func main() {
	var x string = "Hola"
	fmt.Println(x)

	// declaración abreviada
	dato := 20
	fmt.Println("El doble de ", dato, " es: ", 2*dato)

	var (
		nombre string
		edad   int
	)

	nombre = "Juan"
	edad = 23

	fmt.Println("Nombre: ", nombre, " edad: ", edad)
}

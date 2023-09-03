package main

import "fmt"

const ebulicaoK float64 = 373.2

func main()  {

	tempK := ebulicaoK
	tempC := tempK - 273 //Kelvin to Celsius
	fmt.Printf("The boiling point of water in Kelvin = %.2f.\n", tempK)
	fmt.Printf("The boiling point of water in Celsius = %.2f.\n", tempC)
}


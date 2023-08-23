package main

import "fmt"

func printFigura(cantidad int) {

	if cantidad%2 == 0 || cantidad <= 0 {
		fmt.Println("La cantidad debe de ser positiva impar")
	}

	for i := 0; i < cantidad; i += 2 {
		for j := 0; j < (cantidad-i)/2; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := cantidad; i > 0; i -= 2 {
		for j := 0; j < (cantidad-i)/2; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func main() {

	var base = 17

	printFigura(base)

}

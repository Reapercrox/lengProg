package main

import (
	"fmt"
	_ "fmt"
)

func cuentaPalabras(texto string) int { //Se encarga de contar las palabras de un texto

	var total = 0

	for i := 0; i < len(texto); i++ {
		if texto[i] == ' ' || texto[i] == '\n' {
			total += 1
		}
	}

	total += 1

	return total

}

func cuentaLineas(texto string) int { //Se encarga de contar las lineas de un texto

	var numLineas = 1

	for _, char := range texto {
		if char == '\n' {
			numLineas += 1
		}
	}

	return numLineas
}

func ejercicio1(texto string) { //Funcion principal que es invocada en el main para resolver el ejercicio

	var numChars = len([]rune(texto)) //Cuenta la cantidad de caracteres convirtiendo primero a rune

	var numPalabras = cuentaPalabras(texto) //Usa func auxiliar

	var numLineas = cuentaLineas(texto) //Usa func auxiliar

	//Seccion de impresion de resultados
	fmt.Println("Cantidad de caracteres: ", numChars)
	fmt.Println("Cantidad de palabras: ", numPalabras)
	fmt.Println("Cantidad de lineas: ", numLineas)

}

func main() {

	var texto = "All that is gold does not glitter,\nNot all those who wander are lost;\nThe old that is strong does not wither,\nDeep roots are not reached by the frost."

	ejercicio1(texto)
}

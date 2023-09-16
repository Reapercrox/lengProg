package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

//Guillermo Coto Álvarez
//2016134133
//Hecho en intelij

func map3(list [][]rune, f func([][]rune, int, int, int) []rune, x int, y int, n int) []rune {
	result := make([]rune, 0)
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		//s := reflect.ValueOf(list)
		//for i := 0; i < s.Len(); i++ {
		result = append(result, f(list, x, y, n)...)
		//}
	}
	return result
}

func vertical(matriz [][]rune, x int, y int, n int) []rune {
	result := make([]rune, 0)
	if x > len(matriz) {
		println("Error, la cantidad de filas no corresponde al tamaño de la matriz")
		return nil
	} else if y > len(matriz[0]) {
		println("Error, la cantidad de columnas no corresponde al tamaño de la matriz")
		return nil
	} else if n > len(matriz) {
		println("Error, la cantidad de digitos a representar no corresponde al tamaño de la matriz")
		return nil
	}
	for i := 0; i < n; i++ {
		result = append(result, matriz[x+i][y])
	}
	return result
}

func horizontal(matriz [][]rune, x, y, n int) []rune {
	result := make([]rune, 0)
	if x > len(matriz) {
		println("Error, la cantidad de filas no corresponde al tamaño de la matriz")
		return nil
	} else if y > len(matriz[0]) {
		println("Error, la cantidad de columnas no corresponde al tamaño de la matriz")
		return nil
	} else if n > len(matriz) {
		println("Error, la cantidad de digitos a representar no corresponde al tamaño de la matriz")
		return nil
	}
	for i := 0; i < n; i++ {
		result = append(result, matriz[x][y+i])
	}
	return result
}

func diagonal(matriz [][]rune, x, y, n int) []rune {
	result := make([]rune, 0)
	if x > len(matriz) {
		println("Error, la cantidad de filas no corresponde al tamaño de la matriz")
		return nil
	} else if y > len(matriz[0]) {
		println("Error, la cantidad de columnas no corresponde al tamaño de la matriz")
		return nil
	} else if n > len(matriz) {
		println("Error, la cantidad de digitos a representar no corresponde al tamaño de la matriz")
		return nil
	}
	for i := 0; i < n; i++ {
		result = append(result, matriz[x+i][y+i])
	}
	return result
}

func main() {
	//matrix := [][]rune{
	//{1, 'a', 3, 12},
	//{'b', 5, 6, 'f'},
	//{7, 8, 'c', 11},
	//{13, 'd', 'e', 14},
	//}

	rand.Seed(time.Now().UnixNano())
	min := 3
	max := 8
	n := rand.Intn(max-min+1) + min
	matrix := make([][]rune, n)
	for i := range matrix {
		matrix[i] = make([]rune, n)
		for j := range matrix[i] {
			matrix[i][j] = rand.Int31n(126) // generate a random integer between 0 and 9
		}
	}

	for j := range matrix {
		fmt.Println(matrix[j])
	}

	verticalElements := map3(matrix, vertical, 1, 2, 3)
	horizontalElements := map3(matrix, horizontal, 0, 0, 2)
	diagonalElements := map3(matrix, diagonal, 0, 0, 3)

	fmt.Print("Resultado vertical: [ ")
	for i := 0; i < len(verticalElements); i++ {
		if verticalElements[i] >= 97 && verticalElements[i] <= 122 || verticalElements[i] >= 65 && verticalElements[i] < 90 {
			fmt.Printf("%c", verticalElements[i])
			fmt.Print(" ")
		} else {
			fmt.Print(verticalElements[i], " ")
		}
	}
	fmt.Print("]")
	fmt.Println()

	fmt.Print("Resultado horizontal: [ ")
	for i := 0; i < len(horizontalElements); i++ {
		if horizontalElements[i] >= 97 && horizontalElements[i] <= 122 || horizontalElements[i] >= 65 && horizontalElements[i] < 90 {
			fmt.Printf("%c", horizontalElements[i])
			fmt.Print(" ")
		} else {
			fmt.Print(horizontalElements[i], " ")
		}
	}
	fmt.Print("]")
	fmt.Println()

	fmt.Print("Resultado diagonal: [ ")
	for i := 0; i < len(diagonalElements); i++ {
		if diagonalElements[i] >= 97 && diagonalElements[i] <= 122 || diagonalElements[i] >= 65 && diagonalElements[i] < 90 {
			fmt.Printf("%c", diagonalElements[i])
			fmt.Print(" ")
		} else {
			fmt.Print(diagonalElements[i], " ")
		}
	}
	fmt.Print("]")
}

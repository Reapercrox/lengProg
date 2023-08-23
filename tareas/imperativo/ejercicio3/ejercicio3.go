package main

import "fmt"

func rotar(lista *[]string, veces int, direccion string) {

	var listaNueva []string = nil

	var size = len(*lista)

	if direccion == "der" || direccion == "derecha" {

		listaNueva = append(listaNueva, (*lista)[size-veces:]...)
		listaNueva = append(listaNueva, (*lista)[:size-veces]...)
		fmt.Println("Lista rotacion derecha: ", listaNueva)
		//[a,b,c,d,e,f] -> [d,e,f,a,b,c]

	} else if direccion == "izq" || direccion == "izquierda" {

		listaNueva = append(listaNueva, (*lista)[veces:]...)
		listaNueva = append(listaNueva, (*lista)[:veces]...)
		fmt.Println("Lista rotacion izquierda: ", listaNueva)
		//[a,b,c,d,e,f] -> [d,e,f,a,b,c]
	} else {
		fmt.Println("La direccion ingresada es incorrecta")
	}

}

func main() {

	var lista = &[]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	fmt.Print("Secuencia original: [ ")
	for _, item := range *lista {
		fmt.Print(item, " ")
	}
	fmt.Println("]")

	rotar(lista, 5, "izq")
	rotar(lista, 5, "der")
}

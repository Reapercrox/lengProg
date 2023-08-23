package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type producto struct {
	nombre      string
	descripcion string
	montoVenta  int32
}

type Productos []producto

var factura Productos

const rangoPagoImpuestos = 20000
const porcentajeImpuesto = 0.13

func (f *Productos) agregarProducto(nom string, desc string, pre int32) {
	idx := slices.IndexFunc(*f, func(p producto) bool { return p.nombre == nom })
	if idx == -1 {
		*f = append(*f, producto{nom, desc, pre})
	} else {
		fmt.Println("cant add existing product to the list")
	}
}

func (f *Productos) calcularImpuestoFactura() int32 {
	lista := filter1(*f, func(p producto) bool {
		return p.montoVenta > rangoPagoImpuestos
	})
	lista2 := map1(lista, func(p producto) int32 {
		return int32(float64(p.montoVenta) * porcentajeImpuesto)
	})
	return reduce(lista2)
}

func (f *Productos) calcularMontoFactura() int32 {
	lista := map1(*f, func(p producto) int32 {
		return p.montoVenta
	})
	return reduce(lista)
}
func (f *Productos) calcularMontoFactura2() int32 {
	lista := map1(*f, func(p producto) int32 {
		return p.montoVenta
	})
	return reduce(lista)
}

// funciones map y filter para aplicaciones específicas
func map1(list Productos, f func(producto) int32) []int32 {
	mapped := make([]int32, len(list))

	for i, e := range list {
		mapped[i] = f(e)
	}
	return mapped
}

func filter1(list Productos, f func(producto) bool) Productos {
	filtered := make(Productos, 0)

	for _, element := range list {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func reduce(list []int32) int32 {
	var result int32 = 0
	for _, v := range list {
		result += v
	}
	return result
}

// Construir la versión con tipos genéricos de las tres funciones anteriores map/filter para que acepten slices y funciones de cualquier otro tipo
// AYUDA SOBRE TIPOS GENERICOS: https://gobyexample.com/generics

//// Probar su funcionamiento creando una lista/slice de personas y filtrando aquellas personas que sean mayores de edad

func main() {
	factura.agregarProducto("tarjeta madre", "Asus", 54200)
	factura.agregarProducto("mouse", "alámbrico", 15000)
	factura.agregarProducto("teclado", "gamer con luces", 30000)
	factura.agregarProducto("memoria ssd", "2 gb", 65000)
	factura.agregarProducto("cable video", "display port 1m", 18000)

	println(factura.calcularMontoFactura())
	println(factura.calcularImpuestoFactura())

}

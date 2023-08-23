package main

import "fmt"

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	if l.existenciaProducto(nombre) != -1 {
		(*l)[l.existenciaProducto(nombre)].cantidad += cantidad
		(*l)[l.existenciaProducto(nombre)].precio = precio

	} else {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}

	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista

func (l *listaProductos) agregarLista(productos ...producto) {
	for _, p := range productos {
		l.agregarProducto(p.nombre, p.cantidad, p.precio)
	}
}

func (l *listaProductos) existenciaProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) buscarProducto(nombre string) (*producto, int) { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	if result == -1 {
		return nil, 2
	} else {
		return &(*l)[result], 0
	}

}

//modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
//tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la anterior función en funciones posteriores, realizando los cambios
//que el uso de la nueva función ameriten

func (l *listaProductos) venderProducto(nombre string, cantidad int) {
	var prod, err = l.buscarProducto(nombre)
	var pos = l.existenciaProducto(nombre)
	fmt.Println(err)
	if err == 0 {
		if (*prod).cantidad < cantidad {
			fmt.Println("No hay las cantidades solicitadas")
		} else if (*prod).cantidad > cantidad {
			(*prod).cantidad -= cantidad
		} else {
			(*l)[pos] = (*l)[len(*l)-1]
			(*l) = (*l)[:len(*l)-1]
		}

		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar dicha acción
		//modifique posteriormente para que se ingrese de parámetro la cantidad de productos a vender
	} else {
		fmt.Println("No existe el producto")
	}
}

//haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	var nuevaLista listaProductos = nil
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			nuevaLista = append(nuevaLista, producto{(*l)[i].nombre, (*l)[i].cantidad, (*l)[i].precio})
		}
	}

	return nuevaLista
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	fmt.Println(lProductos.buscarProducto("arroz"))
	//venda productos
	lProductos.venderProducto("arroz", 1)
	fmt.Println(lProductos)

	lProductos.agregarProducto("arroz", 5, 3000)
	fmt.Println(lProductos)

	lProductos.venderProducto("arroz", 19)
	fmt.Println(lProductos)

	fmt.Println(lProductos.listarProductosMínimos())
}

package main

import "fmt"

type zapato struct {
	marca    string
	modelo   string
	precio   int
	talla    int
	cantidad int
}

type zapatos []zapato

var stockZapatos zapatos

func (l *zapatos) agregarZapatos(marca string, modelo string, precio int, talla int, cantidad int) {
	if talla < 34 || talla > 44 {
		fmt.Println("La talla ingresada es incorrecta, no ofrecemos esas tallas")
	} else {
		if l.existenciaZapatos(marca, modelo, talla) != -1 {
			(*l)[l.existenciaZapatos(marca, modelo, talla)].cantidad += cantidad
		} else {
			(*l) = append((*l), zapato{marca, modelo, precio, talla, cantidad})
		}
	}
}

// Se necesitan minimo el modelo y la talla para poder verificar la existencia en el inventario
func (l *zapatos) existenciaZapatos(marca string, modelo string, talla int) int {
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].marca == marca && (*l)[i].modelo == modelo && (*l)[i].talla == talla {
			result = i
		}
	}
	return result
}

func (l *zapatos) buscarZapatos(marca string, modelo string, talla int) (*zapato, int) {
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].marca == marca && (*l)[i].modelo == modelo && (*l)[i].talla == talla {
			result = i
		}
	}
	if result == -1 {
		return nil, 2
	} else {
		return &(*l)[result], 0
	}

}

func llenarDatos() {

	stockZapatos.agregarZapatos("Nike", "Jordan 9", 60000, 44, 4)
	stockZapatos.agregarZapatos("Nike", "Air Max 97", 150000, 44, 2)
	stockZapatos.agregarZapatos("Nike", "Jordan 9", 60000, 43, 4)
	stockZapatos.agregarZapatos("Nike", "Jordan 9", 60000, 41, 3)
	stockZapatos.agregarZapatos("Nike", "Air Max 97", 150000, 43, 2)
	stockZapatos.agregarZapatos("Adidas", "Predator", 60000, 44, 6)
	stockZapatos.agregarZapatos("Adidas", "Predator", 60000, 43, 6)
	stockZapatos.agregarZapatos("Adidas", "Predator Messi", 70000, 42, 4)
	stockZapatos.agregarZapatos("New balance", "57/40", 40000, 38, 8)
	stockZapatos.agregarZapatos("New balance", "57/40", 40000, 39, 8)
	stockZapatos.agregarZapatos("New balance", "57/40", 67000, 42, 7)

}

func (l *zapatos) vanderZapatos(marca string, modelo string, talla int, cantidad int) {
	var zapat, err = l.buscarZapatos(marca, modelo, talla)
	var pos = l.existenciaZapatos(marca, modelo, talla)
	fmt.Println(err)
	if err == 0 {
		if (*zapat).cantidad < cantidad {
			fmt.Println("No hay las cantidades solicitadas")
		} else if (*zapat).cantidad > cantidad {
			(*zapat).cantidad -= cantidad
		} else {
			(*l)[pos] = (*l)[len(*l)-1]
			(*l) = (*l)[:len(*l)-1]
		}
	}
}

func main() {
	llenarDatos()
	fmt.Println(stockZapatos)
	stockZapatos.vanderZapatos("Nike", "Jordan 9", 44, 2)
	fmt.Println(stockZapatos)
}

module ListaProductos

type Producto = {
    nombre: string
    cantidad: int
    precio: int
}

let listarExistenciaMinima minimo listaProductos =
    List.filter (fun prod -> prod.cantidad <= minimo) listaProductos

//calcularMontosVenta cantidad listaProductos
//CALCULAR EL MONTO A PAGAR POR LA CANTIDAD DE PRODUCTO INDICADA MÁS EL 13% DE IMPUESTO PARA CADA PRODUCTO DE LA LISTA




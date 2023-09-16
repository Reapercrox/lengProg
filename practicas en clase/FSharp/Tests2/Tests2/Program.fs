open Facturas
open ListaProductos


[<EntryPoint>] 
let main argv = 
    let listaProductos = [
        { nombre = "arroz"; cantidad = 15; precio = 2500 };
        { nombre = "frijoles"; cantidad = 4; precio = 2000 };
        { nombre = "leche"; cantidad = 8; precio = 1200 };
        { nombre = "cafe"; cantidad = 12; precio = 4500 }
    ]
    
    let productosMinimos = listarExistenciaMinima 10 listaProductos
    
    printfn "Productos con existencia mínima:"
    //printfn "%A" productosMinimos

    //productosMinimos |> List.iter (fun prod -> printfn "Nombre: %s, Cantidad: %d" prod.nombre prod.cantidad)
    List.iter (fun prod -> printfn "Nombre: %s, Cantidad: %d" prod.nombre prod.cantidad) productosMinimos 

    let laFactura = [
         { nombre = "tarjeta madre"; descripcion="Asus"; montoVenta=54200}
         { nombre = "mouse"; descripcion="alámbrico"; montoVenta=15000}
         { nombre = "teclado"; descripcion="gamer con luces"; montoVenta=30000}
         { nombre = "memoria ssd"; descripcion="2 gb"; montoVenta=65000}
         { nombre = "cable video"; descripcion="display port 1m"; montoVenta=18000};
    ]
    
    printfn "Monto total de la factura: %d" (calcularMontoFactura laFactura)
    printfn "Impuesto total de la factura: %d" (calcularImpuestoFactura laFactura)

    0 // return an integer exit code

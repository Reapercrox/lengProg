module Facturas

type ElProducto = {
    nombre: string
    descripcion: string
    montoVenta: int32
}
let rangoPagoImpuestos = 20000
let porcentajeImpuesto = 0.13

let lista = [1;2;3;[4]]

let calcularImpuestoFactura factura =
    let lista = factura |> List.filter (fun p -> p.montoVenta > rangoPagoImpuestos)
    let listaImpuestos = lista |> List.map (fun p -> int32 (float p.montoVenta * porcentajeImpuesto))
    List.sum listaImpuestos

let calcularMontoFactura factura =
    let listaMontos = factura |> List.map (fun p -> p.montoVenta)
    List.sum listaMontos

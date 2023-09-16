namespace ClassLibrary1

module lib =
    let rec invertir lista =
        match lista with
        | [] -> []
        | cabeza::cola -> (invertir cola) @ [cabeza]

    let invertirString lista =
        let listaInvertida = invertir lista
        let resultado = System.String( listaInvertida |> List.toArray)
        resultado

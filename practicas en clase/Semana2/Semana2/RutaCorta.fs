module RutaCorta

open Recipientes
// Grafo de prueba
let grafo = [
    ("i", ["a"; "b"]);
    ("a", ["i"; "c"; "d"]);
    ("b", ["i"; "c"; "d"]);
    ("c", ["a"; "b"; "x"]);
    ("d", ["a"; "b"; "f"]);
    ("f", ["d"]);
    ("x", ["c"])
]
let miembro elem lista =
    List.exists (fun x -> x = elem) lista


let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | (head, neighbors) :: rest ->
        if head = nodo then neighbors
        else vecinos nodo rest


let extender ruta grafo = 
    List.filter
        (fun x -> x <> [])
        (List.map  (fun x -> if (miembro x ruta) then [] else x::ruta) (vecinos (List.head ruta) grafo)) 


let rec prof2 ini fin grafo =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            List.rev (List.head ruta) :: prof_aux fin (List.tail ruta) grafo       
        else
            prof_aux fin ((List.tail ruta) @ (extender (List.head ruta) grafo)  ) grafo
    prof_aux fin [[ini]] grafo



        
    

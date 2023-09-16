//let n_esimo (n: int) (lst: int list) =
  //  lst |> List.map (fun x -> x) |> List.item (n)
  
 let n_esimo (n: int) (lst: int list) =
    if n > List.length lst then -1 //Devlueve básicamente un false pero en f# no se puede devolver un int o boolean en una misma funcion
    else lst |> List.mapi (fun i x -> if i = n then Some(x) else None) |> List.choose id |> List.head
    
let result = n_esimo 6 [1;2;3;4;5];;
printf "%A" result
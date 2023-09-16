let shiftRight (lst: int list) (n: int) =
    List.map (fun x -> 0) [1..n] @ lst |> List.take (List.length lst)
    
let shiftLeft (lst: int list) (n: int) =
    List.skip n lst @ List.map (fun x -> 0) [1..n]
    
let der = shiftRight [1;2;3;4;5;6] 2;;
printf $"%A{der}\n"

let izq = shiftLeft [1;2;3;4;5;6] 5;;
printf $"%A{izq}"

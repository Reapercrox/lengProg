// División
let division a b =
    if b = 0 then -1
    else a / b

// Suma desde 1 hasta N
let rec sumarHasta n =
    if n = 0 then 0
    else n + sumarHasta (n - 1)

// Factorial
let rec factorial n =
    if n = 0 then 1
    else n * factorial (n - 1)

// Fibonacci
let rec fibonacci n =
    if n = 0 then 0
    elif n = 1 then 1
    else fibonacci (n - 1) + fibonacci (n - 2)

// Factorial y Fibonacci con recursión de cola
let fact n =
    let rec fact_cola prod cont n =
        if cont > n then prod
        else fact_cola (cont * prod) (cont + 1) n
    fact_cola 1 1 n

(*let fib n =
    let rec fib_cola ini fin acc1 acc2 =
        if ini = fin then acc1 + acc2
        else fib_cola (ini + 1) fin (acc1 + acc2) acc1
    if n = 0 || n = 1 then 1
    else fib_cola 2 n 1 1
*)
let fib n =
    let rec fib_cola (ini:int64) (fin:int64) (acc1:int64) (acc2:int64) : int64 =
        if ini = fin then acc1 + acc2
        else fib_cola (ini + 1L) fin (acc1 + acc2) acc1
    if n = 0L || n = 1L then 1L
    else fib_cola 2L n 1L 1L    
// Determinar si un número es primo o no
let divisible n d = n % d = 0

let primo n =
    let rec primo_aux n d =
        if n = d then true
        elif divisible n d then false
        else primo_aux n (d + 1)
    if n = 0 then true
    else primo_aux n 2

let rec invertir lista =
    match lista with
    | [] -> []
    | cabeza::cola -> (invertir cola) @ [cabeza]

let rec invertir2 lista =
    if List.isEmpty lista then
        []
    else
        let cabeza = List.head lista
        let cola = List.tail lista
        (invertir2 cola) @ [cabeza]


[<EntryPoint>]    
let main argv =
    let resultado = fib 100L
    printfn "%d" resultado
    0 // return an integer exit code

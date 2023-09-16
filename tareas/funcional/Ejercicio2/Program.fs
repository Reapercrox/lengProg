let filterSub (lst: string list) (sub: string) =
    List.filter (fun (s: string) -> s.Contains(sub)) lst
    
let result = filterSub ["la casa"; "el perro"; "pintando la cerca"] "p"
printf "%A" result

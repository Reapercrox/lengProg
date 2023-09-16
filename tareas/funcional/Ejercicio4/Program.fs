open System.Collections.Generic

let vecinos currentNode graph =
    [ for (node, neighbors, weights) in graph do
        if node = currentNode then
            yield List.zip neighbors weights ]
    |> List.concat

let rec depthFirstSearchShortestRoute graph startNode endNode =
    let visited = HashSet<string>()
    let stack = Stack<(string * int list) list>()
    stack.Push([(startNode, [])])

    let rec dfs () =
        if stack.Count > 0 then
            let path = stack.Pop()
            let (currentNode, currentWeights) = List.head path
            if currentNode = endNode then
                currentWeights
            else
                if not (visited.Contains currentNode) then
                    visited.Add currentNode
                    let neighbors = vecinos currentNode graph
                    if List.length neighbors <> List.length currentWeights then
                        failwith "The lists had different lengths"
                    let newPaths = List.map2 (fun (neighbor, weight) currentWeight -> (neighbor, weight::currentWeight)) neighbors currentWeights
                    List.iter (fun newPath -> stack.Push (newPath::path)) newPaths
                dfs ()
        else
            []

    dfs ()
       
let graph = [
    ("i", ["a"; "b"], [4;5]);
    ("a", ["i"; "c"; "d"], [4;1;3]);
    ("b", ["i"; "c"; "d"], [5;2;1]);
    ("c", ["a"; "b"; "x"], [1;2;1]);
    ("d", ["a"; "b"; "f"], [3;1;3]);
    ("f", ["d"], [3]);
    ("x", ["c"], [1]);
]


let shortestRoute = depthFirstSearchShortestRoute graph "i" "x"

printf "%A" shortestRoute
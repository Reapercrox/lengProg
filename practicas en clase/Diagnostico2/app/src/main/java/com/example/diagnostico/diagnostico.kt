package com.example.diagnostico


class linkedList{
    var head: Node? = null
    var tail: Node? = null
    var length: Int = 0
    inner class Node(var value: Any?){
        var next: Node? = null
    }
}

class jugador(var Nombre: String, var Numero: Int){

}

fun main(args: Array<String>){
    println(helloWorld())
}

fun helloWorld():String{
    return "Hello world"
}


package main

import (
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)
import "fmt"
import "bufio"
import "os"

func main() {

	// connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		// what to send?
		fmt.Print("Text to send: ")
		text, err := clientReader.ReadString('\n')

		//EL FOR ES UNA PRUEBA PARA ENVIAR VARIOS MENSAJES DENTRO DE LA MISMA CONEXIÓN Y VER QUEW RECIBE BIEN LAS RESPUESTAS A CADA MENSAJE
		//NO SE HACE, AUNQUE SE PODRÍA, IDENTIFICAR CADA UNA DE LAS RESPUESTAS ASOCIADAS AL MENSAJE QUE SE PASÓ INICIALMENTE
		//PARA EFECTOS DE UNA PRUEBA MÁS ROBUSTA DENTRO DEL PROYECTO PROGRAMADO, SE PUEDE ELIMINAR ESTE FOR, PENSANDO EN QUE DEPORSÍ EL ENVÍO DE MENSAJES MÚLTIPLE
		for i := 1; i <= 5; i++ {

			switch err {
			case nil:
				text := strings.TrimSpace(text)
				if text == "quit" {
					return
				} else {
					if _, err = conn.Write([]byte(text + strconv.Itoa(i) + "\n")); err != nil {
						log.Printf("failed to send the client request: %v\n", err)
					}
				}
			case io.EOF:
				log.Println("client closed the connection")
				return
			default:
				log.Printf("client error: %v\n", err)
				return
			}

			// Waiting for the server response
			serverResponse, err := serverReader.ReadString('\n')

			switch err {
			case nil:
				log.Println("Received: ", strings.TrimSpace(serverResponse))
			case io.EOF:
				log.Println("server closed the connection")
				return
			default:
				log.Printf("server error: %v\n", err)
				return
			}
		}
	}
}

import socket

SERVER_IP = '127.0.0.1'

PORT_NUMBER = 8000
print(f"Server address: {SERVER_IP}:{PORT_NUMBER}")

client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client_socket.connect((SERVER_IP, PORT_NUMBER))

while 1:
    message = input()
    print(f"Send message: {message}")

    #EL FOR ES UNA PRUEBA PARA ENVIAR VARIOS MENSAJES DENTRO DE LA MISMA CONEXIÓN Y VER QUEW RECIBE BIEN LAS RESPUESTAS A CADA MENSAJE
    #NO SE HACE, AUNQUE SE PODRÍA, IDENTIFICAR CADA UNA DE LAS RESPUESTAS ASOCIADAS AL MENSAJE QUE SE PASÓ INICIALMENTE
    #PARA EFECTOS DE UNA PRUEBA MÁS ROBUSTA DENTRO DEL PROYECTO PROGRAMADO, SE PUEDE ELIMINAR ESTE FOR, PENSANDO EN QUE DEPORSÍ EL ENVÍO DE MENSAJES MÚLTIPLE
    for i in range(1,6):
        client_socket.sendall(bytes(message+str(i)+"\n", "utf-8"))
        data = client_socket.recv(1024)
        print(data)

import java.util.LinkedList;
import java.util.Collections;
import java.util.Comparator;

class sortAscendente implements Comparator<diagnostico>{

        public int compare(diagnostico a, diagnostico b){
            return a.numero - b.numero;
    }
}

public class diagnostico{

    private String jugador;
    public int numero;

    public diagnostico(String jugador, int numero){
        this.jugador = jugador;
        this.numero = numero;
    }

    public String getjugador(){
        return jugador;
    }

    public int getnumero(){
        return numero;
    }



    public static void main(String[] args){
        LinkedList<diagnostico> jugadores = new LinkedList<diagnostico>();

        jugadores.add(new diagnostico("Thibaut Curtois",1));
        jugadores.add(new diagnostico("Luka Modric",10));
        jugadores.add(new diagnostico("Toni Kross",8));
        jugadores.add(new diagnostico("Eduardo Camavinga",12));
        jugadores.add(new diagnostico("Karim Benzema",9));
        jugadores.add(new diagnostico("Antonio Rudiger",22));

        System.out.println("Desordenado:");
        for (int i = 0; i < jugadores.size(); i++)
            System.out.println("Jugador: "+jugadores.get(i).jugador+" Numero: "+jugadores.get(i).numero);
        System.out.println("\n");

        Collections.sort(jugadores, new sortAscendente());

        System.out.println("Ordenado:");
        for (int i = 0; i < jugadores.size(); i++)
            System.out.println("Jugador: "+jugadores.get(i).jugador+" Numero: "+jugadores.get(i).numero);
        System.out.println("\n");

    }
}
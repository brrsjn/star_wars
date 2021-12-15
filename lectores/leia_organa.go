package main

import (
	"bufio"
	"context"
	"fmt"

	//"io/ioutil"
	"log"
	"time"

	"star_wars/pb"
	"strings"

	"os"

	"google.golang.org/grpc"
)

const (
	brokeraddress = "localhost:50060"
	defaultBot    = true
)

func main() {
	fmt.Println("\n-Leia Organa: Informacion...")
	//register := make([]string, 1)
	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		_, err := fmt.Println("> ")
		inputScanner.Scan()
		temp := strings.Split(inputScanner.Text(), " ")
		//valid, tempType, large := IsValidInput(temp[0]) //falta arreglar esto
		valid := strings.ToLower(temp[0]) == "getnumbersrebelds"

		if !valid || (len(temp) != 3) {
			fmt.Println("Error, fin de la comunicacion", err)
			continue
		}

		fmt.Println("Conectando")

		fmt.Println()

		reloj_vec, rebelds_number, server := ConnectToBroker(temp[1], temp[2])
		fmt.Println("Reloj: ", reloj_vec)
		fmt.Println("Numero de rebeldes: ", rebelds_number)
		fmt.Println("Servidor Fullcrum: ", server)
	}

}

//Conexion con broker, lectura.
func ConnectToBroker(planet string, city string) (int32, int32, string) {
	conn, err := grpc.Dial(brokeraddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		fmt.Println("Canal seguro")
	}

	defer conn.Close()
	broker := pb.NewBrokerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addres, err := broker.ConnectLeia(ctx, &pb.InstructLeia{Nplanet: planet, Ncity: city})

	if err != nil {
		fmt.Println(err)
		return 0, 0, "error"
	} else {
		return addres.Vreloj, addres.Nrebelds, addres.Server
	}
}

func ConectToServer() {
	fmt.Println("Fulcrum me copia")
}

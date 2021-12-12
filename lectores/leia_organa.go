package main

import (
	"fmt"
	"bufio"
	"context"
	"io/ioutil"
	"log"
	"time"

	protos "star_wars/pb"
	"strconv"
	"strings"

	"os"

	"google.golang.org/grpc"
)

const (
	brokeraddress = "localhost:50051"
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
		valid, tempType, large := IsValidInput(temp[0]) //falta arreglar esto

		if !valid || (len(temp) !=large){
			fmt.Println("Error, fin de la comunicacion")
			continue
		}

		if tempType==1{
			_, err := strconv.Atoi(temp[len(temp)-1])
			if err !=nil{
				fmt.Println("Error, intente nuevamente")
				continue
			}
		}
		if tempType==0{
			fmt.Println("Conexion correcta")
		}

		fmt.Println("Conectando")

		fmt.Println(ConectToBroker(temp[0]))

		ConectToServer()
	}

}

//Conexion con broker, lectura.
func ConectToBroker(message string) string {
	conn, err := grpc.Dial(brokeraddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		fmt.Println("Canal seguro")
	}

	defer conn.Close()
	broker := protos.NewInformerToBrokerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addres, err := broker.ConnectToServer(ctx, &protos.Instruct{Message: message, Lectura: true})
	
	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return addres.Addres
	}
}

func ConectToServer() {
	fmt.Println("Fulcrum me copia")
}


//arreglar!!
/*
func IsValidInput(input string) (bool, int, int){
	switch input {
		case
			"AddCity", "Addcity", "addCity", "addcity":
			return true, 1, 4
	
		case
			"UpdateName", "updateName", "updatename", "Updatename":
			return true, 2, 4
	
		case
			"UpdateNumber", "Updatenumber", "updateNumber", "updatenumber":
			return true, 1, 4
	
		case
			"DeleteCity", "Deletecity", "deleteCity", "deletecity":
			return true, 2, 3
		case
			"Exit", "exit", "EXIT":
			return true, 0, 1
		}
	
		return false, 0, 0
}
*/
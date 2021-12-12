package main

import (
	"bufio"
	"context"
	"fmt"
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
	fmt.Println("\n-Ahsoka: Hola, aqui Fulcrum a su servicio...")
	registro := make([]string, 1)
	buf := bufio.NewScanner(os.Stdin)
	for {
		//mantener encendido e interpretar las entradas
		_, err := fmt.Print("> ")
		buf.Scan()
		comm := strings.Split(buf.Text(), " ")
		valid, CommType, largo := IsValidInput(comm[0])

		//Communication Failled
		if !valid || (len(comm) != largo) {
			fmt.Println("\n-Mos Eisley: Que?, bo. *Llamada desconectada* ")
			fmt.Println("-Ahsoka: Creo que me equivoque de instruccion inicial, intentemoslo otra vez...")
			continue
		}
		if CommType == 1 {
			_, err := strconv.Atoi(comm[len(comm)-1])
			if err != nil {
				fmt.Println("-Ahsoka: formato erroneo, intentemoslo otra vez...")
				continue
			}
		}
		if CommType == 0 {
			fmt.Println("-Ahsoka: Fulcrum Cambio y fuera")
			break
		}

		fmt.Println("-Mos Eisley: Comunicando...")

		//conecta al broker y envia el comando
		ConectToBroker(comm[0])

		//resive el servidor y reenvia el comando
		ConectToServer()

		//resive respuesta de todo ok y suelta al server.

		//Modifica su registro personal.
		//--- en caso de falla no guarda su registro.

		if err != nil {

		} else {
			registro = append(registro, buf.Text())
		}

	}

}

func IsValidInput(input string) (bool, int, int) {
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

func ConectToBroker(message string) string {
	conn, err := grpc.Dial(brokeraddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		fmt.Println("-Ahsoka: Mos Eisley. Aqui Fulcrum solicitando un canal seguro")
	}
	defer conn.Close()
	broker := protos.NewInformerToBrokerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	addres, err := broker.ConnectToServer(ctx, &protos.Instruct{Message: message})
	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		return addres.Addres
	}

}

func ConectToServer() {
	fmt.Println("-Ahsoka: Fulcrum me copia?")
}

func Self_AddCity(texto string, full bool) {
	if !full {
		texto = texto + "\n"
	}
	textByte := []byte(texto)
	err := ioutil.WriteFile("registro.txt", textByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func Self_ReadAll() string {
	allTextBytes, err := ioutil.ReadFile("registro.txt")
	if err != nil {
		log.Fatal(err)
	}
	allTextString := string(allTextBytes)
	return allTextString
}

package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"star_wars/pb"
	"time"

	"strconv"
	"strings"

	"os"

	"google.golang.org/grpc"
)

const (
	brokeraddress = "localhost:50050"
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

		//conecta al broker y envia el comando
		adSv, err := TalkToBroker(comm[0])

		if adSv.Addres == "empty" {
			fmt.Println("-Ahsoka: No se pudo modificar los registros galacticos")
			continue
		}

		//resive el servidor y reenvia el comando
		TalkToServer(adSv.Addres, comm)

		//resive el servidor y reenvia el comando

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

func TalkToBroker(message string) (*pb.Servidor, error) {
	conn, err := grpc.Dial(brokeraddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		fmt.Println("-Ahsoka: Mos Eisley. Aqui Fulcrum solicitando un canal seguro")
	}
	defer conn.Close()
	broker := pb.NewBrokerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addres, err := broker.ConnectToServer(ctx, &pb.Instruct{
		Message: message,
		Lectura: false})

	return addres, err
}

func TalkToServer(address string, input []string) *pb.City {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		fmt.Println("-Ahsoka: Fulcrum me copia?, envio nuevos datos galacticos...")
	}
	defer conn.Close()
	serverObj := pb.NewFulcrumClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// LE ENVIA LA INSTRUCCION AL SERVER
	switch input[0] {
	case
		"AddCity", "Addcity", "addCity", "addcity":
		nume, _ := strconv.Atoi(input[3])
		response, err := serverObj.AddCity(ctx, &pb.CityNewNumber{
			Planet:    input[1],
			City:      input[2],
			Survivors: int32(nume),
		})
		if err != nil {
			return &pb.City{
				Name: " ",
			}
		}
		return response

	case
		"UpdateName", "updateName", "updatename", "Updatename":
		response, err := serverObj.UpdateName(ctx, &pb.CityNewName{
			Planet:  input[1],
			City:    input[2],
			NewName: input[3],
		})
		if err != nil {
			return &pb.City{
				Name: " ",
			}
		}
		return response

	case
		"UpdateNumber", "Updatenumber", "updateNumber", "updatenumber":
		nume, _ := strconv.Atoi(input[3])
		response, err := serverObj.UpdateNumber(ctx, &pb.CityNewNumber{
			Planet:    input[1],
			City:      input[2],
			Survivors: int32(nume),
		})
		if err != nil {
			return &pb.City{
				Name: " ",
			}
		}
		return response

	case
		"DeleteCity", "Deletecity", "deleteCity", "deletecity":
		response, err := serverObj.DeleteCity(ctx, &pb.CityDelete{
			Planet: input[1],
			City:   input[2],
		})
		if err != nil {
			return &pb.City{
				Name: " ",
			}
		}
		return response
	}

	return &pb.City{
		Name: " ",
	}
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

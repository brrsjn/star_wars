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
	brokeraddress = "localhost:50060"
	defaultBot    = true
)

type PlanetsDic struct {
	Planet []PlanetObj
}
type PlanetObj struct {
	Name   string
	Cities map[string]*pb.City
}

type Memoria struct {
	registro string
	reloj    string
	servidor string
}

/*func main() {
	AddlineOnFiles("testing texto", false)
}// */

func main() {
	fmt.Println("\n-Thrawn: Hola, aqui Fulcrum a su servicio...")
	//var planetas map[string]PlanetObj
	planetas := map[string]PlanetObj{}
	registro := []Memoria{}

	buf := bufio.NewScanner(os.Stdin)
	for {
		/*
			mantener encendido e interpretar las entradas
		*/
		fmt.Print("> ")
		buf.Scan()
		comm := strings.Split(buf.Text(), " ")
		valid, CommType, largo := IsValidInput(comm[0])

		/*
			Communication Failled
		*/
		if !valid || (len(comm) != largo) {
			fmt.Println("-Thrawn: Creo que me equivoque de instruccion inicial, intentemoslo otra vez...")
			continue
		}
		if CommType == 1 {
			_, err := strconv.Atoi(comm[len(comm)-1])
			if err != nil {
				fmt.Println("-Thrawn: formato erroneo, intentemoslo otra vez...")
				continue
			}
		}
		if CommType == -1 {
			fmt.Println("-Thrawn: Fulcrum Cambio y fuera")
			break
		}

		/*
			conecta al broker y envia el comando
		*/
		adSv, _ := TalkToBroker(comm[0])

		if adSv.Error {
			fmt.Println("-Mos Eisley: No hay servidores disponibles...")
			continue
		}

		//resive el servidor y reenvia el comando
		cityMod := TalkToServer(adSv.Addres, comm)

		memo := new(Memoria)
		memo.registro = buf.Text() + "; Servidor: " + adSv.Addres + "; Reloj: " + cityMod.Reloj

		if !cityMod.Error {
			//Modifica su registro personal.
			registro = append(registro, *memo)
			AddlineOnFiles(memo.registro, false)
			planetas[cityMod.Planet].Cities[cityMod.Name] = cityMod
		} else {
			//en caso de falla no guarda su registro.
			fmt.Println("ha Fallado el registro planetario.")
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
		return true, -1, 1
	}

	return false, 0, 0
}

func TalkToBroker(message string) (*pb.Servidor, error) {
	conn, err := grpc.Dial(brokeraddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		fmt.Println("-Thrawn: Mos Eisley. Aqui Fulcrum solicitando un canal seguro")
	}
	defer conn.Close()
	broker := pb.NewBrokerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addres, err := broker.RandomServer(ctx, &pb.Instruct{
		Message: message,
		Lectura: false})

	return addres, err
}

func TalkToServer(address string, input []string) *pb.City {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		fmt.Println("-Thrawn: Fulcrum me copia?, envio nuevos datos galacticos...")
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
		response, err := serverObj.AddCity(ctx, &pb.City{
			Planet:    input[1],
			Name:      input[2],
			Survivors: int32(nume),
		})
		if err != nil {
			return &pb.City{
				Error: true,
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
				Error: true,
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
				Error: true,
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
				Error: true,
			}
		}
		return response
	}

	return &pb.City{
		Error: true,
	}
}

func AddlineOnFiles(texto string, full bool) {
	if !full {
		texto = texto + "\n"
	}
	before := SelfReadAll()
	textobyte := before + texto
	writeByte := []byte(textobyte)
	err := ioutil.WriteFile("informantes/almirante_thrawn/registroThrawn.txt", writeByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func SelfReadAll() string {
	allTextBytes, err := ioutil.ReadFile("informantes/almirante_thrawn/registroThrawn.txt")
	if err != nil {
		log.Fatal(err)
	}
	allTextString := string(allTextBytes)
	return allTextString
}

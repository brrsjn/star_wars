package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

		//Communication Failled
		if (len(comm) < 3) || !IsValidInput(comm[0]) || (len(comm) > 4) {
			fmt.Println("\n-Mos Eisley: Que?, aqui no tenemos de esos. *Llamada desconectada* ")
			fmt.Println("-Ahsoka: ok, intentemoslo otra vez...")
			continue
		}

		//conecta al broker y envia el comando
		ConectToBroker()

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

func IsValidInput(input string) bool {
	switch input {
	case
		"AddCity", "UpdateName", "UpdateNumber", "DeleteCity":
		return true
	}
	return false
}

func ConectToBroker() {
	fmt.Println("-Ahsoka: Mos Eisley. Aqui Fulcrum solicitando un canal seguro")
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

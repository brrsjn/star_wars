package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"star_wars/pb"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50061"
)

type FulcrumServer struct {
	pb.UnimplementedFulcrumServer
	savedPlanetas []*Planeta
}

type Planeta struct {
	nombre   string
	ciudades []string
}

func crearArchivo(path string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("File Created Successfully", path)
}

func escribeArchivo(path string, linea string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	_, err = file.WriteString(linea)
	if existeError(err) {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
	fmt.Println("Archivo actualizado existosamente.")
}

func escribeArchivoNuevo(path string, linea string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0660)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	log.Println(linea)
	_, err = file.WriteString(linea)
	if existeError(err) {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
	fmt.Println("Archivo actualizado existosamente.")
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func (s *FulcrumServer) AddCity(ctx context.Context, in *pb.City) (*pb.City, error) {
	//Codigo para guardar la ciudad en archivo
	log.Printf("Se añadirá una nueva ciudad")
	path := fmt.Sprintf("servidores/servidor_fulkrum_1/planetas/%s.txt", in.Planet)
	if len(s.savedPlanetas) > 0 {
		for i := 0; i < len(s.savedPlanetas); i++ {
			if s.savedPlanetas[i].nombre == in.Planet {
				fmt.Println("if planeta")
				isNotCityInPlaneta := true
				for j := 0; j < len(s.savedPlanetas[i].ciudades); j++ {
					fmt.Println("for ciudad")
					if s.savedPlanetas[i].ciudades[j] == in.Name {
						fmt.Println("if ciudad")
						isNotCityInPlaneta = false
						//Solo continua con la funcion, sin informar al usuario
						break
					}
				}
				if isNotCityInPlaneta {
					toWrite := fmt.Sprintf("%s %s %d\n", in.Planet, in.Name, in.Survivors)
					planeta := new(Planeta)
					planeta.nombre = in.Planet
					planeta.ciudades = append(planeta.ciudades, in.Name)
					s.savedPlanetas[i].ciudades = append(planeta.ciudades, in.Name)
					escribeArchivo(path, toWrite)
				}
			}

		}
	} else {
		crearArchivo(path)
		toWrite := fmt.Sprintf("%s %s %d\n", in.Planet, in.Name, in.Survivors)
		planeta := new(Planeta)
		planeta.nombre = in.Planet
		planeta.ciudades = append(planeta.ciudades, in.Name)
		s.savedPlanetas = append(s.savedPlanetas, planeta)
		escribeArchivo(path, toWrite)
	}
	return in, nil

}

func (s *FulcrumServer) DeleteCity(ctx context.Context, in *pb.CityDelete) (*pb.City, error) {
	//Codigo para actualizar una ciudad
	path := fmt.Sprintf("servidores/servidor_fulkrum_1/planetas/%s.txt", in.Planet)
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Entró")

	lines := strings.Split(string(input), "\n")
	var cityName string
	var newLines []string
	for _, line := range lines {
		valor := strings.Split(line, " ")
		if len(valor) == 3 {
			if valor[1] != in.City {
				newLines = append(newLines, line)
			} else {
				cityName = valor[1]
			}
		}

	}
	output := strings.Join(newLines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	return &pb.City{Name: cityName}, nil
}

func (s *FulcrumServer) UpdateName(ctx context.Context, in *pb.CityNewName) (*pb.City, error) {
	//Codigo para actualizar una ciudad
	path := fmt.Sprintf("servidores/servidor_fulkrum_1/planetas/%s.txt", in.Planet)
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Entró")

	lines := strings.Split(string(input), "\n")
	newText := ""
	for i, line := range lines {
		valor := strings.Split(line, " ")
		if len(valor) == 3 {
			if valor[1] == in.City {
				fmt.Println(len(valor))
				newText += fmt.Sprintf("%s %s %s", valor[0], in.NewName, valor[2])
				lines[i] = newText
			}
		}

	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return &pb.City{Name: in.NewName, Planet: in.Planet, Survivors: 0}, nil
}

func (s *FulcrumServer) UpdateNumber(ctx context.Context, in *pb.CityNewNumber) (*pb.City, error) {
	//Codigo cambiar el numero de la ciudad
	path := fmt.Sprintf("servidores/servidor_fulkrum_1/planetas/%s.txt", in.Planet)
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Entró")

	lines := strings.Split(string(input), "\n")
	newText := ""
	for i, line := range lines {
		valor := strings.Split(line, " ")
		if len(valor) == 3 {
			if valor[1] == in.City {
				fmt.Println(len(valor))
				fmt.Println(in.Survivors)
				newText += fmt.Sprintf("%s %s %d", valor[0], valor[1], in.Survivors)
				lines[i] = newText
			}
		}

	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return &pb.City{Name: in.City, Planet: in.Planet, Survivors: 0}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFulcrumServer(s, &FulcrumServer{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

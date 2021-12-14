package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"star_wars/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50061"
)

type BrokerFulcrumServer struct {
	pb.UnimplementedBrokerFulcrumServiceServer
}

func crearArchivo(path string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		log.Printf("Hola")
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("File Created Successfully", path)
}

func escribeArchivo(path string, movimiento int) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	if existeError(err) {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	text := fmt.Sprintf("%d\n", movimiento)
	_, err = file.WriteString(text)
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

func (s *BrokerFulcrumServer) AddCityToFulcrum(ctx context.Context, in *pb.AddCityToFulcrumRequest) (*pb.AddCityToFulcrumReply, error) {
	//Codigo para guardar la ciudad en archivo
	fmt.Println(in.NombreCiudad)

	return &pb.AddCityToFulcrumReply{State: true}, nil
}

func (s *BrokerFulcrumServer) DeleteCityToFulcrum(ctx context.Context, in *pb.DeleteCityToFulcrumRequest) (*pb.DeleteCityToFulcrumReply, error) {
	//Codigo para eliminar una ciudad del archivo
	return &pb.DeleteCityToFulcrumReply{State: true}, nil
}

func (s *BrokerFulcrumServer) UpdateNameToFulcrum(ctx context.Context, in *pb.UpdateNameToFulcrumRequest) (*pb.UpdateNameToFulcrumReply, error) {
	//Codigo para actualizar una ciudad
	return &pb.UpdateNameToFulcrumReply{State: true}, nil
}

func (s *BrokerFulcrumServer) UpdateNumberToFulcrum(ctx context.Context, in *pb.UpdateNumberToFulcrumRequest) (*pb.UpdateNumberToFulcrumReply, error) {
	//Codigo cambiar el numero de la ciudad
	return &pb.UpdateNumberToFulcrumReply{State: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBrokerFulcrumServiceServer(s, &BrokerFulcrumServer{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

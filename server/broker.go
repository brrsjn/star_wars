package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	broker "star_wars/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50070"

	address1 = "localhost:50071"
	address2 = "localhost:50072"
	address3 = "localhost:50073"
)

type BrokerObj struct {
	broker.UnimplementedBrokerServer
	servers    []*broker.Servidor
	conectedSV int32
}

func main() {

	// Set up a connection to the server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	broker.RegisterBrokerServer(s, &BrokerObj{
		conectedSV: 0,
	})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func ConnectToFulcrum(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
}

func (self *BrokerObj) RandomServer(ctx context.Context, req *broker.Instruct) (*broker.Servidor, error) {
	if self.conectedSV == 0 {
		fmt.Println("-Mos Eisley: No hay servidores disponibles...")
		return &broker.Servidor{Error: true}, nil

	} else {
		fmt.Println("-Mos Eisley: Comunicando...")
		req := self.servers[rand.Intn(int(self.conectedSV)-1)]
		return req, nil
	}

}

func (self *BrokerObj) ConnectLeia(ctx context.Context, req *broker.InstructLeia) (*broker.NumbersRebelds, error) {
	adr := self.servers[rand.Intn(int(self.conectedSV)-1)].Addres
	conn, err := grpc.Dial(adr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	serverObj := broker.NewFulcrumClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, _ := serverObj.ReadAll(ctx, &broker.Read{Nplanet: req.Nplanet, Ncity: req.Ncity})
	return &broker.NumbersRebelds{Nrebelds: res.Survivors, Vreloj: res.Reloj, Server: adr}, nil
}

func (self *BrokerObj) ServerIsOpen(ctx context.Context, req *broker.Servidor) (*broker.Servidor, error) {
	if self.conectedSV < 3 {
		req.Error = false
		self.conectedSV = self.conectedSV + 1
		self.servers = append(self.servers, req)
		string := "Adress conectado " + req.Addres
		log.Println(string)
		return req, nil

	} else {
		req.Error = true
		log.Println("No se pudo conectar el addres: " + req.Addres)
		return req, nil
	}

}

func RandomServer() string {
	adressServer := "empty"
	return adressServer
}

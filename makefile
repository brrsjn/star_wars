gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/*.proto

build:
	go mod tidy

clean:
	rm pb/*.go
	
server:
	go run broker/broker.go

lector:
	go run lectores/leia_organa.go

informante_1:
	go run informantes/ahsoka_tano/ahsoka_tano.go

informante_2:
	go run informantes/almirante_thrawn/almirante_thrawn.go

server_1:
	go run servidores/servidor_fulkrum_1/fulkrum_1.go

server_2:
	go run servidores/servidor_fulkrum_2/fulkrum_2.go

server_3:
	go run servidores/servidor_fulkrum_3/fulkrum_3.go
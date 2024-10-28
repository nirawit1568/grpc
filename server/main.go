package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()
	services.RegisterCalculatorServer(grpcServer, services.NewCalculatorServer())

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("gRPC server listening on port 50051")
}

package main

import (
	"github.com/quii/try-grpc"
	"github.com/quii/try-grpc/fridge"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", try_grpc.Port)

	if err != nil {
		log.Fatalf("problem listening to port %s, %v", try_grpc.Port, err)
	}

	server := grpc.NewServer()
	fridge.RegisterFridgeServiceServer(
		server,
		fridge.NewInMemoryFridgeService([]string{"cheese", "eggs"}),
	)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}

package account

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGRPCServer() {
	port := ":3000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	//route.RegisterUserServer(grpcServer, &routeUserServer)
	grpcServer.Serve(lis)
}
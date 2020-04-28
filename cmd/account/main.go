package main

import "github.com/pgnedoy/go-grpc/pkg/account"

func main() {
	//account.StartGRPCServer()
	account.StartHTTPServer()
}
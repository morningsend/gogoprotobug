package main

import (
	"context"
	"gogoprotobug/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

var address = "localhost:6000"

func main() {

	server := grpc.NewServer()

	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := server.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	sayHello := pb.NewGreetingClient(conn)
	_, err = sayHello.SayHello(context.Background(), &pb.MyMessage{
		MyCustomField: pb.NewMyCustomField([]byte("hello world")),
	})

	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net"

	"github.com/Feride3d/banner-rotation-service/internal/pb"
	server "github.com/Feride3d/banner-rotation-service/internal/server/http"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()            // создать сервер
	srv := &server.GRPCServer{}      // переменная со структурой, которая реализует интерфейс сервера
	pb.RegisterRotatorServer(s, srv) // зарегистрировать сервер в качестве сервера для NewServer

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

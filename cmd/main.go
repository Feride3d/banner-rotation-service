package main

import (
	"fmt"
	"os"

	server "github.com/Feride3d/banner-rotation-service/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
	}
}

	s := grpc.NewServer()            // создать сервер
	srv := &server.GRPCServer{}      // переменная со структурой, которая реализует интерфейс сервера
	pb.RegisterRotatorServer(s, srv) // зарегистрировать сервер в качестве сервера для NewServer

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil { */

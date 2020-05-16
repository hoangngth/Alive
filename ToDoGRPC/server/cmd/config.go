package cmd

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	database "Alive/ToDoGRPC/database"
	proto "Alive/ToDoGRPC/proto"
	service "Alive/ToDoGRPC/service"
)

func RunServer() {
	db := database.ConnectDatabase()
	defer db.Close()
	toDoServer := service.NewToDoServiceServer(db)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterToDoServiceServer(srv, toDoServer)
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

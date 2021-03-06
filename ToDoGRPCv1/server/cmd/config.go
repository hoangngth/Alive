package cmd

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	database "Alive/ToDoGRPCv1/database"
	proto "Alive/ToDoGRPCv1/proto"
	service "Alive/ToDoGRPCv1/service"
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

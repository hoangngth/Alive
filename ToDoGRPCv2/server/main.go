package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	database "Alive/ToDoGRPCv2/database"

	profileProto "Alive/ToDoGRPCv2/proto/profile"
	todoProto "Alive/ToDoGRPCv2/proto/todo"

	profileService "Alive/ToDoGRPCv2/service/profile"
	todoService "Alive/ToDoGRPCv2/service/todo"
)

func main() {

	toDoDB := database.ConnectDatabaseToDo()
	defer toDoDB.Close()
	toDoServer := todoService.NewToDoServiceServer(toDoDB)

	profileDB := database.ConnectDatabaseProfile()
	defer profileDB.Close()
	profileServer := profileService.NewProfileServiceServer(profileDB)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	todoProto.RegisterToDoServiceServer(srv, toDoServer)
	profileProto.RegisterProfileServiceServer(srv, profileServer)
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

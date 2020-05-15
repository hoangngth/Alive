package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"reflection"
	"Alive/Todo/pkg/proto/v1"
	"Alive/Todo/pkg/service/v1"
)

func RunServer() error {
	ctx := context.Background()
	port := ServerConfig()4
	v1Api := v1.NewToDoServiceServer(db)
	return grpc.Run(ctx, v1Api, port)
}

func Run(context.Context, v1API v1.ToDoServiceServer, port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)
	
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Println("Starting gRPC server...")
	return server.Serve(listener)
}
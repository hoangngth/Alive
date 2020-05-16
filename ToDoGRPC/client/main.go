package main

import (
	"log"

	"github.com/gin-gonic/gin"

	service "Alive/ToDoGRPC/service"
)

func main() {

	service.DialToServiceServer(8000)

	g := gin.Default()
	g.GET("/user/:userid", service.ReadAllToDo)
	g.GET("/user/:userid/:todoid", service.ReadToDo)
	g.POST("/user/:userid", service.CreateToDo)
	g.PUT("user/:userid/:todoid", service.UpdateToDo)
	g.DELETE("user/:userid/:todoid", service.DeleteToDo)

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}

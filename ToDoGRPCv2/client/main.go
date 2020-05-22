package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"Alive/ToDoGRPCv2/middleware"
	profileService "Alive/ToDoGRPCv2/service/profile"
	todoService "Alive/ToDoGRPCv2/service/todo"
)

func main() {

	todoService.DialToServiceServer(8000)
	profileService.DialToServiceServer(8000)

	g := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	g.Use(sessions.Sessions("mysession", store))

	g.POST("/profile/login", profileService.LoginProfile)
	g.POST("/profile", profileService.CreateProfile)

	authorized := g.Group("/")
	authorized.Use(middleware.AuthorizationRequired)
	{
		authorized.GET("/todo", todoService.ReadAllToDo)
		authorized.GET("/todo/:todoid", todoService.ReadToDo)
		authorized.POST("/todo", todoService.CreateToDo)
		authorized.PUT("/todo/:todoid", todoService.UpdateToDo)
		authorized.DELETE("/todo/:todoid", todoService.DeleteToDo)

		authorized.GET("/profile", profileService.ReadProfile)
		authorized.PUT("/profile", profileService.UpdateProfile)
		authorized.DELETE("/profile", profileService.DeleteProfile)
		authorized.POST("/profile/logout", profileService.LogoutProfile)
	}

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}

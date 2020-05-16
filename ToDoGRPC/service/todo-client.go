package service

import (
	proto "Alive/ToDoGRPC/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func dialToServiceServer(port int) proto.ToDoServiceClient {

	conn, err := grpc.Dial("localhost:"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewToDoServiceClient(conn)

	return client
}

func ReadAllToDo(ctx *gin.Context) {

}

func ReadToDo(ctx *gin.Context) {

	clientService := dialToServiceServer(8000)

	userId, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		log.Fatal(err)
	}

	toDoId, err := strconv.Atoi(ctx.Params("todoid"))
	if err != nil {
		log.Fatal(err)
	}

	request := &proto.ReadRequest{Api: "v1", Id: int64(toDoId), UserId: int64(userId)}

	if response, err := clientService.Read(ctx, request); err == nil {
		ctx.JSON(http.StatusOK, ctx.BindJSON(response.ToDo))
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	// gin.H{
	// 	"id":          response.ToDo.Id,
	// 	"title":       response.ToDo.Title,
	// 	"description": response.ToDo.Description,
	// 	"status":      response.ToDo.Status,
	// 	"tag":         response.ToDo.Tag,
	// 	"createddate": response.ToDo.CreatedDate,
}

func CreateToDo(ctx *gin.Context) {

}

func UpdateToDo(ctx *gin.Context) {

}

func DeleteToDo(ctx *gin.Context) {

}

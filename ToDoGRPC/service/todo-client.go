package service

import (
	proto "Alive/ToDoGRPC/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	clientApiVersion = "v1"
)

var clientService proto.ToDoServiceClient

func DialToServiceServer(port int) {

	conn, err := grpc.Dial("localhost:"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	clientService = proto.NewToDoServiceClient(conn)
}

func ReadAllToDo(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("userid"))
	if err != nil {
		log.Fatal(err)
	}

	request := &proto.ReadAllRequest{Api: clientApiVersion, UserId: int64(userId)}

	if response, err := clientService.ReadAll(ctx, request); err == nil {
		for _, item := range response.ToDos {
			ctx.JSON(200, item)
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}

func ReadToDo(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("userid"))
	if err != nil {
		log.Fatal(err)
	}

	toDoId, err := strconv.Atoi(ctx.Param("todoid"))
	if err != nil {
		log.Fatal(err)
	}

	request := &proto.ReadRequest{Api: clientApiVersion, Id: int64(toDoId), UserId: int64(userId)}

	if response, err := clientService.Read(ctx, request); err == nil {
		ctx.JSON(200, response)
	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

}

func CreateToDo(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("userid"))
	if err != nil {
		log.Fatal(err)
	}

	toDo := proto.ToDoRequest{}
	ctx.BindJSON(&toDo)
	request := &proto.CreateRequest{
		Api: clientApiVersion,
		ToDo: &proto.ToDoRequest{
			Title:       toDo.Title,
			Description: toDo.Description,
			Status:      toDo.Status,
			Tag:         toDo.Tag,
		},
		UserId: int64(userId),
	}

	if response, err := clientService.Create(ctx, request); err == nil {
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func UpdateToDo(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("userid"))
	if err != nil {
		log.Fatal(err)
	}

	toDoId, err := strconv.Atoi(ctx.Param("todoid"))
	if err != nil {
		log.Fatal(err)
	}

	toDo := proto.ToDoRequest{}
	ctx.BindJSON(&toDo)
	request := &proto.UpdateRequest{
		Api: clientApiVersion,
		Id:  int64(toDoId),
		ToDo: &proto.ToDoRequest{
			Title:       toDo.Title,
			Description: toDo.Description,
			Status:      toDo.Status,
			Tag:         toDo.Tag,
		},
		UserId: int64(userId),
	}

	if response, err := clientService.Update(ctx, request); err == nil {
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func DeleteToDo(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("userid"))
	if err != nil {
		log.Fatal(err)
	}

	toDoId, err := strconv.Atoi(ctx.Param("todoid"))
	if err != nil {
		log.Fatal(err)
	}

	request := &proto.DeleteRequest{Api: clientApiVersion, Id: int64(toDoId), UserId: int64(userId)}

	if response, err := clientService.Delete(ctx, request); err == nil {
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

package todo

import (
	proto "Alive/ToDoGRPCv2/proto/todo"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	clientApiVersion = "v2"
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

	session := sessions.Default(ctx)
	profileId := session.Get("profileId")

	tag := ctx.Query("tag")
	if tag == "" {
		tag = "%"
	}

	request := &proto.ReadAllRequest{
		Api:       clientApiVersion,
		ProfileId: profileId.(int64),
		Tag:       tag,
	}

	if response, err := clientService.ReadAll(ctx, request); err == nil {
		for _, item := range response.ToDos {
			ctx.JSON(200, item)
		}
	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func ReadToDo(ctx *gin.Context) {

	session := sessions.Default(ctx)
	profileId := session.Get("profileId")

	toDoId, err := strconv.Atoi(ctx.Param("todoid"))
	if err != nil {
		log.Fatal(err)
	}

	request := &proto.ReadRequest{
		Api:       clientApiVersion,
		Id:        int64(toDoId),
		ProfileId: profileId.(int64),
	}

	if response, err := clientService.Read(ctx, request); err == nil {
		ctx.JSON(200, response)
	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

}

func CreateToDo(ctx *gin.Context) {

	session := sessions.Default(ctx)
	profileId := session.Get("profileId")

	toDo := proto.ToDoRequest{}
	ctx.BindJSON(&toDo)
	request := &proto.CreateRequest{
		Api:       clientApiVersion,
		ToDo:      &toDo,
		ProfileId: profileId.(int64),
	}

	if response, err := clientService.Create(ctx, request); err == nil {
		ctx.JSON(200, response)
	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func UpdateToDo(ctx *gin.Context) {

	session := sessions.Default(ctx)
	profileId := session.Get("profileId")

	toDoId, err := strconv.Atoi(ctx.Param("todoid"))
	if err != nil {
		log.Fatal(err)
	}

	var toDo proto.ToDoRequest
	ctx.BindJSON(&toDo)
	log.Fatal(toDo)
	request := &proto.UpdateRequest{
		Api:       clientApiVersion,
		Id:        int64(toDoId),
		ToDo:      &toDo,
		ProfileId: profileId.(int64),
	}

	if response, err := clientService.Update(ctx, request); err == nil {
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func DeleteToDo(ctx *gin.Context) {

	session := sessions.Default(ctx)
	profileId := session.Get("profileId")

	toDoId, err := strconv.Atoi(ctx.Param("todoid"))
	if err != nil {
		log.Fatal(err)
	}

	request := &proto.DeleteRequest{
		Api:       clientApiVersion,
		Id:        int64(toDoId),
		ProfileId: profileId.(int64),
	}

	if response, err := clientService.Delete(ctx, request); err == nil {
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

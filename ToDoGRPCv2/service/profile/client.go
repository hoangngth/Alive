package profile

import (
	proto "Alive/ToDoGRPCv2/proto/profile"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	clientApiVersion = "v2"
)

var clientService proto.ProfileServiceClient

func DialToServiceServer(port int) {

	conn, err := grpc.Dial("localhost:"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	clientService = proto.NewProfileServiceClient(conn)
}

func LoginProfile(ctx *gin.Context) {

	session := sessions.Default(ctx)

	var login proto.LoginRequest
	ctx.BindJSON(&login)

	request := &login

	if response, err := clientService.Login(ctx, request); err == nil {
		session.Set("auth", true)
		session.Set("profileId", response.GetId())
		session.Save()
		ctx.JSON(200, response)

	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func LogoutProfile(ctx *gin.Context) {

	session := sessions.Default(ctx)
	session.Set("auth", false)
	session.Save()
	profileId := session.Get("profileId")

	request := &proto.LogoutRequest{Api: clientApiVersion, Id: profileId.(int64)}

	if response, err := clientService.Logout(ctx, request); err == nil {
		ctx.JSON(200, response)
	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func ReadProfile(ctx *gin.Context) {

	session := sessions.Default(ctx)
	profileId := session.Get("profileId")

	request := &proto.ReadRequest{Api: clientApiVersion, Id: profileId.(int64)}

	if response, err := clientService.Read(ctx, request); err == nil {
		ctx.JSON(200, response)
	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func CreateProfile(ctx *gin.Context) {

	var profile proto.ProfileRequest
	ctx.BindJSON(&profile)

	request := &proto.CreateRequest{
		Api:     clientApiVersion,
		Profile: &profile,
	}

	if response, err := clientService.Create(ctx, request); err == nil {
		ctx.JSON(200, response)
	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func UpdateProfile(ctx *gin.Context) {

	session := sessions.Default(ctx)
	profileId := session.Get("profileId")

	profile := proto.ProfileRequest{}
	ctx.BindJSON(&profile)
	request := &proto.UpdateRequest{
		Api:     clientApiVersion,
		Id:      profileId.(int64),
		Profile: &profile,
	}

	if response, err := clientService.Update(ctx, request); err == nil {
		ctx.JSON(200, response)
	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func DeleteProfile(ctx *gin.Context) {

	session := sessions.Default(ctx)
	profileId := session.Get("profileId")

	request := &proto.DeleteRequest{Api: clientApiVersion, Id: profileId.(int64)}

	if response, err := clientService.Delete(ctx, request); err == nil {
		ctx.JSON(200, response)
	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

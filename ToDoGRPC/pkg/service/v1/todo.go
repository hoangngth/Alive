package v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"Alive/Todo/pkg/service/v1"
)

const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) proto.ToDoServiceServer {
	return &toDoServiceServer{db: db}
}

func (s *toDoServiceServer) checkApiVersion(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(
				codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'",
				apiVersion, api)
		}
	}
	return nil
}

func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

func (s *toDoServiceServer) Create(ctx context.Context, request *v1.CreateRequest) (*v1.CreateResponse, error) {
	return &v1.CreateResponse{Api: api, Id: 2}, nil
}

func (s *toDoServiceServer) Read(ctx context.Context, request *v1.ReadRequest) (*v1.ReadResponse, error) {
	return &v1.ReadResponse{Api: api, ToDo: v1.ToDo}, nil
}

func (s *toDoServiceServer) Update(ctx context.Context, request *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	return &v1.CreateResponse{Api: api, Updated: 0}, nil
}

func (s *toDoServiceServer) Delete(ctx context.Context, request *v1.CreateRequest) (*v1.CreateResponse, error) {
	return &v1.CreateResponse{Api: api, Deleted: 0}, nil
}

func (s *toDoServiceServer) ReadAll(ctx context.Context, request *v1.CreateRequest) (*v1.CreateResponse, error) {
	return &v1.CreateResponse{Api: api, ToDo: []*v1.ToDo}, nil
}


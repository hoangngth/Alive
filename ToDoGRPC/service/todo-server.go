package service

import (
	proto "Alive/ToDoGRPC/proto"
	"context"
	"database/sql"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type server struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) proto.ToDoServiceServer {
	return &server{db: db}
}

func (s *server) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Get SQL connection from pool
func (s *server) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

func (s *server) ReadAllResponse(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	return nil, nil
}

func (s *server) ReadResponse(ctx context.Context, request *proto.ReadRequest) (*proto.ReadResponse, error) {

	if err := s.checkAPI(request.Api); err != nil {
		return nil, err
	}

	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT * FROM Todo WHERE ID = ? AND UserId = ?",
		request.Id, request.UserId)

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to select from ToDo-> "+err.Error())
	}
	defer rows.Close()

	var toDo proto.ToDoResponse
	for rows.Next() {
		err = rows.Scan(&toDo.Id, &toDo.Title, &toDo.Description, &toDo.Status, &toDo.Tag, &toDo.CreatedDate)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &proto.ReadResponse{Api: apiVersion, ToDo: &toDo}, nil
}

func (s *server) CreateResponse(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	return nil, nil
}

func (s *server) UpdateResponse(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}

func (s *server) DeleteResponse(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	return nil, nil
}

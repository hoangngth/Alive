package service

import (
	proto "Alive/ToDoGRPC/proto"
	"context"
	"database/sql"
	"log"

	"github.com/golang/protobuf/ptypes"

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

func (s *server) ReadAll(ctx context.Context, request *proto.ReadAllRequest) (*proto.ReadAllResponse, error) {
	return nil, nil
}

func (s *server) Read(ctx context.Context, request *proto.ReadRequest) (*proto.ReadResponse, error) {

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
		createdDate, err := ptypes.Timestamp(toDo.CreatedDate)
		err = rows.Scan(&toDo.Id, &toDo.UserId, &toDo.Title, &toDo.Description, &toDo.Status, &toDo.Tag, &createdDate)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &proto.ReadResponse{Api: apiVersion, ToDo: &toDo}, nil
}

func (s *server) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	return nil, nil
}

func (s *server) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	return nil, nil
}

func (s *server) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}

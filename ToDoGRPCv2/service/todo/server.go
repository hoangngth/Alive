package todo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "Alive/ToDoGRPCv2/proto/todo"
)

const (
	serverApiVersion = "v2"
)

type server struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) proto.ToDoServiceServer {
	return &server{db: db}
}

func (s *server) checkAPI(api string) error {
	if len(api) > 0 {
		if serverApiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", serverApiVersion, api)
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

	if err := s.checkAPI(request.GetApi()); err != nil {
		return nil, err
	}

	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT * FROM Profile_Todo WHERE ProfileId = ? AND Tag LIKE ?",
		request.GetProfileId(), request.GetTag())

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to select from Profile_ToDo-> "+err.Error())
	}
	defer rows.Close()

	var toDos []*proto.ToDoResponse
	for rows.Next() {
		toDo := new(proto.ToDoResponse)
		var createdDate time.Time
		if err != nil {
			log.Println(err)
		}
		err = rows.Scan(&toDo.Id, &toDo.ProfileId, &toDo.Title, &toDo.Description, &toDo.Status, &toDo.Tag, &createdDate)
		if err != nil {
			return nil, status.Error(codes.Unknown, fmt.Sprintf("Failed to retrieve field value"))
		} else {
			toDo.CreatedDate, _ = ptypes.TimestampProto(createdDate)
			toDos = append(toDos, toDo)
		}
	}

	return &proto.ReadAllResponse{Api: serverApiVersion, ToDos: toDos}, nil
}

func (s *server) Read(ctx context.Context, request *proto.ReadRequest) (*proto.ReadResponse, error) {

	if err := s.checkAPI(request.GetApi()); err != nil {
		return nil, err
	}

	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT * FROM Profile_Todo WHERE ID = ? AND ProfileId = ?",
		request.GetId(), request.GetProfileId())

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to select from Profile_ToDo-> "+err.Error())
	}
	defer rows.Close()

	var toDo proto.ToDoResponse
	var createdDate time.Time
	for rows.Next() {
		toDo.CreatedDate, _ = ptypes.TimestampProto(createdDate)
		err = rows.Scan(&toDo.Id, &toDo.ProfileId, &toDo.Title, &toDo.Description, &toDo.Status, &toDo.Tag, &createdDate)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &proto.ReadResponse{Api: serverApiVersion, ToDo: &toDo}, nil
}

func (s *server) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {

	if err := s.checkAPI(request.GetApi()); err != nil {
		return nil, err
	}

	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	res, err := conn.ExecContext(ctx, "INSERT INTO Profile_Todo (ProfileId, Title, Description, Status, Tag) VALUES (?, ?, ?, ?, ?)",
		request.GetProfileId(), request.GetToDo().GetTitle(), request.GetToDo().GetDescription(),
		request.GetToDo().GetStatus(), request.GetToDo().GetTag())

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to insert into Profile_ToDo-> "+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to retrieve id for created Profile_ToDo-> "+err.Error())
	}
	fmt.Printf("ToDo with ID='%d' is created", id)

	return &proto.CreateResponse{Api: serverApiVersion, Created: true}, nil
}

func (s *server) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {

	if err := s.checkAPI(request.GetApi()); err != nil {
		return nil, err
	}

	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	res, err := conn.ExecContext(ctx, "UPDATE Profile_Todo SET Title = ?, Description = ?, Status = ?, Tag = ? WHERE Id = ? AND ProfileId = ?",
		request.GetToDo().GetTitle(), request.GetToDo().GetDescription(), request.GetToDo().GetStatus(),
		request.GetToDo().GetTag(), request.GetId(), request.GetProfileId())

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to update Profile_ToDo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found or unchanged", request.GetId()))
	}

	return &proto.UpdateResponse{Api: serverApiVersion, Updated: true}, nil
}

func (s *server) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {

	if err := s.checkAPI(request.GetApi()); err != nil {
		return nil, err
	}

	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	res, err := conn.ExecContext(ctx, "DELETE FROM Profile_Todo WHERE Id = ? AND ProfileId = ?",
		request.GetId(), request.GetProfileId())

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to delete Profile_ToDo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found", request.GetId()))
	}

	return &proto.DeleteResponse{Api: serverApiVersion, Deleted: true}, nil
}

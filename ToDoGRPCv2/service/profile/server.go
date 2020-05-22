package profile

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "Alive/ToDoGRPCv2/proto/profile"
)

const (
	serverApiVersion = "v2"
)

type server struct {
	db *sql.DB
}

func NewProfileServiceServer(db *sql.DB) proto.ProfileServiceServer {
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

func (s *server) Read(ctx context.Context, request *proto.ReadRequest) (*proto.ReadResponse, error) {

	if err := s.checkAPI(request.GetApi()); err != nil {
		return nil, err
	}

	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT Id, Name, Email, Phone, Address, CreatedDate FROM Profile WHERE Id = ?", request.GetId())

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to select from Profile-> "+err.Error())
	}
	defer rows.Close()

	var profile proto.ProfileResponse
	var strCreateDate string
	formatTime := "2006-01-02 15:04:05"
	for rows.Next() {
		err = rows.Scan(&profile.Id, &profile.Name, &profile.Email, &profile.Phone, &profile.Address, &strCreateDate)
		if err != nil {
			log.Fatal(err)
		}
		createdDate, _ := time.Parse(formatTime, strCreateDate)
		profile.CreatedDate, _ = ptypes.TimestampProto(createdDate)
	}

	return &proto.ReadResponse{Api: serverApiVersion, Profile: &profile}, nil
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

	res, err := conn.ExecContext(ctx, "INSERT INTO Profile (Username, Password, Name, Email, Phone, Address) VALUES (?, ?, ?, ?, ?, ?)",
		request.GetProfile().GetUsername(), request.GetProfile().GetPassword(), request.GetProfile().GetName(),
		request.GetProfile().GetEmail(), request.GetProfile().GetPhone(), request.GetProfile().GetAddress())

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to insert into Profile-> "+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to retrieve id for created Profile-> "+err.Error())
	}

	return &proto.CreateResponse{Api: serverApiVersion, Id: id}, nil
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

	res, err := conn.ExecContext(ctx, "UPDATE Profile SET Name = ?, Email = ?, Phone = ?, Address = ? WHERE  Id = ?",
		request.GetProfile().GetName(), request.GetProfile().GetEmail(), request.GetProfile().GetPhone(), request.GetProfile().GetAddress(), request.GetId())

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to update Profile-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Profile with ID='%d' is not found or unchanged", request.GetId()))
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

	res, err := conn.ExecContext(ctx, "DELETE FROM Profile WHERE Id = ?",
		request.GetId())

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to delete Profile-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Profile with ID='%d' is not found", request.GetId()))
	}

	return &proto.DeleteResponse{Api: serverApiVersion, Deleted: true}, nil
}

func (s *server) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {

	if err := s.checkAPI(request.GetApi()); err != nil {
		return nil, err
	}

	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT Id FROM Profile WHERE Username = ? AND Password = ?",
		request.GetUsername(), request.GetPassword())

	if err != nil {
		return nil, status.Error(codes.Unknown, "Failed to select from Profile-> "+err.Error())
	}
	defer rows.Close()

	var loggedInId int
	for rows.Next() {
		err = rows.Scan(&loggedInId)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &proto.LoginResponse{Api: serverApiVersion, Id: int64(loggedInId), LoggedIn: true}, nil
}

func (s *server) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.LogoutResponse, error) {

	if err := s.checkAPI(request.GetApi()); err != nil {
		return nil, err
	}

	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return &proto.LogoutResponse{Api: serverApiVersion, LoggedOut: true}, nil
}

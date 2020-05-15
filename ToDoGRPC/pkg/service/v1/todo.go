import (
	"context"
	"database/sql"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	proto "Alive/ToDoGRPC/pkg/proto/v1"
	v1 "Alive/ToDoGRPC/pkg/proto/v1"
)

const (
	apiVersion = "v1"
)

type server struct {
	db *sql.DB
}

func main() {
	
}

func NewToDoServiceServer(db *sql.DB) v1.ToDoServiceServer {
	return &server{db: db}
}

func (s *toDoServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *server) ReadResponse(ctx context.Context, request *proto.ReadRequest) (*proto.ReadResponse, error) {

	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	return
}

func (s *server) ReadAllResponse(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	return
}

func (s *server) CreateResponse(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	return
}

func (s *server) UpdateResponse(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	return
}

func (s *server) DeleteResponse(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	return
}

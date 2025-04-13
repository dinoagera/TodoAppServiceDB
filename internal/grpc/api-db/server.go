package apidb

import (
	"context"
	"fmt"

	"github.com/dinoagera/api-db/internal/domain/models"
	pb "github.com/dinoagera/proto/gen/go/myservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkDB interface {
	CreateTask(ctx context.Context, title string, description string) (int64, error)
	DeleteTask(ctx context.Context, title string) error
	DoneTask(ctx context.Context, title string) error
	GetAllTask(ctx context.Context) ([]models.Task, error)
}
type serverAPI struct {
	pb.UnimplementedDBWorkServer
	workdb WorkDB
}

func Register(gRPC *grpc.Server, workdb WorkDB) {
	pb.RegisterDBWorkServer(gRPC, &serverAPI{workdb: workdb})
}
func (s *serverAPI) CreateTask(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	if req.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "title is empty")
	}
	id, err := s.workdb.CreateTask(ctx, req.Title, req.Description)
	if err != nil {
		return &pb.CreateResponse{
			Id:    0,
			Reply: "error while creating task",
		}, fmt.Errorf("creat to failed")
	}
	return &pb.CreateResponse{
		Id:    id,
		Reply: "create task is successfully",
	}, nil
}
func (s *serverAPI) DeleteTask(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	if req.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "title is empty")
	}
	err := s.workdb.DeleteTask(ctx, req.Title)
	if err != nil {
		return &pb.DeleteResponse{
			Reply: "task was not deleted",
		}, fmt.Errorf("task was not deleted")
	}
	return &pb.DeleteResponse{
		Reply: "task deleted successfully",
	}, nil
}
func (s *serverAPI) DoneTask(ctx context.Context, req *pb.DoneRequest) (*pb.DoneResponse, error) {
	if req.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "title is empty")
	}
	err := s.workdb.DoneTask(ctx, req.Title)
	if err != nil {
		return &pb.DoneResponse{
			Reply: "task is not done",
		}, fmt.Errorf("task is not done")
	}
	return &pb.DoneResponse{
		Reply: "task done successfully",
	}, nil
}
func (s *serverAPI) GetAllTask(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	tasks, err := s.workdb.GetAllTask(ctx)
	if err != nil {
		return &pb.GetAllResponse{
			Tasks: nil,
		}, fmt.Errorf("get all failed")
	}
	pbTasks := make([]*pb.Task, 0, len(tasks))
	for _, task := range tasks {
		pbTasks = append(pbTasks, &pb.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
		})
	}
	return &pb.GetAllResponse{
		Tasks: pbTasks,
	}, nil
}

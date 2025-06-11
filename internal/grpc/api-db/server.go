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
	CreateTask(ctx context.Context, title string, description string, uid int64) (int64, error)
	DeleteTask(ctx context.Context, id int64, uid int64) error
	DoneTask(ctx context.Context, id int64, uid int64) error
	GetAllTask(ctx context.Context, uid int64) ([]models.Task, error)
	ChangeTask(ctx context.Context, id int64, title string, description string, uid int64) error
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
	id, err := s.workdb.CreateTask(ctx, req.Title, req.Description, req.Userid)
	if err != nil {
		return &pb.CreateResponse{
			Id:      0,
			Message: "error while creating task",
		}, fmt.Errorf("creat to failed")
	}
	return &pb.CreateResponse{
		Id:      id,
		Message: "create task is successfully",
	}, nil
}
func (s *serverAPI) DeleteTask(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Id is correctly")
	}
	err := s.workdb.DeleteTask(ctx, req.Id, req.Userid)
	if err != nil {
		return &pb.DeleteResponse{
			Message: "task was not deleted",
		}, fmt.Errorf("task was not deleted")
	}
	return &pb.DeleteResponse{
		Message: "task deleted successfully",
	}, nil
}
func (s *serverAPI) DoneTask(ctx context.Context, req *pb.DoneRequest) (*pb.DoneResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Id is correctly")
	}
	err := s.workdb.DoneTask(ctx, req.Id, req.Userid)
	if err != nil {
		return &pb.DoneResponse{
			Message: "task is not done",
		}, fmt.Errorf("task is not done")
	}
	return &pb.DoneResponse{
		Message: "task done successfully",
	}, nil
}
func (s *serverAPI) GetAllTask(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	tasks, err := s.workdb.GetAllTask(ctx, req.Userid)
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
			Uid:         task.ID,
		})
	}
	return &pb.GetAllResponse{
		Tasks: pbTasks,
	}, nil
}
func (s *serverAPI) ChangeTask(ctx context.Context, req *pb.ChangeRequest) (*pb.ChangeResponse, error) {
	err := s.workdb.ChangeTask(ctx, req.Id, req.Title, req.Description, req.Userid)
	if err != nil {
		return &pb.ChangeResponse{
			Message: "Task is not changed",
		}, fmt.Errorf("failed to changed task")
	}
	return &pb.ChangeResponse{
		Message: "Task is changed",
	}, nil
}

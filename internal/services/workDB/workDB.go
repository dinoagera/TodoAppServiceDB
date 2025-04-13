package workdb

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/dinoagera/api-db/internal/domain/models"
)

type WorkDB struct {
	log        *slog.Logger
	createTask TaskCreate
	deleteTask TaskDelete
	donetask   TaskDone
	getalltask TaskGetAll
}
type TaskCreate interface {
	CreateTask(ctx context.Context, title string, description string) (int64, error)
}
type TaskDelete interface {
	DeleteTask(ctx context.Context, title string) error
}
type TaskDone interface {
	DoneTask(ctx context.Context, title string) error
}
type TaskGetAll interface {
	GetAllTask(ctx context.Context) ([]models.Task, error)
}

func New(log *slog.Logger,
	taskcreate TaskCreate,
	taskdelete TaskDelete,
	taskdone TaskDone,
	taskgetall TaskGetAll,
) *WorkDB {
	return &WorkDB{
		log:        log,
		createTask: taskcreate,
		deleteTask: taskdelete,
		donetask:   taskdone,
		getalltask: taskgetall,
	}
}

func (w *WorkDB) CreateTask(ctx context.Context, title string, description string) (int64, error) {
	id, err := w.createTask.CreateTask(ctx, title, description)
	if err != nil {
		w.log.Info("create task to failed")
		return 0, fmt.Errorf("created task to failed, err:%s", err.Error())
	}
	return id, nil

}
func (w *WorkDB) DeleteTask(ctx context.Context, title string) error {
	err := w.deleteTask.DeleteTask(ctx, title)
	if err != nil {
		w.log.Info("delete task to failed")
		return fmt.Errorf("delete task to failed, err:%s", err.Error())
	}
	return nil
}
func (w *WorkDB) DoneTask(ctx context.Context, title string) error {
	err := w.donetask.DoneTask(ctx, title)
	if err != nil {
		w.log.Info("delete task to failed")
		return fmt.Errorf("created task to failed, err:%s", err.Error())
	}
	return nil
}
func (w *WorkDB) GetAllTask(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	tasks, err := w.getalltask.GetAllTask(ctx)
	if err != nil {
		w.log.Info("getall tasks to failed")
		return nil, fmt.Errorf("getall tasks to failed, err:%s", err.Error())
	}
	return tasks, nil
}

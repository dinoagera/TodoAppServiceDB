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
	changetask TaskChange
}
type TaskCreate interface {
	CreateTask(ctx context.Context, title string, description string, uid int64) (int64, error)
}
type TaskDelete interface {
	DeleteTask(ctx context.Context, id int64, uid int64) error
}
type TaskDone interface {
	DoneTask(ctx context.Context, id int64, uid int64) error
}
type TaskGetAll interface {
	GetAllTask(ctx context.Context, uid int64) ([]models.Task, error)
}
type TaskChange interface {
	ChangeTask(ctx context.Context, id int64, title string, description string, uid int64) error
}

func New(log *slog.Logger,
	taskcreate TaskCreate,
	taskdelete TaskDelete,
	taskdone TaskDone,
	taskgetall TaskGetAll,
	taskchange TaskChange,
) *WorkDB {
	return &WorkDB{
		log:        log,
		createTask: taskcreate,
		deleteTask: taskdelete,
		donetask:   taskdone,
		getalltask: taskgetall,
		changetask: taskchange,
	}
}

func (w *WorkDB) CreateTask(ctx context.Context, title string, description string, uid int64) (int64, error) {
	id, err := w.createTask.CreateTask(ctx, title, description, uid)
	if err != nil {
		w.log.Info("create task to failed")
		return 0, fmt.Errorf("created task to failed, err:%s", err.Error())
	}
	return id, nil

}
func (w *WorkDB) DeleteTask(ctx context.Context, id int64, uid int64) error {
	err := w.deleteTask.DeleteTask(ctx, id, uid)
	if err != nil {
		w.log.Info("delete task to failed")
		return fmt.Errorf("delete task to failed, err:%s", err.Error())
	}
	return nil
}
func (w *WorkDB) DoneTask(ctx context.Context, id int64, uid int64) error {
	err := w.donetask.DoneTask(ctx, id, uid)
	if err != nil {
		w.log.Info("delete task to failed")
		return fmt.Errorf("created task to failed, err:%s", err.Error())
	}
	return nil
}
func (w *WorkDB) GetAllTask(ctx context.Context, uid int64) ([]models.Task, error) {
	var tasks []models.Task
	tasks, err := w.getalltask.GetAllTask(ctx, uid)
	if err != nil {
		w.log.Info("getall tasks to failed")
		return nil, fmt.Errorf("getall tasks to failed, err:%s", err.Error())
	}
	return tasks, nil
}
func (w *WorkDB) ChangeTask(ctx context.Context, id int64, title string, description string, uid int64) error {
	err := w.changetask.ChangeTask(ctx, id, title, description, uid)
	if err != nil {
		w.log.Info("change task to failed")
		return fmt.Errorf("changed task to failed")
	}
	return nil
}

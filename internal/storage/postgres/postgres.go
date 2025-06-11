package postgres

import (
	"context"
	"fmt"

	"github.com/dinoagera/api-db/internal/domain/models"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}

func New(storagePath string) (*Storage, error) {
	conn, err := pgx.Connect(context.Background(), storagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("db ping failed: %w", err)
	}

	return &Storage{db: conn}, nil
}

func (s *Storage) CreateTask(ctx context.Context, title string, description string, uid int64) (int64, error) {
	var id int64
	err := s.db.QueryRow(ctx,
		"INSERT INTO tasks(title, description, uid) VALUES($1, $2, $3) RETURNING id",
		title, description, uid).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %w", err)
	}
	return id, nil
}

func (s *Storage) DeleteTask(ctx context.Context, id int64, uid int64) error {
	result, err := s.db.Exec(ctx, "DELETE FROM tasks WHERE id=$1 AND uid=$2", id, uid)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("task not found or access denied")
	}
	return nil
}

func (s *Storage) DoneTask(ctx context.Context, id int64, uid int64) error {
	result, err := s.db.Exec(ctx,
		"UPDATE tasks SET done = true WHERE id = $1 AND done = false AND uid=$2",
		id, uid)
	if err != nil {
		return fmt.Errorf("failed to mark task as done: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("task not found or access denied")
	}
	return nil
}

func (s *Storage) GetAllTask(ctx context.Context, uid int64) ([]models.Task, error) {
	rows, err := s.db.Query(ctx,
		"SELECT id, title, description, done FROM tasks WHERE uid=$1", uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Done,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tasks, nil
}
func (s *Storage) ChangeTask(ctx context.Context, id int64, title string, description string, uid int64) error {
	result, err := s.db.Exec(ctx,
		"UPDATE tasks SET title = $1, description=$2 WHERE id = $3 AND uid=$4",
		title, description, id, uid)
	if err != nil {
		return fmt.Errorf("change task is failed")
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("task not found or access denied")
	}
	return nil
}

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

func (s *Storage) CreateTask(ctx context.Context, title string, description string) (int64, error) {
	var id int64
	err := s.db.QueryRow(ctx,
		"INSERT INTO tasks(title, description) VALUES($1, $2) RETURNING id",
		title, description).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %w", err)
	}
	return id, nil
}

func (s *Storage) DeleteTask(ctx context.Context, title string) error {
	_, err := s.db.Exec(ctx, "DELETE FROM tasks WHERE title=$1", title)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (s *Storage) DoneTask(ctx context.Context, title string) error {
	_, err := s.db.Exec(ctx,
		"UPDATE tasks SET done = true WHERE title = $1 AND done = false",
		title)
	if err != nil {
		return fmt.Errorf("failed to mark task as done: %w", err)
	}
	return nil
}

func (s *Storage) GetAllTask(ctx context.Context) ([]models.Task, error) {
	rows, err := s.db.Query(ctx,
		"SELECT id, title, description, done FROM tasks")
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

package models

type Task struct {
	ID          int64
	Title       string
	Description string
	Done        bool
	Uid         int64
}

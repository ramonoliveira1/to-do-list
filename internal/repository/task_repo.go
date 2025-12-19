package repository

import (
	"database/sql"
	"fmt"
	"to-do-list/internal/entity"
)

type TaskRepository interface {
	Create(task *entity.Task) error
	FindById(id int) (entity.Task, error)
	FindAll() ([]entity.Task, error)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Create(task *entity.Task) error {
	query := `INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, task.Title, task.Description, task.Completed)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = id
	return nil
}

func (r *taskRepository) FindById(id int) (entity.Task, error) {
	query := `SELECT id, title, description, completed FROM tasks WHERE id = ?`

	var task entity.Task
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Completed)

	if err == sql.ErrNoRows {
		return entity.Task{}, fmt.Errorf("task not found")
	}
	if err != nil {
		return entity.Task{}, fmt.Errorf("failed to find task: %w", err)
	}

	return task, nil
}

func (r *taskRepository) FindAll() ([]entity.Task, error) {
	query := `SELECT id, title, description, completed FROM tasks`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return tasks, nil
}

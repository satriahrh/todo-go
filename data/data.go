package data

import (
	"context"
	"time"
)

// BlueprintOfDatabase interface of data and its derivatives
type BlueprintOfDatabase interface {
	FetchProjects(ctx context.Context) (projects []Project, err error)
	FetchProjectByID(ctx context.Context, id uint64) (project Project, err error)
	FetchTasksByProjectIDAndTaskType(ctx context.Context, projectID uint64, taskType string) (tasks []Task, err error)

	InsertProject(ctx context.Context, project Project) (err error)
	InsertTask(ctx context.Context, task Task) (err error)

	UpdateProject(ctx context.Context, projectID uint64, updatedProject Project) (err error)
	UpdateTask(ctx context.Context, taskID uint64, updatedTask Task) (err error)

	DeleteProject(ctx context.Context, projectID uint64) (err error)
}

// Project hold project information
type Project struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// Task hold task information
type Task struct {
	ID           uint64    `json:"id"`
	CurrentState string    `json:"task_type"`
	Data         string    `json:"data"`
	TodoSince    time.Time `json:"todo_since"`
	DoneSince    time.Time `json:"done_since"`
	BlockedSince time.Time `json:"blocked_since"`
	ArchivedAt   time.Time `json:"archived_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdateAt     time.Time `json:"update_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	ProjectID    uint64    `json:"project_id"`
}

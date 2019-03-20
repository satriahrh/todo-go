package todo_go

import (
	"context"
	"github.com/satriahrh/todo-go/data"
)

// TodoGo main struct of the project
type TodoGo struct {
	Database data.BlueprintOfDatabase
}

// NewTodoGo constructor like of TodoGo
func NewTodoGo(database data.BlueprintOfDatabase) (todoGo *TodoGo, err error) {
	return &TodoGo{
		Database: database,
	}, nil
}

// BlueprintOfTodoGo main interfaces of the project
type BlueprintOfTodoGo interface {
	GetProjects(ctx context.Context) (projects []data.Project, err error)
	GetProjectByID(ctx context.Context, projectID uint64) (project data.Project, err error)
	GetTodoByProjectID(ctx context.Context, projectID uint64) (tasks []data.Task, err error)
	GetDoneByProjectID(ctx context.Context, projectID uint64) (tasks []data.Task, err error)
	GetBlockByProjectID(ctx context.Context, projectID uint64) (tasks []data.Task, err error)

	CreateProject(ctx context.Context, project data.Project) (projects []data.Project, err error)
	CreateTask(ctx context.Context, task data.Task) (tasks []data.Task, err error)

	UpdateProject(ctx context.Context, projectID uint64, updatedProject data.Project) (project data.Project, err error)
	UpdateTaskState(ctx context.Context, taskID uint64, destination string) (originTasks, destinationTasks []data.Task, err error)
	ArchiveTasks(ctx context.Context, projectID uint64) (err error)

	DeleteProject(ctx context.Context, projectID uint64) (projects []data.Project, err error)
}

// GetProjects get list of projects
func (t *TodoGo) GetProjects(ctx context.Context) (projects []data.Project, err error) {
	projects, err = t.Database.FetchProjects(ctx)
	return
}

// GetProjectByID get project information by ID
func (t *TodoGo) GetProjectByID(ctx context.Context, id uint64) (project data.Project, err error) {
	project, err = t.Database.FetchProjectByID(ctx, id)
	return
}

// GetTodoByProjectID get a list of task of todo by project id
func (t *TodoGo) GetTodoByProjectID(ctx context.Context, projectID uint64) (tasks []data.Task, err error) {
	return
}

// GetDoneByProjectID get a list of task of done by project id
func (t *TodoGo) GetDoneByProjectID(ctx context.Context, projectID uint64) (tasks []data.Task, err error) {
	return
}

// GetBlockByProjectID get a list of task of block by project id
func (t *TodoGo) GetBlockByProjectID(ctx context.Context, projectID uint64) (tasks []data.Task, err error) {
	return
}

// CreateProject create a brand new project
func (t *TodoGo) CreateProject(ctx context.Context, project data.Project) (projects []data.Project, err error) {
	err = t.Database.InsertProject(ctx, project)
	if err != nil {
		return
	}
	projects, err = t.Database.FetchProjects(ctx)
	return
}

// CreateTask create a task on a project
func (t *TodoGo) CreateTask(ctx context.Context, task data.Task) (tasks []data.Task, err error) {
	return
}

// UpdateProject update some attributes on project
func (t *TodoGo) UpdateProject(ctx context.Context, projectID uint64, newProject data.Project) (project data.Project, err error) {
	err = t.Database.UpdateProject(ctx, projectID, newProject)
	if err != nil {
		return
	}

	project, err = t.Database.FetchProjectByID(ctx, projectID)
	return
}

// UpdateTaskState to update task state.
// todo --> done; todo --> blocked; blocked --> todo
func (t *TodoGo) UpdateTaskState(ctx context.Context, taskID uint64, destination string) (originTasks, destinationTasks []data.Task, err error) {
	return
}

// ArchiveTasks to archieve done project, and update blocked into todo
func (t *TodoGo) ArchiveTasks(ctx context.Context, projectID uint64) (err error) {
	return
}

// DeleteProject
func (t *TodoGo) DeleteProject(ctx context.Context, projectID uint64) (projects []data.Project, err error) {
	err = t.Database.DeleteProject(ctx, projectID)
	if err != nil {
		return
	}

	projects, err = t.Database.FetchProjects(ctx)
	return
}

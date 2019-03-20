package mysql

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/satriahrh/todo-go/data"
	"time"
)

// MySql struct of mysql
type MySql struct {
	DB *sql.DB
}

// NewMySql construct MySql
func NewMySql(dataSourceName string) (mySql *MySql, err error) {
	cfg, err := mysql.ParseDSN(dataSourceName)
	if err != nil {
		return
	}
	cfg.ParseTime = true
	cfg.Loc = time.Local
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return
	}
	return &MySql{
		DB: db,
	}, err
}

// FetchProjects fetch a list of projects
func (m *MySql) FetchProjects(ctx context.Context) (projects []data.Project, err error) {
	rows, err := m.DB.QueryContext(ctx,
		"SELECT id, title, created_at, updated_at FROM projects WHERE deleted_at IS NULL;",
	)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var project data.Project
		var updatedAt mysql.NullTime
		if err = rows.Scan(
			&project.ID,
			&project.Title,
			&project.CreatedAt,
			&updatedAt,
		); err != nil {
			return
		}
		project.UpdatedAt = updatedAt.Time
		projects = append(projects, project)
	}
	return
}

// FetchProjectByID fetch a project by its ID
func (m *MySql) FetchProjectByID(ctx context.Context, id uint64) (project data.Project, err error) {
	row := m.DB.QueryRowContext(ctx,
		"SELECT id, title, created_at, updated_at FROM projects WHERE id = ?",
		id,
	)
	var updatedAt mysql.NullTime
	if err = row.Scan(
		&project.ID,
		&project.Title,
		&project.CreatedAt,
		&updatedAt,
	); err != nil {
		return
	}
	project.UpdatedAt = updatedAt.Time
	return
}

// FetchTasksByProjectIDAndTaskType fetch list of task by its project id and task type.
// Useful to fetch todo, done, and block.
func (m *MySql) FetchTasksByProjectIDAndTaskType(ctx context.Context, projectID uint64, taskType string) (tasks []data.Task, err error) {
	return
}

// InsertProject insert new project
func (m *MySql) InsertProject(ctx context.Context, project data.Project) (err error) {
	_, err = m.DB.ExecContext(ctx,
		"INSERT INTO projects (title) values (?);",
		project.Title,
	)
	return
}

// InsertTask insert new task with foreign key of project id
func (m *MySql) InsertTask(ctx context.Context, task data.Task) (err error) {
	return
}

// UpdateProject update partially a project by id with a new value
func (m *MySql) UpdateProject(ctx context.Context, projectID uint64, updatedProject data.Project) (err error) {
	_, err = m.DB.ExecContext(ctx,
		"UPDATE projects SET title = ? WHERE id = ?",
		updatedProject.Title, projectID,
	)
	return
}

// UpdateTask update partially a task by id with a new value
func (m *MySql) UpdateTask(ctx context.Context, taskID uint64, updatedTask data.Task) (err error) {
	return
}

// DeleteProject delete a project by setting deleted_at to now
func (m *MySql) DeleteProject(ctx context.Context, projectID uint64) (err error) {
	_, err = m.DB.ExecContext(ctx,
		"UPDATE projects SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?",
		projectID,
	)
	return
}

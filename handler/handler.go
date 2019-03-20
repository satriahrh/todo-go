package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/satriahrh/todo-go"
	"github.com/satriahrh/todo-go/data"
	"github.com/satriahrh/todo-go/handler/response"
	"net/http"
	"strconv"
)

type Handler struct {
	TodoGo *todo_go.TodoGo
}

func NewHandler(todoGo *todo_go.TodoGo) (handler *Handler, err error) {
	return &Handler{
		TodoGo: todoGo,
	}, err
}

// Ping check application status
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	response.Message(200, w, "Pong")
}

// GetProjects handle getting all projects
func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	var err error
	var projects []data.Project

	if projects, err = h.TodoGo.GetProjects(r.Context()); err != nil {
		response.Error(w, err)
		return
	}
	response.Success(200, w, projects)
}

// GetProjectByID handle getting a project by id
func (h *Handler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	var err error
	var project data.Project

	projectID, _ := strconv.ParseUint(getUrlParams(r, "id"), 10, 64)

	if project, err = h.TodoGo.GetProjectByID(r.Context(), projectID); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(200, w, project)
}

// UpdateProject handle updating some project's attributes
func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var err error
	var project data.Project

	projectID, _ := strconv.ParseUint(getUrlParams(r, "id"), 10, 64)

	if project, err = h.TodoGo.GetProjectByID(r.Context(), projectID); err != nil {
		response.Error(w, err)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&project); err != nil {
		response.Error(w, err)
		return
	}

	if project, err = h.TodoGo.UpdateProject(r.Context(), projectID, project); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(200, w, project)
}

// CreateProject handle project creation
func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var err error
	var project data.Project
	var projects []data.Project

	if err = json.NewDecoder(r.Body).Decode(&project); err != nil {
		response.Error(w, err)
		return
	}

	if projects, err = h.TodoGo.CreateProject(r.Context(), project); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(201, w, projects)
}

// DeleteProject handle project deletion
func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	var err error
	var projects []data.Project

	projectID, _ := strconv.ParseUint(getUrlParams(r, "id"), 10, 64)

	if projects, err = h.TodoGo.DeleteProject(r.Context(), projectID); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(200, w, projects)
}

func getUrlParams(r *http.Request, key string) string {
	return mux.Vars(r)["id"]
}

func getQueryParams(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

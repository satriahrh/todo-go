package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"github.com/satriahrh/todo-go"
	"github.com/satriahrh/todo-go/data/mysql"
	"github.com/satriahrh/todo-go/handler"
	"github.com/subosito/gotenv"
	"net/http"
	"os"
	"path"
)

func main() {
	cwd, _ := os.Getwd()
	err := gotenv.Load(path.Join(
		cwd,
		".env",
	))
	if err != nil {
		log.Errorln(err)
	}

	db, err := mysql.NewMySql(os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatalln(err)
	}

	tg, err := todo_go.NewTodoGo(db)
	if err != nil {
		log.Fatalln(err)
	}

	hdr, err := handler.NewHandler(tg)
	if err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/ping", hdr.Ping).Methods("GET")
	router.HandleFunc("/projects", hdr.CreateProject).Methods("POST")
	router.HandleFunc("/projects", hdr.GetProjects).Methods("GET")
	router.HandleFunc("/projects/{id:[0-9]+}", hdr.GetProjectByID).Methods("GET")
	router.HandleFunc("/projects/{id:[0-9]+}", hdr.UpdateProject).Methods("PATCH")
	router.HandleFunc("/projects/{id:[0-9]+}", hdr.DeleteProject).Methods("DELETE")

	log.Infof("Ready at %v:%v", "", os.Getenv("PORT"))
	log.Fatalln(http.ListenAndServe(
		fmt.Sprintf(":%v", os.Getenv("PORT")),
		router,
	))
}

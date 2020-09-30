package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"todoApp/controller"
	"todoApp/driver"

	"github.com/gorilla/mux"
)

type app struct {
	Db *sql.DB
}

func main() {

	r := mux.NewRouter()

	db := driver.ConnectDB()

	a := app{Db: db}

	r.HandleFunc("/todos", controller.ListTodoHandler(a.Db)).Methods("GET")
	r.HandleFunc("/add", controller.CreateTodoHandler(a.Db)).Methods("POST")
	r.HandleFunc("/delete/id={id}", controller.DeleteTodoHandler(a.Db)).Methods("DELETE")
	r.HandleFunc("/update/id={id}", controller.UpdateTodoHandler(a.Db)).Methods("PUT")

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}

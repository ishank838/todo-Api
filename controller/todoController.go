package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"todoApp/model"
	"todoApp/repository"

	"github.com/gorilla/mux"
)

//UpdateTodoHandler creates a new todo
func UpdateTodoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		id, err := strconv.ParseUint(params["id"], 10, 32)

		if err != nil {
			log.Fatal(err)
		}

		msg, todo := repository.UpdateTodo(db, id)

		if len(msg) == 0 {
			respondwithJSON(w, todo)
		} else {
			respondWithError(w, msg)
		}
	}
}

//DeleteTodoHandler creates a new todo
func DeleteTodoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		id, err := strconv.ParseUint(params["id"], 10, 32)

		if err != nil {
			log.Fatal(err)
		}

		msg := repository.DeleteTodo(db, id)

		respondwithJSON(w, msg)
	}
}

//ListTodoHandler creates a new todo
func ListTodoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var responseTodo []model.TodoResponse

		todos := repository.ListTodo(db)

		for _, t := range *todos {
			curt := getResponseTodo(t)
			responseTodo = append(responseTodo, curt)
		}

		respondwithJSON(w, responseTodo)
	}
}

//CreateTodoHandler creates a new todo
func CreateTodoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var t model.Todo

		err := json.NewDecoder(r.Body).Decode(&t)
		log.Println(t)

		if err == io.EOF {
			respondWithError(w, "Empty Body")
			return
		}

		id, msg := repository.CreateTodo(db, &t)
		t.Status = model.TodoValue

		if id == 0 {
			respondWithError(w, msg)
		} else {
			respondwithJSON(w, t)
		}
	}
}

func respondwithJSON(w http.ResponseWriter, u interface{}) {
	response, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, msg string) {
	respondwithJSON(w, map[string]string{"message": msg})
}

func getResponseTodo(t model.Todo) model.TodoResponse {

	statusVal := getStatusValue(t.Status)

	return model.TodoResponse{
		ID:           t.ID,
		Description:  t.Description,
		Status:       statusVal,
		CompleteDate: t.CompleteDate,
	}
}

func getStatusValue(i uint) string {

	switch i {
	case model.TodoValue:
		return "To DO"
	case model.OngoingValue:
		return "Ongoing"
	case model.CompletedValue:
		return "Completed"
	}

	return ""
}

package repository

import (
	"database/sql"
	"log"
	"todoApp/model"
)

//CreateTodo function for create db operation
func CreateTodo(db *sql.DB, t *model.Todo) (uint, string) {

	var todoID uint

	if t.CompleteDate == "" {
		return 0, "No Complete Date"
	}

	if t.Description == "" {
		return 0, "No Description"
	}

	if len(t.Description) > 15 {
		return 0, "Description Too Long. Length should be atmost 15"
	}

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
		return 0, "System Error"
	}

	err = tx.QueryRow("INSERT INTO todos VALUES (DEFAULT,$1,$2,$3) RETURNING id;",
		t.Description, t.Status, t.CompleteDate).Scan(&todoID)

	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		return todoID, "Database Insert Error"
	}

	tx.Commit()
	return 0, "Success"
}

//UpdateTodo function for create db operation
func UpdateTodo(db *sql.DB, id uint64) (string, *model.Todo) {
	tx, err := db.Begin()

	if err != nil {
		log.Println(err)
		return "Error", nil
	}

	var todo model.Todo

	r := tx.QueryRow("Select * from todos where id=$1;", id)

	err = r.Scan(&todo.ID, &todo.Description, &todo.Status, &todo.CompleteDate)

	if err == sql.ErrNoRows {
		return "ID doesn't Exist", nil
	}

	newStatus := getNextStatus(todo.Status)

	if newStatus == 0 {
		return "Task Already Completed", &todo
	}

	_, err = tx.Exec("UPDATE todos set status=$1 where id=$2;", newStatus, id)

	var newTodo model.Todo

	if err != nil {
		log.Println(err)
		return "Error", nil
	}

	newTodo = model.Todo{
		ID:           todo.ID,
		Description:  todo.Description,
		Status:       newStatus,
		CompleteDate: todo.CompleteDate,
	}

	tx.Commit()
	return "Success", &newTodo
}

//DeleteTodo function for create db operation
func DeleteTodo(db *sql.DB, id uint64) string {

	tx, err := db.Begin()

	if err != nil {
		log.Println(err)
		return "Error"
	}

	var rowCount uint64

	err = tx.QueryRow("Select count(id) from todos where id=$1;", id).Scan(&rowCount)

	if err != nil {
		panic(err)
	}

	if rowCount == 0 {
		return "ID doesn't exist"
	}

	_, err = tx.Exec("DELETE FROM todos WHERE id=$1", id)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return "Item Doesn't exist"
	}

	tx.Commit()
	return "Sccessfully Deleted"
}

//ListTodo function for create db operation
func ListTodo(db *sql.DB) *[]model.Todo {

	var todos []model.Todo

	tx, err := db.Begin()

	if err != nil {
		log.Println(err)
		return nil
	}

	rows, err := tx.Query("SELECT * FROM todos ORDER BY status ASC, enddate ASC;")

	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		var t model.Todo
		rows.Scan(&t.ID, &t.Description, &t.Status, &t.CompleteDate)
		todos = append(todos, t)
	}

	return &todos
}

func getNextStatus(cur uint) uint {

	switch cur {
	case model.TodoValue:
		return model.OngoingValue
	case model.OngoingValue:
		return model.CompletedValue
	default:
		return 0
	}
}

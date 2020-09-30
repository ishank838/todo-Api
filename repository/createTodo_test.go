package repository

import (
	"log"
	"testing"
	"todoApp/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNoCompleteDate(t *testing.T) {
	db, _, _ := sqlmock.New()

	defer db.Close()

	todo := model.Todo{
		Description:  "Todo 1",
		CompleteDate: "",
	}

	id, msg := CreateTodo(db, &todo)

	assert.Equal(t, id, uint(0))
	assert.Equal(t, msg, "No Complete Date")
}

func TestNoDescription(t *testing.T) {
	db, _, _ := sqlmock.New()

	defer db.Close()

	todo := model.Todo{
		Description:  "",
		CompleteDate: "5-10-2020",
	}

	id, msg := CreateTodo(db, &todo)

	assert.Equal(t, id, uint(0))
	assert.Equal(t, msg, "No Description")
}

func TestDescriptionTooLong(t *testing.T) {
	db, _, _ := sqlmock.New()

	defer db.Close()

	todo := model.Todo{
		Description:  "Decription Length is greater than 15",
		CompleteDate: "5-10-2020",
	}

	id, msg := CreateTodo(db, &todo)

	assert.Equal(t, id, uint(0))
	assert.Equal(t, msg, "Description Too Long. Length should be atmost 15")
}

func TestCreateTodo(t *testing.T) {
	db, mock, err := sqlmock.New()

	defer db.Close()

	if err != nil {
		log.Println(err)
	}

	todo := model.Todo{
		Description:  "test Todo",
		Status:       1,
		CompleteDate: "5-10-2020",
	}

	columns := []string{"id"}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO todos VALUES").
		WithArgs(todo.Description, todo.Status, todo.CompleteDate).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("0"))
	mock.ExpectCommit()

	id, msg := CreateTodo(db, &todo)

	assert.Equal(t, id, uint(0))
	assert.Equal(t, msg, "Success")
}

package repository

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTodo(t *testing.T) {
	db, mock, err := sqlmock.New()

	defer db.Close()

	if err != nil {
		log.Println(err)
	}

	columns := []string{"id"}
	mock.ExpectBegin()

	mock.ExpectQuery("Select count(.+) from todos").
		WithArgs(uint64(1)).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	mock.ExpectExec("DELETE FROM todos*").
		WithArgs(uint64(1)).WillReturnResult(sqlmock.NewResult(1, 1))

	msg := DeleteTodo(db, uint64(1))

	assert.Equal(t, msg, "Sccessfully Deleted")
}

func TestDeleteTodoDoesNotExist(t *testing.T) {
	db, mock, err := sqlmock.New()

	defer db.Close()

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	columns := []string{"count"}

	mock.ExpectBegin()
	mock.ExpectQuery("Select count(.+) from todos*").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("0"))

	msg := DeleteTodo(db, 1)

	assert.Equal(t, "ID doesn't exist", msg)
}

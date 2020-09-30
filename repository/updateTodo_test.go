package repository

import (
	"testing"
	"todoApp/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateStatus(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	todo := model.Todo{
		ID:           1,
		Description:  "Todo 1",
		Status:       2,
		CompleteDate: "5-10-2020",
	}

	columns := []string{"id", "description", "status", "enddate"}

	mock.ExpectBegin()

	mock.ExpectQuery("Select . from todos where *").
		WithArgs(uint64(1)).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow("1", "Todo 1", "1", "5-10-2020"))

	mock.ExpectExec("UPDATE todos set *").
		WithArgs(2, 1).WillReturnResult(sqlmock.NewResult(0, 0))

	mock.ExpectCommit()

	msg, todoRes := UpdateTodo(db, 1)

	assert.Equal(t, todo, *todoRes)
	assert.Equal(t, "Success", msg)
}

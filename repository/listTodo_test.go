package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"todoApp/model"
)

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	result := []model.Todo{
		{ID: 4, Description: "Todo 2", Status: 1, CompleteDate: "30-9-2020"},
		{ID: 5, Description: "Todo 3", Status: 1, CompleteDate: "5-10-2020"},
		{ID: 3, Description: "Todo 1", Status: 3, CompleteDate: "29-9-2020"},
	}

	columns := sqlmock.NewRows([]string{"id", "description", "status", "enddate"}).
		AddRow("4", "Todo 2", "1", "30-9-2020").
		AddRow("5", "Todo 3", "1", "5-10-2020").
		AddRow("3", "Todo 1", "3", "29-9-2020")

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT . FROM todos ORDER BY status ASC, enddate ASC;").
		WillReturnRows(columns)

	todos := ListTodo(db)

	for i, a := range *todos {
		assert.Equal(t, result[i], a)
	}
}

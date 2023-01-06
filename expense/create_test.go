//go:build unit

package expense

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {

	exp := Expense{
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	db, mock, _ := sqlmock.New()

	rows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).AddRow(0, "strawberry smoothie", 79, "night market promotion discount 10 bath", pq.Array([]string{"food", "beverage"}))

	mock.ExpectQuery("INSERT INTO expenses ").WithArgs("strawberry smoothie", 79.0, "night market promotion discount 10 bath", pq.Array([]string{"food", "beverage"})).WillReturnRows(rows)

	result, err := CreateExpense(db, exp)
	assert.Nil(t, err)
	assert.EqualValues(t, result.ID, exp.ID)
	assert.EqualValues(t, result.Title, exp.Title)
	assert.EqualValues(t, result.Amount, exp.Amount)
	assert.EqualValues(t, result.Note, exp.Note)
	assert.EqualValues(t, result.Tags, exp.Tags)

}

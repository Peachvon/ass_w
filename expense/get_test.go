//go:build unit

package expense

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetExpense(t *testing.T) {

	exp := Expense{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	db, mock, _ := sqlmock.New()

	row := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).AddRow("1", "strawberry smoothie", 79, "night market promotion discount 10 bath", pq.Array([]string{"food", "beverage"}))

	mock.ExpectPrepare("SELECT id,title, amount,note,tags FROM expenses").ExpectQuery().WillReturnRows(row)

	result, err := GetExpense(db, "1")
	assert.Nil(t, err)
	assert.EqualValues(t, result.ID, exp.ID)
	assert.EqualValues(t, result.Title, exp.Title)
	assert.EqualValues(t, result.Amount, exp.Amount)
	assert.EqualValues(t, result.Note, exp.Note)
	assert.EqualValues(t, result.Tags, exp.Tags)

}

func TestGetAllExpenses(t *testing.T) {
	db, mock, _ := sqlmock.New()
	rows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).AddRow("1", "strawberry smoothie", 79, "night market promotion discount 10 bath", pq.Array([]string{"food", "beverage"})).AddRow("2", "strawberry smoothie", 45, "night market promotion discount 50 bath", pq.Array([]string{"food"}))
	mock.ExpectPrepare("SELECT id,title, amount,note,tags FROM expenses").ExpectQuery().WillReturnRows(rows)
	result, err := GetExpenses(db)
	assert.Nil(t, err)
	assert.EqualValues(t, len(result), 2)

}

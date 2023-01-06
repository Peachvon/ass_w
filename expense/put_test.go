//go:build unit

package expense

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestPutExpense(t *testing.T) {

	exp := Expense{
		ID:     1,
		Title:  "strawberry smoothie",
		Amount: 79,
		Note:   "night market promotion discount 10 bath",
		Tags:   []string{"food", "beverage"},
	}

	db, mock, _ := sqlmock.New()

	r := sqlmock.NewResult(1, 1)

	mock.ExpectPrepare("UPDATE expenses SET").ExpectExec().WithArgs(1, "strawberry smoothie", 79.0, "night market promotion discount 10 bath", pq.Array([]string{"food", "beverage"})).WillReturnResult(r)

	err := PutExpense(db, "1", exp)
	assert.Nil(t, err)

}

// go:build unit
package expense

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTableExpenses(t *testing.T) {
	db, mock, _ := sqlmock.New()

	mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(0, 0))

	err := CreateTableExpenses(db)
	assert.Nil(t, err)

}

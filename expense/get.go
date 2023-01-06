package expense

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func GetExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	exp, err := GetExpense(db, id)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})

	}
	return c.JSON(http.StatusOK, exp)

}
func GetExpense(db *sql.DB, id string) (Expense, error) {
	result := Expense{}
	stmt, err := db.Prepare("SELECT id,title, amount,note,tags FROM expenses where id=$1")
	if err != nil {
		fmt.Println("can't scan expenses", err)
		return result, errors.New("can't scan expenses")
	}
	row := stmt.QueryRow(id)

	err = row.Scan(&result.ID, &result.Title, &result.Amount, &result.Note, pq.Array(&result.Tags))
	if err != nil {
		fmt.Println("can't scan id", err)
		return result, errors.New("can't scan id")
	}
	return result, nil
}

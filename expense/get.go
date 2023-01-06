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

func GetExpensesHandler(c echo.Context) error {
	exps, err := GetExpenses(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, exps)

}
func GetExpenses(db *sql.DB) ([]Expense, error) {
	result := []Expense{}
	stmt, err := db.Prepare("SELECT id,title, amount,note,tags FROM expenses")
	if err != nil {
		fmt.Println("can't query sql", err)

		//c.JSON(http.StatusInternalServerError, Err{Message: "can't query sql" + err.Error()})
		return nil, errors.New("can't query sql")
	}

	rows, err := stmt.Query()
	if err != nil {
		//c.JSON(http.StatusInternalServerError, Err{Message: "can't query sql" + err.Error()})
		fmt.Println("can't query sql", err)

		return nil, errors.New("can't query sql")
	}

	for rows.Next() {
		exp := Expense{}
		err := rows.Scan(&exp.ID, &exp.Title, &exp.Amount, &exp.Note, pq.Array(&exp.Tags))
		if err != nil {
			fmt.Println("can't scan expenses", err)

			// c.JSON(http.StatusInternalServerError, Err{Message: "can't scan expenses:" + err.Error()})
			return nil, errors.New("can't scan expenses")
		}
		result = append(result, exp)
	}
	return result, nil
}

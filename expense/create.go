package expense

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func CreateExpenseHandler(c echo.Context) error {
	var exp Expense
	err := c.Bind(&exp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	result, err := CreateExpense(db, exp)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})

	}
	fmt.Println("insert todo success id : ", exp)
	return c.JSON(http.StatusCreated, result)

}

func CreateExpense(db *sql.DB, exp Expense) (Expense, error) {

	row := db.QueryRow("INSERT INTO expenses (title, amount,note,tags) values ($1, $2,$3,$4) RETURNING id,title, amount,note,tags", exp.Title, exp.Amount, exp.Note, pq.Array(&exp.Tags))

	result := Expense{}
	err := row.Scan(&result.ID, &result.Title, &result.Amount, &result.Note, pq.Array(&result.Tags))
	if err != nil {
		fmt.Println("can't scan id", err)
		return exp, err
	}
	return result, nil
}

type User struct {
	ID   int
	Name string
	Age  int
}

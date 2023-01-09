package expense

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func CheckIdRequest(ParamId, JsonId string) error {

	if ParamId != JsonId {
		return errors.New("ID BadRequest")
	}
	return nil
}

func (h *handler) PutExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	var exp Expense

	err := c.Bind(&exp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	err = CheckIdRequest(id, strconv.Itoa(exp.ID))
	if err != nil {

		return c.JSON(http.StatusBadRequest, Err{Message: "can't query sql: " + err.Error()})

	}

	err = PutExpense(DB, id, exp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	h.GetExpenseHandler(c)
	return nil

}

func PutExpense(db *sql.DB, id string, exp Expense) error {

	stmt, err := db.Prepare("UPDATE expenses SET title=$2 ,amount=$3,note=$4,tags=$5 WHERE id=$1")
	if err != nil {
		return errors.New("can't query sql: " + err.Error())

	}
	row, err := stmt.Exec(exp.ID, exp.Title, exp.Amount, exp.Note, pq.Array(&exp.Tags))
	if err != nil {
		fmt.Println(row)
		return errors.New("can't query sql: " + err.Error())
	}

	return nil

}

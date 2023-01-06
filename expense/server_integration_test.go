//go:build integration

package expense

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
func TestIntegrationCreateExpense(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79.0,
		"note":"night market promotion discount 10 bath",
		"tags":   ["food", "beverage"]
	}`)
	var exp Expense

	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&exp)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, exp.ID)
	assert.Equal(t, "strawberry smoothie", exp.Title)
	assert.Equal(t, 79.0, exp.Amount)
	assert.Equal(t, "night market promotion discount 10 bath", exp.Note)
	assert.Equal(t, []string{"food", "beverage"}, exp.Tags)

}
func TestIntegrationGetExpenseById(t *testing.T) {
	c := seedExpense(t)

	var latest Expense
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(c.ID)), nil)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.ID, latest.ID)
	assert.NotEmpty(t, latest.Tags)
	assert.NotEmpty(t, latest.Amount)
	assert.NotEmpty(t, latest.Note)
	assert.NotEmpty(t, latest.Tags)

}

func seedExpense(t *testing.T) Expense {
	var cexp Expense
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79.0,
		"note":"night market promotion discount 10 bath",
		"tags":   ["food", "beverage"]
	}`)
	err := request(http.MethodPost, uri("expenses"), body).Decode(&cexp)
	if err != nil {
		t.Fatal("can't create expenses:", err)
	}
	return cexp
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	fmt.Println(os.Getenv("AUTH_TOKEN"))
	return &Response{res, err}
}

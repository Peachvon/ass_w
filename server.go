package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type Expenses struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Err struct {
	Message string `json:"message"`
}

func createExpenses(c echo.Context) error {
	var exp Expenses

	err := c.Bind(&exp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	tags := "{" + strings.Join(exp.Tags, ",") + "}"

	//return c.JSON(http.StatusBadRequest, exp)
	row := db.QueryRow("INSERT INTO expenses (title, amount,note,tags) values ($1, $2,$3,$4) RETURNING id", exp.Title, exp.Amount, exp.Note, tags)

	err = row.Scan(&exp.ID)
	if err != nil {
		fmt.Println("can't scan id", err)
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
	}

	fmt.Println("insert todo success id : ", exp)
	return c.JSON(http.StatusCreated, exp)

}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error file .env")
	}
}

var db *sql.DB

func main() {

	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	url := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table success")

	e.POST("/expenses", createExpenses)
	// e.POST("/users", createUserHandler)
	// e.GET("/users", getUsersHandler)
	// e.GET("/users/:id", getUserHandler)

	log.Fatal(e.Start(os.Getenv("PORT")))
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
}

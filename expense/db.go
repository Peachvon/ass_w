package expense

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	url := os.Getenv("DATABASE_URL")
	var err error
	DB, err = sql.Open("postgres", url)

	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer DB.Close()
	CreateTableExpenses(DB)
}
func CreateTableExpenses(db *sql.DB) error {

	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`
	_, err := db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table success")

	return nil
}

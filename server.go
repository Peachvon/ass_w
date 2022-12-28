package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

type User struct {
	Name string
	Age  int
}

var users = []User{
	{Name: "peach", Age: 19},
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {

		u := User{Name: "von", Age: 12}
		users = append(users, u)
		return c.JSON(http.StatusOK, users)
	})
	// e.POST("/users", createUserHandler)
	// e.GET("/users", getUsersHandler)
	// e.GET("/users/:id", getUserHandler)

	log.Fatal(e.Start(":2565"))
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
}

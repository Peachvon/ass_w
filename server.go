package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Peachvon/assessment/expense"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u, p, ok := c.Request().BasicAuth()
		if !ok {

			return c.JSON(http.StatusUnauthorized, expense.Err{Message: "Unauthorized"})
		}

		if u != "peachvon" || p != "20232566" {

			return c.JSON(http.StatusUnauthorized, expense.Err{Message: "Unauthorized"})
		}

		next(c)
		return nil
	}
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error file .env")
	}
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "peachvon" && password == "20232566" {
	// 		return true, nil
	// 	}
	// 	return false, nil

	// }))
	expense.InitDB()
	e.POST("/expenses", expense.CreateExpenseHandler)
	e.GET("/expenses/:id", expense.GetExpenseHandler)
	e.GET("/expenses", AuthMiddleware(expense.GetExpensesHandler))
	e.PUT("/expenses/:id", expense.PutExpenseHandler)

	go func() {
		err := e.Start(os.Getenv("PORT"))
		if err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}

	}()
	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("shutdown. . .")
	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("shutdown err:", err)
	}
}

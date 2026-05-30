package main

import (
	"os"
	"study/internal/http_handler"
	"study/internal/service"
	"study/internal/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	CONN_STR := os.Getenv("CONN_STR")
	CONN_STR_RDB := os.Getenv("CONN_STR_RDB")

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	store := storage.NewStorage(CONN_STR)
	redis := storage.NewRedis(CONN_STR_RDB)

	srv := service.NewService(store, redis)

	handler := http_handler.NewHandler(srv)

	e.POST("/user", handler.PostUser)
	e.GET("/users", handler.GetUsers)
	e.GET("/user", handler.GetUser)

	e.Logger.Fatal(e.Start(":8080"))
}

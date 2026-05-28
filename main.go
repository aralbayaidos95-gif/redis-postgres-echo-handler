package main

import (
	"study/internal/http_handler"
	"study/internal/service"
	"study/internal/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	connStr := "postgres://postgres:2055@localhost:5432/postgres"
	connStrRDB := "127.0.0.1:6379"

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	store := storage.NewStorage(connStr)
	redis := storage.NewRedis(connStrRDB)

	srv := service.NewService(store, redis)

	handler := http_handler.NewHandler(srv)

	e.POST("/user", handler.PostUser)
	e.GET("/users", handler.GetUsers)
	e.GET("/user", handler.GetUser)

	e.Logger.Fatal(e.Start(":8080"))
}

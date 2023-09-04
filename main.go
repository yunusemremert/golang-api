package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-api/app"
	"golang-api/configs"
	"golang-api/repository"
	"golang-api/services"
)

func main() {
	appRoute := fiber.New()

	dbClient := configs.GetCollection(configs.ConnectMongoDB(), "todos")

	TodoRepositoryDB := repository.NewTodoRepositoryDB(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDB)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)

	err := appRoute.Listen(":8080")
	if err != nil {
		return
	}
}

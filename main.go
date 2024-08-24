package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nizsimsek/go-fiber-mongodb/app"
	"github.com/nizsimsek/go-fiber-mongodb/configs"
	"github.com/nizsimsek/go-fiber-mongodb/repository"
	"github.com/nizsimsek/go-fiber-mongodb/services"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()

	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDB := repository.NewTodoRepositoryDb(dbClient)

	th := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDB)}

	appRoute.Post("/api/todo", th.CreateTodo)
	appRoute.Get("/api/todos", th.GetAllTodo)
	appRoute.Delete("/api/todo/:id", th.DeleteTodo)
	err := appRoute.Listen(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}

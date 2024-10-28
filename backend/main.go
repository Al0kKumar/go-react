package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID          int     `json:"id"` 
	Completed   bool     `json:"completed"`
	Body        string   `json:"body"`
}


var todos []Todo

func main() {
	fmt.Println("hello world")

	app := fiber.New()


	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg":"hello world"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil{
			return err
		}

		if todo.Body == ""{
			return c.Status(400).JSON(fiber.Map{"error":"todo body is reqiured"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	log.Fatal(app.Listen(":3000"))
}

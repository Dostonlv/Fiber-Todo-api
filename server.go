package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"todo/Controllers"
	"todo/model"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	//app.Use(middleware.CORS())

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&model.Todo{})

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/todos", Controllers.GetTodos(db))
	app.Get("/todos/:id", Controllers.GetTodo(db))
	app.Post("/todos", Controllers.NewTodo(db))
	app.Delete("/todos/:id", Controllers.DeleteTodo(db))
	app.Patch("/todos/:id", Controllers.UpdateTodo(db))

	log.Fatal(app.Listen(":4000"))
}

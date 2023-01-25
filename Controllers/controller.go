package Controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"todo/model"
)

func GetTodos(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var todos []model.Todo
		db.Find(&todos)
		return c.JSON(todos)
	}
}

func GetTodo(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var todo model.Todo
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(todo)
	}
}

func NewTodo(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var todo model.Todo
		if err := c.BodyParser(&todo); err != nil {
			return c.Status(500).SendString(err.Error())

		}
		db.Create(&todo)
		return c.JSON(todo)
	}
}

func DeleteTodo(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var todo model.Todo
		d := db.Where("id = ?", id).Delete(&todo)
		fmt.Println(d)
		return c.JSON("ok")
	}
}

func UpdateTodo(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var todo model.Todo
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			return c.Status(500).SendString(err.Error())

		}
		if err := c.BodyParser(&todo); err != nil {
			return c.Status(500).SendString(err.Error())

		}
		db.Save(&todo)
		return c.JSON(todo)
	}
}

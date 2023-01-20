package Controllers

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"todo/model"
)

func GetTodos(db *gorm.DB) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		var todos []model.Todo
		db.Find(&todos)
		c.JSON(todos)
	}
}

func GetTodo(db *gorm.DB) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id := c.Params("id")
		var todo model.Todo
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			c.Status(500).Send(err)
			return
		}
		c.JSON(todo)
	}
}

func NewTodo(db *gorm.DB) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		var todo model.Todo
		if err := c.BodyParser(&todo); err != nil {
			c.Status(500).Send(err)
			return
		}
		db.Create(&todo)
		c.JSON(todo)
	}
}

func DeleteTodo(db *gorm.DB) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id := c.Params("id")
		var todo model.Todo
		d := db.Where("id = ?", id).Delete(&todo)
		fmt.Println(d)
		c.Send("ok")
	}
}

func UpdateTodo(db *gorm.DB) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id := c.Params("id")
		var todo model.Todo
		if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
			c.Status(500).Send(err)
			return
		}
		if err := c.BodyParser(&todo); err != nil {
			c.Status(500).Send(err)
			return
		}
		db.Save(&todo)
		c.JSON(todo)
	}
}

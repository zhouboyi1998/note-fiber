package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	// 新建 Fiber 实例
	app := fiber.New()

	// Hello World
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.JSON(fiber.Map{
			"code":    http.StatusOK,
			"message": "Hello, " + name,
		})
	})

	// 启动服务
	app.Listen(":18099")
}

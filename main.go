package main

import (
	"github.com/gofiber/fiber/v2"
	"note-fiber/src/application"
	"note-fiber/src/router"
)

func main() {
	// 新建 Fiber 实例
	app := fiber.New()
	// 注册路由
	router.RegisterRouter(app)
	// 启动服务
	app.Listen(application.App.Server.Host + ":" + application.App.Server.Port)
}

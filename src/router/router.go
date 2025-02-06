package router

import (
	"github.com/gofiber/fiber/v2"
	"note-fiber/src/controller"
)

// RegisterRouter 注册路由
func RegisterRouter(f *fiber.App) {
	// 新建命令路由组
	g := f.Group("/command")
	{
		// 添加命令相关路由
		g.Get("/:commandId", controller.One)
		g.Get("/", controller.List)
		g.Post("/", controller.Insert)
		g.Post("/batch", controller.InsertBatch)
		g.Put("/", controller.Update)
		g.Put("/batch", controller.UpdateBatch)
		g.Delete("/:commandId", controller.Delete)
		g.Delete("/batch", controller.DeleteBatch)
		g.Get("/select/:commandName", controller.Select)
		g.Get("/name-list", controller.NameList)
	}
}

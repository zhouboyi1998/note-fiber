package controller

import (
	"github.com/gofiber/fiber/v2"
	"note-fiber/src/repository"
)

// One 根据id查询命令
func One(ctx *fiber.Ctx) error {
	command := repository.One(ctx)
	return ctx.JSON(command)
}

// List 查询命令列表
func List(ctx *fiber.Ctx) error {
	commandArray := repository.List(ctx)
	return ctx.JSON(commandArray)
}

// Insert 新增命令
func Insert(ctx *fiber.Ctx) error {
	result, commandName := repository.Insert(ctx)
	return ctx.JSON(fiber.Map{
		"result":  result,
		"command": commandName,
	})
}

// InsertBatch 批量新增命令
func InsertBatch(ctx *fiber.Ctx) error {
	result := repository.InsertBatch(ctx)
	return ctx.JSON(result)
}

// Update 修改命令
func Update(ctx *fiber.Ctx) error {
	result := repository.Update(ctx)
	return ctx.JSON(result)
}

// UpdateBatch 批量修改命令
func UpdateBatch(ctx *fiber.Ctx) error {
	result := repository.UpdateBatch(ctx)
	return ctx.JSON(result)
}

// Delete 删除命令
func Delete(ctx *fiber.Ctx) error {
	result, objectId := repository.Delete(ctx)
	return ctx.JSON(fiber.Map{
		"result": result,
		"_id":    objectId,
	})
}

// DeleteBatch 批量删除命令
func DeleteBatch(ctx *fiber.Ctx) error {
	result, objectIds := repository.DeleteBatch(ctx)
	return ctx.JSON(fiber.Map{
		"result": result,
		"_ids":   objectIds,
	})
}

// Select 查询命令
func Select(ctx *fiber.Ctx) error {
	command := repository.Select(ctx)
	return ctx.JSON(command)
}

// NameList 查询命令名称列表
func NameList(ctx *fiber.Ctx) error {
	nameArray := repository.NameList(ctx)
	return ctx.JSON(nameArray)
}

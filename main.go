package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// 定义全局的中间件
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("🥇 First 2222handler")
		return c.Next()
	})
	// 定义路由
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("你好小钟同学!")
	})

	app.Listen(":3000")
}

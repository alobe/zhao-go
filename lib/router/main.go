package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitRouter() {
	app := fiber.New()

	// middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET",
	}))

	//路由绑定
	images := app.Group("/api/image")
	images.Get("/all", getAll)

	app.Listen(":3000")
}

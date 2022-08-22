package router

// refer to https://github.dev/NikSchaefer/go-fiber
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

	api := app.Group("/api")

	//路由绑定
	images := api.Group("/image")
	images.Get("/all", getAll)
	images.Post("/create", createImage)

	app.Listen(":3000")
}

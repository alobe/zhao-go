package router

// refer to https://github.dev/NikSchaefer/go-fiber
import (
	"zhao-go/lib/middleware"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitRouter() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	// middleware
	app.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET",
		}),
		middleware.Logger(),
	)

	api := app.Group("/api")

	//路由绑定
	images := api.Group("/image")
	images.Get("/all", getAll)
	images.Post("/create", createImage)
	images.Post("/batch_create", createImages)

	app.Listen(":3000")
}

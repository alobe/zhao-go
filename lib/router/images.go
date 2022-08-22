package router

import (
	"zhao-go/lib/database"
	"zhao-go/lib/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func imageDb() *gorm.DB {
	return database.DB.Model(&model.Image{}).Table("Images")
}

type Res fiber.Map

func getAll(c *fiber.Ctx) error {
	db := imageDb()
	var images []model.Image
	db.Limit(100).Find(&images)
	return c.JSON(Res{
		"code": 200,
		"data": images,
	})
}

func createImage(c *fiber.Ctx) error {
	json := new(model.Image)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(Res{
			"code":    400,
			"message": "invalid image strut",
		})
	}

	if err := imageDb().Create(&model.Image{Url: json.Url}); err != nil {
		return c.JSON(Res{
			"code":    400,
			"message": "create image fail",
		})
	}

	return c.JSON(Res{
		"code": 400,
	})
}

package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"zhao-go/lib/database"
	"zhao-go/lib/model"
)

func imageDb() *gorm.DB {
	return database.DB.Model(&model.Image{}).Table("Images")
}

func getAll(c *fiber.Ctx) error {
	db := imageDb()
	var images []model.Image
	db.Limit(100).Find(&images)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    images,
	})
}

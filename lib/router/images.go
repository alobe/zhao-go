package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"zhao-go/lib/database"
	"zhao-go/lib/model"
)

func imageDb() *gorm.DB {
	return database.DB.Model(&model.Image{})
}

func getAll(c *fiber.Ctx) error {
	//有bug
	db := imageDb()
	images := db.Order("id asc").Limit(100).Find(model.Image{})
	log.Print(images)
	//有bug
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    0, /*images*/
	})
}

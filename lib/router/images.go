package router

import (
	"zhao-go/lib/database"
	"zhao-go/lib/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func imageDb() *gorm.DB {
	return database.DB.Model(&model.Image{}).Table("images")
}

type Res fiber.Map

func getAll(c *fiber.Ctx) error {
	db := imageDb()
	var images []model.Image
	db.Limit(100).Find(&images)
	return res(c, 200, nil, images)
}

func createImage(c *fiber.Ctx) error {
	json := new(model.Image)
	if err := c.BodyParser(json); err != nil {
		return res(c, 400, "invalid image strut")
	}

	tx := imageDb().Create(&model.Image{Url: json.Url})
	if tx.Error != nil {
		return res(c, 400, "create image fail")
	}

	return res(c, 200)
}

type ImagesBody struct {
	images []model.Image
}

func createImages(c *fiber.Ctx) error {
	json := new (ImagesBody)
	if err := c.BodyParser(json); err != nil {
		return res(c, 400, "invalid images strut")
	}

	tx := imageDb().CreateInBatches(&json.images, 200)
	if tx.Error != nil {
		return res(c, 400, "images create fail")
	}

	return res(c, 200, nil)
}

func res(ctx *fiber.Ctx, code int, msg string | nil, data interface {} | nil) error {
	return ctx.JSON(Res{
		"code": code,
		"message": msg,
		"data": data,
	})
}

package router

import (
	"zhao-go/lib/database"
	"zhao-go/lib/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var imageDb *gorm.DB

func ImageDb() *gorm.DB {
	if imageDb == nil {
		imageDb = database.GetTableDb(&model.Image{}, "images")
	}
	return imageDb
}

type Res fiber.Map

func getAll(c *fiber.Ctx) error {
	var images []model.Image
	ImageDb().Limit(100).Find(&images)
	return res(c, 200, "", images)
}

func createImage(c *fiber.Ctx) error {
	json := new(model.Image)
	if err := c.BodyParser(json); err != nil {
		return res(c, 400, "invalid image strut", nil)
	}

	tx := ImageDb().Create(&model.Image{Url: json.Url})
	if tx.Error != nil {
		return res(c, 400, "create image fail", nil)
	}

	return res(c, 200, "", nil)
}

type ImagesBody struct {
	Images []model.Image `json:"images"`
}

func createImages(c *fiber.Ctx) error {
	json := new(ImagesBody)
	if err := c.BodyParser(json); err != nil {
		return res(c, 400, "invalid images strut", nil)
	}

	tx := ImageDb().CreateInBatches(&json.Images, 200)
	if tx.Error != nil {
		return res(c, 400, "images create fail", nil)
	}

	return res(c, 200, "", nil)
}

func res(ctx *fiber.Ctx, code int, msg string, data interface{}) error {
	return ctx.JSON(Res{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}

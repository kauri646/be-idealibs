package handlersimages

import (
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kauri646/be-idealibs/config"
	"github.com/kauri646/be-idealibs/internal/models/images"
)

func ImageHandlerGetAll(ctx *fiber.Ctx) error {

	var images []images.Image

	result := config.DB.Find(&images)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(images)

}

func ImageHandlerCreate(ctx *fiber.Ctx) error {
	// Menginisialisasi variabel baru untuk menampung data yang dikirim dalam body request
	image := new(images.Image)

	// Parsing body request ke dalam struct image
	if err := ctx.BodyParser(image); err != nil {
		return err
	}

	// Simpan gambar di server
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	// Generate unique file name
	fileName := uuid.New().String() + filepath.Ext(file.Filename)
	if err := ctx.SaveFile(file, fileName); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	// Set nama file sebagai URL gambar
	image.File = fileName

	// Menambahkan data baru ke dalam basis data
	if err := config.DB.Create(image).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to store data",
			"error":   err.Error(),
		})
	}

	// Mengembalikan respons JSON yang berisi data yang baru saja ditambahkan
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    image,
	})
}

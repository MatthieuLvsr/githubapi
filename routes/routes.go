package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/download", Download)
}

func Download(c *fiber.Ctx) error {
	fileName := c.Query("name")
	filepath := fmt.Sprintf("./repos/%s.zip", fileName)

	exist := func(filename string) bool {
		f, err := os.Open(filepath)
		if err != nil {
			return false
		}
		if err != nil {
			log.Fatalf("Could not download file %s \n", filepath)
			return false
		}
		defer f.Close()
		return true
		}(fileName)
		if !exist {
			return c.Status(400).SendString("Error")
		}
		return c.Download(filepath)
}

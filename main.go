package main

import (
	"fmt"
	"log"

	"github.com/MatthieuLvsr/githubapi/csvmaker"
	"github.com/MatthieuLvsr/githubapi/gitapi"
	"github.com/MatthieuLvsr/githubapi/requester"
	"github.com/MatthieuLvsr/githubapi/routes"
	"github.com/MatthieuLvsr/githubapi/zipmaker"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})
	routes.SetupRoutes(app)
	gitapi.Setup()
	result := requester.Request()
	repos := requester.ParseRepos(result)
	gitapi.Clone(repos)
	csvmaker.Record(repos)
	err := zipmaker.ZipDirectory(fmt.Sprintf("./repos/%s",gitapi.GIT_USER),fmt.Sprintf("./repos/%s.zip",gitapi.GIT_USER))
	if err != nil {
		log.Fatal(err)
	}
	app.Listen(":8081")
}
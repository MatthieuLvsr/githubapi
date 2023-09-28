package main

import (
	"fmt"
	"log"

	"github.com/MatthieuLvsr/githubapi/csvmaker"
	"github.com/MatthieuLvsr/githubapi/gitapi"
	"github.com/MatthieuLvsr/githubapi/requester"
	"github.com/MatthieuLvsr/githubapi/zipmaker"
)

func main() {
	gitapi.Setup()
	result := requester.Request()
	repos := requester.ParseRepos(result)
	gitapi.Clone(repos)
	csvmaker.Record(repos)
	err := zipmaker.ZipDirectory(fmt.Sprintf("./repos/%s",gitapi.GIT_USER),fmt.Sprintf("./repos/%s.zip",gitapi.GIT_USER))
	if err != nil {
		log.Fatal(err)
	}
}
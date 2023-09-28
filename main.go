package main

import (
	"fmt"

	"github.com/MatthieuLvsr/githubapi/gitapi"
	"github.com/MatthieuLvsr/githubapi/requester"
)

func main() {
	gitapi.Setup()
	result := requester.Request()
	fmt.Println(requester.ParseRepos(result))
}
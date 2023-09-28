package main

import (
	"github.com/MatthieuLvsr/githubapi/csvmaker"
	"github.com/MatthieuLvsr/githubapi/gitapi"
	"github.com/MatthieuLvsr/githubapi/requester"
)

func main() {
	gitapi.Setup()
	result := requester.Request()
	repos := requester.ParseRepos(result)
	gitapi.Clone(repos)
	csvmaker.Record(repos)
}
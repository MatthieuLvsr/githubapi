package gitapi

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/MatthieuLvsr/githubapi/models"
)

var API_KEY string
var GIT_USER string
var REPOS_PER_PAGE string

func Setup() {
	key,ok := os.LookupEnv("GITHUB_KEY")
	if !ok {
		log.Fatal("Key not found")
	}
	user,ok := os.LookupEnv("GITHUB_USRER")
	if !ok {
		log.Fatal("User not found")
	}
	per_page,ok := os.LookupEnv("REPOS_PER_PAGE")
	if !ok {
		log.Fatal("Per_page not found")
	}
	fmt.Println("Key setup : ",key)
	fmt.Println("User setup : ",user)
	fmt.Println("Per_page setup : ",per_page)
	API_KEY = key
	GIT_USER = user
	REPOS_PER_PAGE = per_page
}

func Clone(repos []models.Repos){
	var wg sync.WaitGroup
	wg.Add(len(repos))
	
	for _,repo:=range repos {
		go func(url string, name string){
			defer wg.Done()
			path := fmt.Sprintf("./repos/%s/%s",GIT_USER,name)
			cmd := exec.Command("git","clone",url,path)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println("Error : ",err)
			}

			cmd = exec.Command("git", "pull")
			cmd.Dir = path
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				fmt.Println("Error : ",err)
			}

			cmd = exec.Command("git", "fetch")
			cmd.Dir = path
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				fmt.Println("Error : ",err)
			}

		}(repo.Html_url, repo.Name)
	}
	wg.Wait()
}
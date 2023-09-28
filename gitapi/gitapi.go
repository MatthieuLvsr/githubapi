package gitapi

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"

	"github.com/MatthieuLvsr/githubapi/models"
)

var API_KEY string
var GIT_USER string
var REPOS_PER_PAGE int
var ORG_MODE bool

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
	int_per_page, err := strconv.Atoi(per_page)
	if err != nil{
		log.Fatal(err)
	}
	org_mode,ok := os.LookupEnv("ORG_MODE")
	if !ok {
		log.Fatal("Org_mode not found")
	}
	bool_org_mode, err := strconv.ParseBool(org_mode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Key setup : ",key)
	fmt.Println("User setup : ",user)
	fmt.Println("Per_page setup : ",int_per_page)
	fmt.Println("Org_mode setup : ",bool_org_mode)
	API_KEY = key
	GIT_USER = user
	REPOS_PER_PAGE = int_per_page
	ORG_MODE = bool_org_mode
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
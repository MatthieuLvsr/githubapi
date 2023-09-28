package requester

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/MatthieuLvsr/githubapi/gitapi"
	"github.com/MatthieuLvsr/githubapi/models"
)

func ParseRepos(body []byte) []models.Repos {
	var result []models.Repos
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Request()[]byte{
	client := &http.Client{}
	mode := "users"
	if gitapi.ORG_MODE {
		mode = "orgs"
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/%s/%s/repos?per_page=%d",mode,gitapi.GIT_USER,gitapi.REPOS_PER_PAGE), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization","Bearer " + gitapi.API_KEY) 
	req.Header.Add("Accept", "application/vnd.github+json") 
	
	resp, err := client.Do(req)
		if err != nil {
		log.Fatal(err)
	}	
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}
	return body
}
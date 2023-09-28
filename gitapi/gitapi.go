package gitapi

import (
	"fmt"
	"log"
	"os"
)

var API_KEY string
var GIT_USER string

func Setup() {
	key,ok := os.LookupEnv("GITHUB_KEY")
	if !ok {
		log.Fatal("Key not found")
	}
	user,ok := os.LookupEnv("GITHUB_USRER")
	if !ok {
		log.Fatal("User not found")
	}
	fmt.Println("Key setup : ",key)
	fmt.Println("User setup : ",user)
	API_KEY = key
	GIT_USER = user
}
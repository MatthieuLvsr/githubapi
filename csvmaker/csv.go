package csvmaker

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/MatthieuLvsr/githubapi/gitapi"
	"github.com/MatthieuLvsr/githubapi/models"
)

func Record(repos []models.Repos) {
	f,err:= os.OpenFile(fmt.Sprintf("./repos/%s/%s.csv",gitapi.GIT_USER,gitapi.GIT_USER), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	w.Write([]string{"NAME","URL","DESCRIPTION","CREATION","UPDATE"})
	for _, repo := range repos {
		if err := w.Write(models.Tostring(repo)); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
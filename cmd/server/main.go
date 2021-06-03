package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rubberyconf/rubberyconf/api"
	"github.com/rubberyconf/rubberyconf/internal/config"
)

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	conf := config.NewConfiguration(filepath.Join(path, "../../config/local.yml"))

	router := api.NewRouter()

	log.Printf(
		"Api started at port: %s",
		conf.Api.Port,
	)

	log.Fatal(http.ListenAndServe(":"+conf.Api.Port, router))

}

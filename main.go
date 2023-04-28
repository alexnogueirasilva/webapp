package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookeis"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Load()
	cookeis.Setup()
	utils.LoadTemplates()
	r := router.Generate()

	fmt.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

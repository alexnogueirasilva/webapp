package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.LoadTemplates()

	fmt.Println("Running WebApp listening at port 3000...")
	r := router.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"restapiwithgo/router"
)

func main() {
	rou := router.Router()
	fmt.Println("Running....")
	log.Fatal(http.ListenAndServe(":8001", rou))
}

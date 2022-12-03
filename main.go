package main

import (
	"log"
	"net/http"

	"github.com/anon1689501/AOC2022web/controllers"
)

func main() {
	http.HandleFunc("/", controllers.MakeIndex)

	log.Fatal(http.ListenAndServe(":80", nil))
}

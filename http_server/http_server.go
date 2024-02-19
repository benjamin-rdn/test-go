package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Up and running!")

	http.HandleFunc("/", RootHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

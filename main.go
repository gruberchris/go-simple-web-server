package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("http server started on :8181")

	err := http.ListenAndServe(":8181", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

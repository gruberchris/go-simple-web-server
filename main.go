package main

import (
	"fmt"
	"log"
	"net/http"
)

func logRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("%s request for %s from client address %s", r.Method, r.RequestURI, r.RemoteAddr))
		h.ServeHTTP(w, r) // call original
	})
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", logRequest(fs))

	log.Println("http server started on :8181")

	err := http.ListenAndServe(":8181", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Fatal()
}

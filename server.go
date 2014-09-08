package main

import "net/http"
import "github.com/gorilla/mux"

func FooHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("おっぱい"))
}

func main() {
	routes := mux.NewRouter()
	routes.HandleFunc("/", FooHandler)
	http.Handle("/", routes)

	http.ListenAndServe(":12345", nil)
}

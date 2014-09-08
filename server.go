package main

import "net/http"
import "github.com/gorilla/mux"
import "github.com/otiai10/ocr-kcwidget/handler"

func FooHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("foooo"))
}

func main() {
	routes := mux.NewRouter()
	routes.HandleFunc("/", FooHandler)

	routes.Path("/oppai/{count}").Handler(handler.Oppai{})
	routes.Path("/oppai").Handler(handler.Oppai{})

	http.Handle("/", routes)

	http.ListenAndServe(":12345", nil)
}

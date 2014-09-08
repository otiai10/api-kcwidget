package handler

import "github.com/gorilla/mux"
import "net/http"

type Handler struct{}

func (h Handler) Vars(req *http.Request) map[string]string {
	return mux.Vars(req)
}

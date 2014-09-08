package handler

import "net/http"

type Oppai struct {
	Handler
}

func (h Oppai) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	res := []byte("おっぱい")
	if count, ok := h.Vars(req)["count"]; ok {
		res = append(res, []byte("が"+count+"こ")...)
	}
	rw.Write(res)
}

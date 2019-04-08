package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"videoservice/api/session"
)

type middlerWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middlerWareHandler{}
	m.r = r
	return m
}

func (m middlerWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandler() *httprouter.Router {
    router := httprouter.New()

    router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)

	return  router
}

func Prepare(){
	session.LoadAllSession()
}

func main() {
    Prepare()
    r := RegisterHandler()
    mr := NewMiddleWareHandler(r)

    http.ListenAndServe(":8080", mr)
}

package routes

import (
	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
	user UserAPI
}

func ProvideRouter(api UserAPI) *Router {
	router := Router{mux.NewRouter(), api}
	router.Setup()
	return &router
}

func (router *Router) Setup() *Router {
	s := router.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/register", router.user.Register).Methods("POST")
	return router
}


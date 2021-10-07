package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
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

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		response, _ := json.Marshal(payload)
		w.Write(response)
	}
}

func getQueryParams(queryParams url.Values) (filter string, page int64, pageSize int64) {
	filter = queryParams.Get("search")
	page = 1
	if n, err := strconv.Atoi(queryParams.Get("page")); err == nil {
		page = int64(n)
	}
	pageSize = 5
	if n, err := strconv.Atoi(queryParams.Get("pageSize")); err == nil {
		pageSize = int64(n)
	}
	return filter, page, pageSize
}

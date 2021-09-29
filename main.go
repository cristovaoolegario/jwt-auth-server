package free_auth_server

import (
	"fmt"
	"github.com/cristovaoolegario/free-auth-server/db"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	db.Connect(os.Getenv("ENV"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_NAME"))

	router := mux.Router{}

	Run(os.Getenv("PORT"),
		os.Getenv("ENV"),
		&router)
}

func Run(port, env string, r *mux.Router) {
	fmt.Println("Server running in port:", port)
	if env == "dev" {
		corsWrapper := cors.New(cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
		})
		log.Fatal(http.ListenAndServe(port, corsWrapper.Handler(r)))
	} else {
		log.Fatal(http.ListenAndServe(port, r))
	}
}

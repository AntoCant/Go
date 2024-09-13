package main

import (
	"log"
	"net/http"

	"Api-users/internal/handler"
	"Api-users/pkg/db"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDb()
	router := mux.NewRouter()
	handler.SetupRoutes(router)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

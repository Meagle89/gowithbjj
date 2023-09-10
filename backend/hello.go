package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var PORT = ":8080"

func main() {
	db, err := gorm.Open(sqlite.Open("./bjjTechniques.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	db.AutoMigrate(&Technique{}) //gorm will create the table if it doesn't exist

	fmt.Println("Server running on port", PORT)

	SeedDatabase(db)

	// Initialize a new http.ServeMux (it's an HTTP request multiplexer)
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.Handle("/techniques", LoggerMiddleWare(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			AddTechnique(w, r, db)
			return
		}

		if r.Method == "GET" {
			GetAllTechniques(w, r, db)
			return
		}
	}))).Methods("GET", "POST")

	router.Handle("/techniques/{id:[0-9]+}", LoggerMiddleWare(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			UpdateTechnique(w, r, db)
			return
		}

		if r.Method == "DELETE" {
			DeleteTechnique(w, r, db)
			return
		}
	}))).Methods("PUT", "DELETE")

	http.ListenAndServe(PORT, router)

}

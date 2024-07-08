package main

import (
	"football-forum/controllers"
	"football-forum/database"
	"football-forum/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	log.Println("Uygulama başlatılıyor...")

	database.InitDB()

	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	router.HandleFunc("/topics", middleware.Authenticate(controllers.CreateTopic)).Methods("POST")
	router.HandleFunc("/comments", middleware.Authenticate(controllers.CreateComment)).Methods("POST")
	router.HandleFunc("/topics", controllers.GetTopics).Methods("GET")
	router.HandleFunc("/comments/{id}", controllers.GetComments).Methods("GET")

	router.HandleFunc("/admin/delete-topic/{id}", middleware.AdminAuthenticate(controllers.DeleteTopic)).Methods("DELETE")
	router.HandleFunc("/admin/delete-comment/{id}", middleware.AdminAuthenticate(controllers.DeleteComment)).Methods("DELETE")

	handler := cors.Default().Handler(router)
	log.Println("Port 8080 üzerinde dinleniyor")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

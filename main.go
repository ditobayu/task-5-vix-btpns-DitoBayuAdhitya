package main

import (
	"log"
	"net/http"

	"github.com/ditobayu/task-5-vix-btpns-Dito-Bayu-Adhitya/controllers"
	"github.com/ditobayu/task-5-vix-btpns-Dito-Bayu-Adhitya/middlewares"
	"github.com/ditobayu/task-5-vix-btpns-Dito-Bayu-Adhitya/models"

	// "github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

 
func main() {
	// r := gin.Default()
	models.ConnectToDatabase()

	
	r := mux.NewRouter()

	r.HandleFunc("/users/register", controllers.Register).Methods("POST")
	r.HandleFunc("/users/login", controllers.Login).Methods("POST")
	r.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE")

	
	api := r.PathPrefix("/").Subrouter()
	api.HandleFunc("/photos", controllers.AddPhoto).Methods("POST")
	api.HandleFunc("/photos", controllers.GetPhoto).Methods("GET")
	api.HandleFunc("/{photoId}", controllers.UpdatePhoto).Methods("PUT")
	api.HandleFunc("/{photoId}", controllers.DeletePhoto).Methods("DELETE")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))


	// User
	// r.POST("/users/register", controllers.Register)
	// r.GET("/users/login", controllers.Login)
	// r.PUT("/users/:userId", controllers.UpdateUser)
	// r.DELETE("/users/:userId", controllers.DeleteUser)
	
	// // Photo
	// r.POST("/photos", controllers.CreatePhoto)
	// r.GET("/photos", controllers.ReadPhoto)
	// r.PUT("/:photoId", controllers.UpdatePhoto)
	// r.DELETE("/:photoId", controllers.DeletePhoto)

	// r.Run()
}
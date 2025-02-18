// /router/router.go
package router

import (
	"suriyaa.com/handlers" // Adjust with your actual module name

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter sets up the router and routes
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	// Define your routes here
	r.HandleFunc("/api/v1/create-player", handlers.CreatePlayer).Methods("POST")
	r.HandleFunc("/api/v1/get-player", handlers.GetPlayer).Methods("GET")
	r.HandleFunc("/api/v1/update-player", handlers.UpdatePlayer).Methods("POST")
	r.HandleFunc("/api/v1/start-game", handlers.StartGame).Methods("POST")
	r.HandleFunc("/api/v1/end-game", handlers.EndGame).Methods("POST")
	r.HandleFunc("/api/v1/mode-statistics", handlers.GetPopularModes).Methods("GET")
	r.HandleFunc("/api/v1/printAll", handlers.PrintAll).Methods("GET") // Just for debugging/logging purpose nothing to display in http reponse
	// Serve Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return r
}

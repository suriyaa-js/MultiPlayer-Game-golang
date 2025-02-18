package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"suriyaa.com/cache"
	_ "suriyaa.com/docs"
	"suriyaa.com/router"
	"suriyaa.com/storage"
)

func homeHandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	fmt.Fprintf(resp, "Hello world \n")

}

func loadConfig() (string, time.Duration, time.Duration) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := os.Getenv("SERVER_ADDRESS")
	writeTimeout, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	readTimeout, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))

	return addr, time.Duration(writeTimeout) * time.Second, time.Duration(readTimeout) * time.Second
}

// Initiate web server
func main() {
	addr, writeTimeout, readTimeout := loadConfig()
	storage.InitializeDB()
	cache.InitializeCache()

	addr = "0.0.0.0:9100"
	// }
	//cache.Random()
	router := router.NewRouter()
	
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}
	log.Printf("Starting server at %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}

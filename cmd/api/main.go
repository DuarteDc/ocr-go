package api

import (
	"fmt"
	"log"

	"github.com/DuarteDc/ocr-golang.git/internal/infraestructure/http"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	router := http.NewRouter()

	log.Println("Server started on port 8080")

	if err := router.Run(fmt.Sprintf(":%s", "8080")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

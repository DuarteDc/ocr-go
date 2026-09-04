package main

import (
	"log"
	"os"

	"github.com/DuarteDc/ocr-golang.git/internal/infraestructure/datasource"
	"github.com/DuarteDc/ocr-golang.git/internal/infraestructure/datasource/repository"
	"github.com/DuarteDc/ocr-golang.git/internal/infraestructure/http"
	"github.com/DuarteDc/ocr-golang.git/internal/infraestructure/pdf"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, proceeding with environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	dbPool, err := datasource.NewPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbPool.Close()

	log.Println("Connected to the database successfully")

	documentRepository := repository.NewPgDocumentRepository(dbPool)
	documentPageRepository :=
		repository.NewPostgresDocumentPageRepository(dbPool)

	pdfExtractor :=
		pdf.NewExtractor()

	router := http.NewRouter(documentRepository, documentPageRepository, pdfExtractor)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

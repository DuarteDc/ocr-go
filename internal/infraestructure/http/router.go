package http

import (
	"github.com/DuarteDc/ocr-golang.git/internal/application"
	"github.com/DuarteDc/ocr-golang.git/internal/infraestructure/http/handlers"
	"github.com/DuarteDc/ocr-golang.git/internal/ports"
	"github.com/gin-gonic/gin"
)

func NewRouter(documentRepository ports.DocumentRepository, documentPageRepository ports.DocumentPageRepository,
	pdfExtractor ports.PDFExtractor, askService *application.AskService) *gin.Engine {

	router := gin.Default()

	documentHadler := handlers.NewDocumentHandler(documentRepository, documentPageRepository, pdfExtractor)
	askHandler := handlers.NewAskHandler(askService)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/search", documentHadler.Search)
	router.POST("/documents", documentHadler.Upload)
	router.POST("/ask", askHandler.Ask)

	return router

}

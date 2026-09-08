package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/DuarteDc/ocr-golang.git/internal/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DocumentHandler struct {
	documentRepository     ports.DocumentRepository
	documentPageRepository ports.DocumentPageRepository
	pdfExtractor           ports.PDFExtractor
}

func NewDocumentHandler(
	documentRepository ports.DocumentRepository,
	documentPageRepository ports.DocumentPageRepository,
	pdfExtractor ports.PDFExtractor,
) *DocumentHandler {
	return &DocumentHandler{
		documentRepository:     documentRepository,
		documentPageRepository: documentPageRepository,
		pdfExtractor:           pdfExtractor,
	}
}

func (h *DocumentHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File is required",
		})
		return
	}
	extension := filepath.Ext(file.Filename)
	storageName := uuid.New().String() + extension

	filePath := filepath.Join("storage/documents", storageName)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	documentID, err := h.documentRepository.Create(c.Request.Context(), file.Filename, storageName, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	pages, err := h.pdfExtractor.Extract(c.Request.Context(), filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to extract PDF text",
		})
		return
	}

	if err := h.documentPageRepository.CreateMany(
		c.Request.Context(),
		documentID,
		pages,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save document pages",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":           documentID,
		"name":         file.Filename,
		"storage_name": storageName,
		"path":         filePath,
	})
}

func (h *DocumentHandler) Search(c *gin.Context) {
	query := c.Query("q")

	if query == "" {
		c.JSON(400, gin.H{
			"error": "q is required",
		})
		return
	}

	results, err := h.documentPageRepository.Search(c.Request.Context(), query, 10)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"results": results,
	})
}

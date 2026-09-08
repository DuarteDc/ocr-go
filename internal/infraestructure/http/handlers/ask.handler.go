package handlers

import (
	"github.com/DuarteDc/ocr-golang.git/internal/application"
	"github.com/gin-gonic/gin"
)

type AskHandler struct {
	service *application.AskService
}

func NewAskHandler(service *application.AskService) *AskHandler {
	return &AskHandler{
		service: service,
	}
}

type AskRequest struct {
	Question string `json:"question"`
}

func (h *AskHandler) Ask(c *gin.Context) {
	var request AskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	answer, err := h.service.Ask(c.Request.Context(), request.Question)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, answer)
}

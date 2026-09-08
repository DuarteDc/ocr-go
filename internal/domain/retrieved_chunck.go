package domain

import "github.com/google/uuid"

type RetrievedChunk struct {
	DocumentID uuid.UUID `json:"documentId"`
	FileName   string    `json:"fileName"`
	PageNumber int       `json:"pageNumber"`
	Content    string    `json:"content"`
}

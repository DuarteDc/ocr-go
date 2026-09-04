package domain

import "github.com/google/uuid"

type SearchResult struct {
	DocumentID uuid.UUID `json:"documentId`
	FileName   string    `json:"fileName"`
	PageNumber int       `json:"pageNumber"`
	Snippet    string    `json:"snippet"`
}

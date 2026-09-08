package ports

import (
	"context"

	"github.com/DuarteDc/ocr-golang.git/internal/domain"
	"github.com/google/uuid"
)

type DocumentPageRepository interface {
	CreateMany(
		ctx context.Context,
		documentID uuid.UUID,
		pages []domain.DocuementPage,
	) error

	Search(ctx context.Context, query string, limit int) ([]domain.RetrievedChunk, error)
}

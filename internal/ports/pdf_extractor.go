package ports

import (
	"context"

	"github.com/DuarteDc/ocr-golang.git/internal/domain"
)

type PDFExtractor interface {
	Extract(ctx context.Context, filePath string) ([]domain.DocuementPage, error)
}

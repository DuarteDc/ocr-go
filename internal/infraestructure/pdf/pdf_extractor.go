package pdf

import (
	"context"
	"fmt"

	"github.com/DuarteDc/ocr-golang.git/internal/domain"

	"github.com/ledongthuc/pdf"
)

type Extarctor struct{}

func NewExtractor() *Extarctor {
	return &Extarctor{}
}

func (e *Extarctor) Extract(ctx context.Context, filePath string) ([]domain.DocuementPage, error) {
	file, reader, err := pdf.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open ODF file: %w", err)
	}

	defer file.Close()

	totalPages := reader.NumPage()
	pages := make([]domain.DocuementPage, 0, totalPages)

	for pageNumber := 1; pageNumber <= totalPages; pageNumber++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		page := reader.Page(pageNumber)
		if page.V.IsNull() {
			continue
		}

		content, err := page.GetPlainText(nil)

		if err != nil {
			return nil, fmt.Errorf("get plain text from page %d: %w", pageNumber, err)
		}

		pages = append(pages, domain.DocuementPage{
			PageNumber: pageNumber,
			Content:    content,
		})
	}
	return pages, nil
}

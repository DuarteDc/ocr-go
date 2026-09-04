package repository

import (
	"context"
	"fmt"

	"github.com/DuarteDc/ocr-golang.git/internal/domain"
	"github.com/DuarteDc/ocr-golang.git/internal/ports"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDocumentPageRepository struct {
	db *pgxpool.Pool
}

func NewPostgresDocumentPageRepository(
	db *pgxpool.Pool,
) ports.DocumentPageRepository {
	return &PostgresDocumentPageRepository{
		db: db,
	}
}

func (r *PostgresDocumentPageRepository) CreateMany(
	ctx context.Context,
	documentID uuid.UUID,
	pages []domain.DocuementPage,
) error {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	for _, page := range pages {
		_, err := tx.Exec(
			ctx,
			`
			INSERT INTO document_pages (
				id,
				document_id,
				page_number,
				content,
				search_vector
			)
			VALUES ($1, $2, $3, $4, to_tsvector('spanish', $4))
			`,
			uuid.New(),
			documentID,
			page.PageNumber,
			page.Content,
		)

		if err != nil {
			return fmt.Errorf(
				"insert page %d: %w",
				page.PageNumber,
				err,
			)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *PostgresDocumentPageRepository) Search(
	ctx context.Context,
	query string,
) ([]domain.SearchResult, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
		 	d.id,
			d.file_name,
			dp.page_number,
			ts_headline('spanish', dp.content, plainto_tsquery('spanish', $1)) AS snippet
		FROM 
			document_pages dp
		INNER JOIN documents d ON dp.document_id = d.id
		WHERE 
			dp.search_vector @@ plainto_tsquery('spanish', $1)
		ORDER BY dp.page_number LIMIT 50
	`, query)

	if err != nil {
		return nil, fmt.Errorf("search query: %w", err)
	}

	defer rows.Close()
	results := make([]domain.SearchResult, 0)
	for rows.Next() {
		var result domain.SearchResult
		if err := rows.Scan(
			&result.DocumentID,
			&result.FileName,
			&result.PageNumber,
			&result.Snippet,
		); err != nil {
			return nil, fmt.Errorf("scan result: %w", err)
		}
		results = append(results, result)
	}
	return results, nil
}

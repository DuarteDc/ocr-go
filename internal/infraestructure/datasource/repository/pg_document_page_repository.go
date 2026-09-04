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
				content
			)
			VALUES ($1, $2, $3, $4)
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

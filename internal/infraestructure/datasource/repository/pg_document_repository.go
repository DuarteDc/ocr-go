package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgDocumentRepository struct {
	db *pgxpool.Pool
}

func NewPgDocumentRepository(db *pgxpool.Pool) *PgDocumentRepository {
	return &PgDocumentRepository{
		db: db,
	}
}

func (r *PgDocumentRepository) Create(ctx context.Context, fileName string, storageName string, filePath string) (uuid.UUID, error) {
	var id uuid.UUID
	err := r.db.QueryRow(ctx, "INSERT INTO documents (file_name, storage_name, file_path) VALUES ($1, $2, $3) RETURNING id", fileName, storageName, filePath).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil

}

package ports

import (
	"context"

	"github.com/google/uuid"
)

type DocumentRepository interface {
	Create(ctx context.Context, fileName string, storageName string, filePath string) (uuid.UUID, error)
}

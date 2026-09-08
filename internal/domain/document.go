package domain

import "time"

type Document struct {
	FileName    string
	FilePath    string
	StorageName string
	CreatedAt   time.Time
}

package ports

import "context"

type LLM interface {
	Generate(ctx context.Context, question string, context string) (string, error)
}

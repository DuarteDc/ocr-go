package application

import (
	"context"
	"fmt"
	"strings"

	"github.com/DuarteDc/ocr-golang.git/internal/domain"
	"github.com/DuarteDc/ocr-golang.git/internal/ports"
)

type AskService struct {
	repository ports.DocumentPageRepository
	llm        ports.LLM
}

func NewAskService(repository ports.DocumentPageRepository, llm ports.LLM) *AskService {
	return &AskService{
		repository: repository,
		llm:        llm,
	}
}

func (s *AskService) Ask(ctx context.Context, question string) (*domain.Answer, error) {
	chunks, err := s.repository.Search(ctx, question, 10)

	if err != nil {
		return nil, err
	}

	contextText := buildContext(chunks)

	fmt.Printf("Contexto generado:\n%s\n", contextText)

	answer, err := s.llm.Generate(ctx, question, contextText)

	if err != nil {
		return nil, err
	}

	return &domain.Answer{
		Answer:  answer,
		Sources: chunks,
	}, nil

}

func buildContext(chunks []domain.RetrievedChunk) string {
	var builder strings.Builder

	for _, chunk := range chunks {
		builder.WriteString(
			fmt.Sprintf(
				"Documento: %s\nPágina: %d\n%s\n\n",
				chunk.FileName,
				chunk.PageNumber,
				chunk.Content,
			),
		)
	}

	return builder.String()

}

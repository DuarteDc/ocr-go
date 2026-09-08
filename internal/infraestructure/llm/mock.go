package llm

import "context"

type MockLLM struct{}

func NewMockLLM() *MockLLM {
	return &MockLLM{}
}

func (m *MockLLM) Generate(cxt context.Context, question string, context string) (string, error) {
	return context, nil
}

package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OllamaClient struct {
	baseURL string
	model   string
	client  *http.Client
}

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type GenerateResponse struct {
	Response string `json:"response"`
}

func NewOllamaClient(baseURL string, model string) *OllamaClient {
	return &OllamaClient{
		baseURL: baseURL,
		model:   model,
		client:  &http.Client{},
	}
}

func (o *OllamaClient) Generate(ctx context.Context, question string, contextText string) (string, error) {
	prompt := fmt.Sprintf(`
Eres un asistente que responde preguntas sobre documentos PDF.

Reglas:
- Usa únicamente la información del contexto.
- Si la respuesta no aparece en el contexto, responde:
  "No encontré información suficiente."
- Menciona el nombre del documento cuando sea posible.
- Sé breve y preciso.

CONTEXTO:
%s

PREGUNTA:
%s
`, contextText, question)

	payload := GenerateRequest{
		Model:  o.model,
		Prompt: prompt,
		Stream: false,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+"/api/generate", bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := o.client.Do(req)
	if err != nil {
		return "", err
	}

	defer req.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result GenerateResponse

	if err := json.Unmarshal(responseBody, &result); err != nil {
		return "", err
	}

	return result.Response, nil

}

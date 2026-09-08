package domain

type Answer struct {
	Answer  string           `json:"answer"`
	Sources []RetrievedChunk `json:"sources"`
}

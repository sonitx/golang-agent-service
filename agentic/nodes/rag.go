package nodes

import (
	"context"
	"main/models"
)

type AgenticRag struct {
}

func NewAgenticRag() AgenticInterface {
	return &AgenticRag{}
}

// GenerateResponse implements AgenticInterface.
func (a *AgenticRag) GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error) {
	panic("unimplemented")
}

package nodes

import (
	"context"
	"main/models"
)

type AgenticDirect struct {
}

func NewAgenticDirect() AgenticInterface {
	return &AgenticDirect{}
}

// GenerateResponse implements AgenticInterface.
func (a *AgenticDirect) GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error) {
	panic("unimplemented")
}

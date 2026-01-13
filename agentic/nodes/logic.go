package nodes

import (
	"context"
	"main/models"
)

type AgenticLogic struct {
}

func NewAgenticLogic() AgenticInterface {
	return &AgenticLogic{}
}

// GenerateResponse implements AgenticInterface.
func (a *AgenticLogic) GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error) {
	panic("unimplemented")
}

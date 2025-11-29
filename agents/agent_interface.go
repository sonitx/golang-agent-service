package agents

import (
	"context"
	"main/models"
)

type AgentInterface interface {
	GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error)
}

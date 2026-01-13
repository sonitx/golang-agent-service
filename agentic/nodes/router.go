package nodes

import (
	"context"
	"main/models"
)

type AgenticInterface interface {
	GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error)
}

package implements

import (
	"context"
	"main/agents"
	"main/models"
)

type AgentSearchWeb struct {
}

func NewAgentSearchWeb() agents.AgentInterface {
	return &AgentSearchWeb{}
}

// GenerateResponse implements agents.AgentInterface.
func (a *AgentSearchWeb) GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error) {
	return agents.GenerateResponse(ctx, prompt)
}

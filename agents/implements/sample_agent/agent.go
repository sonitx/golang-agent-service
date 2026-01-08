package sampleagent

import (
	"context"
	"fmt"
	"main/agents"
	"main/models"
)

type SampleAgent struct {
}

func NewSampleAgent() agents.AgentInterface {
	return &SampleAgent{}
}

// GenerateResponse implements agents.AgentInterface.
func (s *SampleAgent) GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error) {
	fullPrompt := fmt.Sprintf(instruction, prompt)
	return agents.GenerateResponse(ctx, fullPrompt, s.Tools(nil))
}

// Tools implements agents.AgentInterface.
func (s *SampleAgent) Tools(inputData any) []models.AITool {
	return []models.AITool{}
}

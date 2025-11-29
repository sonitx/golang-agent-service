package services

import (
	"context"
	"errors"
	"main/agents"
	"main/agents/implements"
	"main/models"
)

type AgentService struct{}

func NewAgentService() *AgentService {
	return &AgentService{}
}

func (s *AgentService) initAgent() map[string]agents.AgentInterface {
	mapAgent := make(map[string]agents.AgentInterface)
	mapAgent["agent_search_web"] = implements.NewAgentSearchWeb()

	return mapAgent
}

func (s *AgentService) GenerateResponse(ctx context.Context, agentKey string, prompt string) (*models.ChatResponse, error) {
	mapAgent := s.initAgent()
	if _, ok := mapAgent[agentKey]; !ok {
		return nil, errors.New("agent not found")
	}
	return mapAgent[agentKey].GenerateResponse(ctx, prompt)
}

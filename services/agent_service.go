package services

import (
	"context"
	"errors"
	"main/agents"
	"main/agents/cores"
	"main/agents/implements"
	"main/models"
	"main/utils"
)

type AgentService struct{}

func NewAgentService() *AgentService {
	return &AgentService{}
}

func (s *AgentService) initAgent() map[string]agents.AgentInterface {
	mapAgent := make(map[string]agents.AgentInterface)

	// AI AUTO GENERATED
	mapAgent["agent_search_web"] = cores.NewAgentSearchWeb()
	mapAgent["agent_sale"] = implements.NewSaleAgent()
	// END AI AUTO GENERATED

	return mapAgent
}

func (s *AgentService) GenerateResponse(ctx context.Context, agentKey string, prompt string) (*models.ChatResponse, error) {
	mapAgent := s.initAgent()
	if _, ok := mapAgent[agentKey]; !ok {
		return nil, errors.New("agent not found")
	}
	return mapAgent[agentKey].GenerateResponse(ctx, prompt)
}

func (s *AgentService) ListAgents() []string {
	mapAgent := s.initAgent()
	var agents []string
	for key := range mapAgent {
		agents = append(agents, key)
	}
	return agents
}
func (s *AgentService) ListAgentInfos() []struct {
	Key  string `json:"key"`
	Name string `json:"name"`
} {
	var infos []struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	}
	for _, item := range utils.AppConfig.AgentConfig {
		if item.Enable {
			infos = append(infos, struct {
				Key  string `json:"key"`
				Name string `json:"name"`
			}{Key: item.Key, Name: item.Name})
		}
	}
	return infos
}

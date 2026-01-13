package services

import (
	"context"
	"errors"
	"main/agents"
	"main/agents/cores"
	implements "main/agents/implements/sale_agent"
	"main/agents/implements/spvai"
	"main/models"
	"main/repositories"
	"main/utils"
)

type AgentService struct {
	sessionStatusRepo *repositories.SessionStatusRepository
}

func NewAgentService(sessionStatusRepo *repositories.SessionStatusRepository) *AgentService {
	return &AgentService{
		sessionStatusRepo: sessionStatusRepo,
	}
}

func (s *AgentService) initAgent() map[string]agents.AgentInterface {
	mapAgent := make(map[string]agents.AgentInterface)

	// AI AUTO GENERATED
	mapAgent["agent_search_web"] = cores.NewAgentSearchWeb()
	mapAgent["agent_sale"] = implements.NewSaleAgent()
	mapAgent["agent_spvai"] = spvai.NewSPVAIAgent(s.sessionStatusRepo)
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

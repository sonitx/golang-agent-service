package cores

import (
	"context"
	"fmt"
	"main/agents"
	"main/models"
	"main/utils"

	"github.com/firebase/genkit/go/ai"
)

type AgentSearchWeb struct {
}

func NewAgentSearchWeb() agents.AgentInterface {
	return &AgentSearchWeb{}
}

// GenerateResponse implements agents.AgentInterface.
func (a *AgentSearchWeb) GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error) {
	return agents.GenerateResponse(ctx, prompt, a.Tools(nil))
}

// Tools implements agents.AgentInterface.
func (a *AgentSearchWeb) Tools(inputData any) []models.AITool {
	return []models.AITool{
		a.searchGoogleTool(),
	}
}

func (a *AgentSearchWeb) searchGoogleTool() models.AITool {
	return models.AITool{
		Name:        "search_google",
		Description: "Search Google for information",
		Function: func(ctx *ai.ToolContext, input any) (string, error) {
			utils.ShowInfoLogs("Call search google tool with params: %s", input)
			return fmt.Sprintf("The result of question %s is Son Nguyen", input), nil
		},
	}
}

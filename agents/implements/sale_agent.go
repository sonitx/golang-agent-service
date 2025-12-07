package implements

import (
	"context"
	"main/agents"
	"main/agents/tools"
	"main/models"
)

type SaleAgent struct {
}

func NewSaleAgent() agents.AgentInterface {
	return &SaleAgent{}
}

// GenerateResponse implements agents.AgentInterface.
func (s *SaleAgent) GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error) {
	return agents.GenerateResponse(ctx, prompt, s.Tools())
}

// Tools implements agents.AgentInterface.
func (s *SaleAgent) Tools() []models.AITool {
	productsTool, err := tools.NewProductTool().GetListProducts()
	cartsTool, err := tools.NewCartsTool().GetListCarts()

	if err != nil {
		return nil
	}

	return []models.AITool{
		productsTool,
		cartsTool,
	}
}

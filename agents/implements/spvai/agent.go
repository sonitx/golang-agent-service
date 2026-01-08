package spvai

import (
	"context"
	"errors"
	"fmt"
	"main/agents"
	"main/models"
	"main/repositories"
	"main/utils"
)

type SPVAIAgent struct {
	sessionStatusRepo *repositories.SessionStatusRepository
}

func NewSPVAIAgent(sessionStatusRepo *repositories.SessionStatusRepository) agents.AgentInterface {
	return &SPVAIAgent{sessionStatusRepo: sessionStatusRepo}
}

// GenerateResponse implements agents.AgentInterface.
func (s *SPVAIAgent) GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error) {
	maxTime := 5
	dataResp := &models.ChatResponse{}

	for maxTime > 0 {
		// Convert text to sql
		utils.ShowInfoLogs("Retry remaining: %d", maxTime)
		sqlResp, err := s.ConvertTextToSQLQuery(ctx, prompt)
		if err != nil {
			maxTime--
			continue
		}
		utils.ShowInfoLogs("ConvertTextToSQLQuery DONE: %s", sqlResp.Answer)

		// Validate SQL query
		if err := utils.ValidateSQLQuery(sqlResp.Answer); err != nil {
			maxTime--
			continue
		}
		utils.ShowInfoLogs("ValidateSQLQuery DONE")

		// Call DB to get data
		data, err := s.sessionStatusRepo.GetSessionStatus(ctx, sqlResp.Answer)
		if err != nil {
			maxTime--
			continue
		}
		utils.ShowInfoLogs("GetSessionStatus DONE")

		// Summary the data
		jData, err := utils.MarshalToString(data)
		if err != nil {
			maxTime--
			continue
		}
		fullPrompt := fmt.Sprintf(instructionSummary, prompt, jData)
		dataResp, err = agents.GenerateResponse(ctx, fullPrompt, s.Tools(jData))
		if err != nil {
			maxTime--
			continue
		}
		return dataResp, nil
	}
	return nil, errors.New("max time")
}

// Tools implements agents.AgentInterface.
func (s *SPVAIAgent) Tools(inputData any) []models.AITool {
	return nil //[]models.AITool{}
}

// AI Agent convert text to sql query
func (s *SPVAIAgent) ConvertTextToSQLQuery(ctx context.Context, prompt string) (*models.ChatResponse, error) {
	fullPrompt := fmt.Sprintf(instructionSQLQuerySessionStatus, prompt)
	return agents.GenerateResponse(ctx, fullPrompt, s.Tools(nil))
}

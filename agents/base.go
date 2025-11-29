package agents

import (
	"context"
	"errors"
	"main/models"
	"main/utils"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/firebase/genkit/go/plugins/ollama"
)

func GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error) {
	g, modelName := getModel(ctx, utils.AppConfig.ModelConfig)
	if g == nil {
		return nil, errors.New("model not found")
	}

	response, err := genkit.Generate(ctx, g,
		ai.WithPrompt(prompt),
		ai.WithModelName(modelName),
	)
	if err != nil {
		utils.ShowErrorLogs(err)
		return nil, err
	}
	return &models.ChatResponse{
		Answer:     response.Text(),
		TotalToken: response.Usage.TotalTokens,
	}, nil
}

func getModel(ctx context.Context, modelConfig utils.ModelConfig) (*genkit.Genkit, string) {
	var g *genkit.Genkit
	var modelName string

	if modelConfig.Gemini.Enable {
		g = genkit.Init(ctx, genkit.WithPlugins(&googlegenai.GoogleAI{
			APIKey: modelConfig.Gemini.APIKey,
		}))
		modelName = modelConfig.Gemini.ModelName
	} else if modelConfig.Ollama.Enable {
		g = genkit.Init(ctx, genkit.WithPlugins(
			&ollama.Ollama{
				ServerAddress: modelConfig.Ollama.ServerAddress,
			},
		))
		modelName = modelConfig.Ollama.ModelName
	}
	return g, modelName
}

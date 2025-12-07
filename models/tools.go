package models

import "github.com/firebase/genkit/go/ai"

type AITool struct {
	Name        string
	Description string
	Function    func(ctx *ai.ToolContext, input any) (string, error)
}

package tools

import (
	"fmt"
	"main/models"
	"main/utils"

	"github.com/firebase/genkit/go/ai"
)

type CartsTool struct {
}

func NewCartsTool() *CartsTool {
	return &CartsTool{}
}

func (t *CartsTool) GetListCarts() (models.AITool, error) {
	utils.ShowInfoLogs("Call get list carts tool")

	return models.AITool{
		Name:        "get_carts",
		Description: "Get list carts from database",
		Function: func(ctx *ai.ToolContext, input any) (string, error) {
			fullUrl := "https://dummyjson.com/carts?limit=5"
			code, data := utils.DoGet(fullUrl, nil, nil)
			if code != 200 {
				return "", fmt.Errorf("get list carts failed")
			}
			return string(data), nil
		},
	}, nil
}

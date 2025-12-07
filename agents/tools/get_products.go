package tools

import (
	"fmt"
	"main/models"
	"main/utils"

	"github.com/firebase/genkit/go/ai"
)

type ProductTool struct {
}

func NewProductTool() *ProductTool {
	return &ProductTool{}
}

func (t *ProductTool) GetListProducts() (models.AITool, error) {
	utils.ShowInfoLogs("Call get list products tool")

	return models.AITool{
		Name:        "get_products",
		Description: "Get list products from database",
		Function: func(ctx *ai.ToolContext, input any) (string, error) {
			fullUrl := "https://dummyjson.com/products?limit=5"
			code, data := utils.DoGet(fullUrl, nil, nil)
			if code != 200 {
				return "", fmt.Errorf("get list products failed")
			}
			return string(data), nil
		},
	}, nil
}

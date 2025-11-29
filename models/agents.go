package models

type ChatRequest struct {
	Question string `json:"question"`
}

type ChatResponse struct {
	Answer     string `json:"answer"`
	TotalToken int    `json:"total_token,omitempty"`
}

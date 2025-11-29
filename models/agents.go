package models

type ChatRequest struct {
	Question string `json:"question"`
}

type ChatResponse struct {
	Answer   string `json:"answer"`
	SQLQuery string `json:"sql_query,omitempty"`
}

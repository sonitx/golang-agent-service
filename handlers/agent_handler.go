package handlers

import (
	"encoding/json"
	"fmt"
	"main/api"
	"main/models"
	"main/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AgentHandler struct {
	agentSvc *services.AgentService
}

func NewAgentHandler(agentSvc *services.AgentService) *AgentHandler {
	return &AgentHandler{
		agentSvc: agentSvc,
	}
}

func (h *AgentHandler) GenerateResponse(w http.ResponseWriter, r *http.Request) {
	agentKey := chi.URLParam(r, "agent")
	if agentKey == "" {
		api.Error(w, http.StatusBadRequest, "agent is required")
		return
	}

	var reqBody models.ChatRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		api.Error(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	agentResp, err := h.agentSvc.GenerateResponse(r.Context(), agentKey, reqBody.Question)
	if err != nil {
		api.Error(w, http.StatusInternalServerError, fmt.Sprintf("Failed to generate response: %v", err))
		return
	}

	api.Ok(w, agentResp)
}

package server

import "github.com/go-chi/chi/v5"

func (apiServer *apiServer) initApiRoutes(r *chi.Mux) {
	// Basic API
	r.Get("/", apiServer.baseHandler.Home)
	r.Get("/ping", apiServer.baseHandler.Ping)

	// Agent API
	r.Get("/api/agents", apiServer.agentHandler.ListAgents)
	r.Post("/api/{agent}/generate", apiServer.agentHandler.GenerateResponse)
}

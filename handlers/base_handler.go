package handlers

import (
	"main/api"
	"net/http"
)

type BaseHandler struct{}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) Home(w http.ResponseWriter, r *http.Request) {
	api.Ok(w, "Welcome to Chi Router Framework!")
}

func (h *BaseHandler) Ping(w http.ResponseWriter, r *http.Request) {
	api.Ok(w, "pong")
}

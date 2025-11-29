package api

import (
	"encoding/json"
	"fmt"
	"main/utils"
	"net/http"
)

type apiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func getMessage(code int) string {
	strMessage := ""
	switch code {
	case http.StatusOK:
		strMessage = "Success"
	case http.StatusBadRequest:
		strMessage = "Bad request"
	case http.StatusUnauthorized:
		strMessage = "Unauthorized"
	case http.StatusForbidden:
		strMessage = "You don't have permission to access"
	case http.StatusInternalServerError:
		strMessage = "Internal server error"
	case http.StatusBadGateway:
		strMessage = "Bad gateway"
	default:
		strMessage = "Unknown error"
	}
	return strMessage
}

func Ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	apiResp := apiResponse{
		Code:    http.StatusOK,
		Message: getMessage(http.StatusOK),
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(apiResp); err != nil {
		utils.ShowErrorLogs(fmt.Errorf("Error encoding response: %v", err))
		http.Error(w, getMessage(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if message == "" {
		message = getMessage(code)
	}
	utils.ShowErrorLogs(fmt.Errorf("Error code: %d, meg: %s", code, message))

	apiResp := apiResponse{
		Code:    code,
		Message: getMessage(code),
		Data:    nil,
	}

	if err := json.NewEncoder(w).Encode(apiResp); err != nil {
		utils.ShowErrorLogs(fmt.Errorf("Error encoding response: %v", err))
		http.Error(w, getMessage(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

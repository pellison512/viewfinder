package handlers

import (
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() HealthHandler {
	return HealthHandler{}
}

func (handler *HealthHandler) HealthGETHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

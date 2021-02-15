package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// HealthCheck handle health check
func (eh *ExampleHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, eh.service.CheckHealth())
}

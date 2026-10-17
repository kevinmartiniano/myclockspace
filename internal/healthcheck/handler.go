package healthcheck

import (
	"net/http"

	"myclockspace.kevinmartiniano.com/utils"
)

type HealthcheckHandler struct{}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) Index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string `json:"message"`
	}{Message: "Ok!"}

	utils.JSONResponse(w, http.StatusOK, data)
}

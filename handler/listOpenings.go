package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//@BasePath /api/v1

// @Summary List openings
// @Description Listt all jobs openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func (h *Handler) ListOpeningsHandler(ctx *gin.Context) {
	openings, err := h.db.List()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marceloxhenrique/gopportunities/schemas"
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
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}

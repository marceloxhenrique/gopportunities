package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marceloxhenrique/gopportunities/schemas"
)

//@BasePath /api/v1

// @Summary Show Opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string  true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func (h *Handler) ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := h.db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening with not found")
		return
	}

	sendSuccess(ctx, "showing-opening", opening)
}

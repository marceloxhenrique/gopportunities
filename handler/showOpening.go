package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	idParam := ctx.Query("id")
	if idParam == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	idUint64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	opening, err := h.db.GetById(uint(idUint64))
	if err != nil {
		sendError(ctx, http.StatusNotFound, "opening with not found")
		return
	}

	sendSuccess(ctx, "showing-opening", opening)
}

package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//@BasePath /api/v1

// @Summary Delete Opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string  true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func (h *Handler) DeleteOpeningHandler(ctx *gin.Context) {
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
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", idParam))
		return
	}

	if err := h.db.Delete(uint(idUint64)); err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", idParam))
		return
	}
	sendSuccess(ctx, "deleting-opening", opening)

}

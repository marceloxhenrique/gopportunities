package handler

import (
	"github.com/marceloxhenrique/gopportunities/config"
	"gorm.io/gorm"
)

type Handler struct {
	logger *config.Logger
	db     *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db:     db,
		logger: config.GetLogger("handler"),
	}
}

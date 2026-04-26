package handler

import (
	"github.com/marceloxhenrique/gopportunities/config"
	"github.com/marceloxhenrique/gopportunities/repository"
)

type Handler struct {
	logger *config.Logger
	db     repository.OpeningRepository
}

func NewHandler(repo repository.OpeningRepository) *Handler {
	return &Handler{
		db:     repo,
		logger: config.GetLogger("handler"),
	}
}

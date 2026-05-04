package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/marceloxhenrique/gopportunities/repository"
	"github.com/marceloxhenrique/gopportunities/schemas"
)

func TestDeleteOpening(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewInMemoryRepository()
	h := NewHandler(repo)

	r := gin.Default()
	r.DELETE("/opening", h.DeleteOpeningHandler)

	opening := &schemas.Opening{
		Role:     "Backend Dev",
		Company:  "Google",
		Location: "Remote",
		Remote:   true,
		Salary:   5000,
		Link:     "www.google.com/opening",
	}
	repo.Create(opening)

	assert.Equal(t, "Backend Dev", repo.Data[1].Role)

	req := httptest.NewRequest("DELETE", "/opening?id=1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	_, exist := repo.Data[1]
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, false, exist)
}

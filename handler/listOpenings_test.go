package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/marceloxhenrique/gopportunities/repository"
	"github.com/marceloxhenrique/gopportunities/schemas"
)

func TestListOpeningsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := repository.NewInMemoryRepository()
	h := NewHandler(repo)
	r := gin.Default()

	r.GET("/openings", h.ListOpeningsHandler)

	opening := &schemas.Opening{
		Role:     "Backend Dev",
		Company:  "Google",
		Location: "Remote",
		Remote:   true,
		Salary:   5000,
		Link:     "www.google.com/opening",
	}
	repo.Create(opening)

	req := httptest.NewRequest("GET", "/openings", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	type ListResponse struct {
		Data    []schemas.Opening `json:"data"`
		Message string            `json:"message"`
	}
	var response ListResponse

	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, "Backend Dev", response.Data[0].Role)
	assert.Equal(t, "Google", response.Data[0].Company)
	assert.Equal(t, int64(5000), response.Data[0].Salary)

}

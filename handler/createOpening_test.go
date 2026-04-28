package handler

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/marceloxhenrique/gopportunities/repository"
)

func TestCreateOpeningHandler_Router(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewInMemoryRepository()
	h := NewHandler(repo)

	r := gin.Default()
	r.POST("/opening", h.CreateOpeningHandler)

	body := `{
		"role": "Backend Dev",
		"company": "Google",
		"location": "Remote",
		"remote":true,
		"salary": 5000,
		"link": "www.google.com/opening"
	}`

	req := httptest.NewRequest("POST", "/opening", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(repo.Data))
	assert.Equal(t, "Backend Dev", repo.Data[1].Role)
	assert.Equal(t, "Google", repo.Data[1].Company)
	assert.Equal(t, int64(5000), repo.Data[1].Salary)

}

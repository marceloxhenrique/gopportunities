package handler

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/marceloxhenrique/gopportunities/repository"
	"github.com/marceloxhenrique/gopportunities/schemas"
)

func TestUpdateOpening(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewInMemoryRepository()
	h := NewHandler(repo)

	r := gin.Default()
	r.PUT("/opening", h.UpdateOpeningHandler)

	opening := &schemas.Opening{
		Role:     "Backend Dev",
		Company:  "Google",
		Location: "Remote",
		Remote:   true,
		Salary:   5000,
		Link:     "www.google.com/opening",
	}
	repo.Create(opening)

	body := `{
		"role": "Backend Dev",
		"company": "IBM",
		"location": "Remote",
		"remote":true,
		"salary": 5000,
		"link": "www.ibm.com/opening"
	}`

	req := httptest.NewRequest("PUT", "/opening?id=1", strings.NewReader(body))

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	type Response struct {
		Data    schemas.Opening `json:"data"`
		Message string          `json:"message"`
	}

	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)

	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, repo.Data[1].Company, response.Data.Company)
	assert.Equal(t, repo.Data[1].Link, response.Data.Link)
}

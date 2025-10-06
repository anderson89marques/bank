package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anderson89marques/bank/config"
	"github.com/anderson89marques/bank/internal/adapter/database/postgres/repository"
	"github.com/anderson89marques/bank/internal/core/services"
	"github.com/anderson89marques/bank/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountThenReturnSuccess(t *testing.T) {
	db, err := tests.SetupTests(t)
	assert.NoError(t, err)

	// setup test running migrations
	payload := AccountInputSchema{
		Document: "123458",
	}
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	basePath := config.GetEnv().BasePath
	baseGroup := router.Group(basePath)

	v1 := baseGroup.Group("/api/v1")

	// Accounts
	accountRepo := repository.NewAccountRepository(db)
	accountSrv := services.NewAccountService(accountRepo)
	accountHandler := NewAccountHandler(accountSrv)

	accountsGroup := v1.Group("/accounts")
	accountsGroup.POST("/", accountHandler.Create)

	endpoint := "/api/v1/accounts/"
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	assert.NoError(t, err)
	router.ServeHTTP(recorder, r)

	var response map[string]any
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}

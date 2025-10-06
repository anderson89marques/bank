package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anderson89marques/bank/config"
	"github.com/anderson89marques/bank/internal/adapter/database/postgres/repository"
	"github.com/anderson89marques/bank/internal/core/domain"
	"github.com/anderson89marques/bank/internal/core/services"
	"github.com/anderson89marques/bank/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionWithNormalPurchaseThenReturnSuccess(t *testing.T) {
	db, err := tests.SetupTests(t)
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	basePath := config.GetEnv().BasePath
	baseGroup := router.Group(basePath)

	v1 := baseGroup.Group("/api/v1")

	// Creating accounting
	accountRepo := repository.NewAccountRepository(db)
	accountSrv := services.NewAccountService(accountRepo)
	account := domain.Account{
		Document: "123459",
	}
	_, err = accountSrv.Create(context.Background(), &account)
	assert.NoError(t, err)

	operationTypeRepo := repository.NewOperationTypeRepository(db)
	operationTypeService := services.NewOperationTypeService(operationTypeRepo)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionSrv := services.NewTransactionService(transactionRepo, operationTypeService)
	transactionHandler := NewTransacationHandler(transactionSrv)

	transactionsGroup := v1.Group("/transactions")
	transactionsGroup.POST("/", transactionHandler.Create)

	endpoint := "/api/v1/transactions/"
	// setup test running migrations
	payload := TransactionInputSchema{
		AccountID:       account.ID,
		OperationTypeID: 1,
		Amount:          123.45,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	assert.NoError(t, err)
	router.ServeHTTP(recorder, r)

	var response map[string]any
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, -123.45, response["amount"])
}

func TestCreateTransactionWithPurchaseInstallmentsThenReturnSuccess(t *testing.T) {
	db, err := tests.SetupTests(t)
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	basePath := config.GetEnv().BasePath
	baseGroup := router.Group(basePath)

	v1 := baseGroup.Group("/api/v1")

	// Creating accounting
	accountRepo := repository.NewAccountRepository(db)
	accountSrv := services.NewAccountService(accountRepo)
	account := domain.Account{
		Document: "123459",
	}
	_, err = accountSrv.Create(context.Background(), &account)
	assert.NoError(t, err)

	operationTypeRepo := repository.NewOperationTypeRepository(db)
	operationTypeService := services.NewOperationTypeService(operationTypeRepo)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionSrv := services.NewTransactionService(transactionRepo, operationTypeService)
	transactionHandler := NewTransacationHandler(transactionSrv)

	transactionsGroup := v1.Group("/transactions")
	transactionsGroup.POST("/", transactionHandler.Create)

	endpoint := "/api/v1/transactions/"
	// setup test running migrations
	payload := TransactionInputSchema{
		AccountID:       account.ID,
		OperationTypeID: 2,
		Amount:          123.45,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	assert.NoError(t, err)
	router.ServeHTTP(recorder, r)

	var response map[string]any
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, -123.45, response["amount"])
}

func TestCreateTransactionWithWithDrawalThenReturnSuccess(t *testing.T) {
	db, err := tests.SetupTests(t)
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	basePath := config.GetEnv().BasePath
	baseGroup := router.Group(basePath)

	v1 := baseGroup.Group("/api/v1")

	// Creating accounting
	accountRepo := repository.NewAccountRepository(db)
	accountSrv := services.NewAccountService(accountRepo)
	account := domain.Account{
		Document: "123459",
	}
	_, err = accountSrv.Create(context.Background(), &account)
	assert.NoError(t, err)

	operationTypeRepo := repository.NewOperationTypeRepository(db)
	operationTypeService := services.NewOperationTypeService(operationTypeRepo)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionSrv := services.NewTransactionService(transactionRepo, operationTypeService)
	transactionHandler := NewTransacationHandler(transactionSrv)

	transactionsGroup := v1.Group("/transactions")
	transactionsGroup.POST("/", transactionHandler.Create)

	endpoint := "/api/v1/transactions/"
	// setup test running migrations
	payload := TransactionInputSchema{
		AccountID:       account.ID,
		OperationTypeID: 3,
		Amount:          123.45,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	assert.NoError(t, err)
	router.ServeHTTP(recorder, r)

	var response map[string]any
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, -123.45, response["amount"])
}

func TestCreateTransactionWithCreditThenReturnSuccess(t *testing.T) {
	db, err := tests.SetupTests(t)
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	basePath := config.GetEnv().BasePath
	baseGroup := router.Group(basePath)

	v1 := baseGroup.Group("/api/v1")

	// Creating accounting
	accountRepo := repository.NewAccountRepository(db)
	accountSrv := services.NewAccountService(accountRepo)
	account := domain.Account{
		Document: "123459",
	}
	_, err = accountSrv.Create(context.Background(), &account)
	assert.NoError(t, err)

	operationTypeRepo := repository.NewOperationTypeRepository(db)
	operationTypeService := services.NewOperationTypeService(operationTypeRepo)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionSrv := services.NewTransactionService(transactionRepo, operationTypeService)
	transactionHandler := NewTransacationHandler(transactionSrv)

	transactionsGroup := v1.Group("/transactions")
	transactionsGroup.POST("/", transactionHandler.Create)

	endpoint := "/api/v1/transactions/"
	// setup test running migrations
	payload := TransactionInputSchema{
		AccountID:       account.ID,
		OperationTypeID: 4,
		Amount:          123.45,
	}

	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	assert.NoError(t, err)
	router.ServeHTTP(recorder, r)

	var response map[string]any
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, 123.45, response["amount"])
}

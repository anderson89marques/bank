package rest

import (
	"github.com/anderson89marques/bank/config"
	database "github.com/anderson89marques/bank/internal/adapter/database/postgres"
	"github.com/anderson89marques/bank/internal/adapter/database/postgres/repository"
	"github.com/anderson89marques/bank/internal/core/services"
	"github.com/gin-gonic/gin"

	docs "github.com/anderson89marques/bank/docs"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func swaggerInfo() {
	docs.SwaggerInfo.BasePath = "/bank"
	docs.SwaggerInfo.Title = "Bank"
	docs.SwaggerInfo.Description = "Bank exchange"
	docs.SwaggerInfo.Version = "1.0.0"
}

func RegisterRoutes(engine gin.IRouter) {
	db, err := database.NewPostgresDB(config.GetEnv())
	if err != nil {
		panic(err)
	}
	swaggerInfo()
	basePath := config.GetEnv().BasePath
	baseGroup := engine.Group(basePath)
	baseGroup.GET("/docs/*any", ginSwagger.WrapHandler(swagFiles.Handler))

	v1 := baseGroup.Group("/api/v1")

	// Accounts
	accountRepo := repository.NewAccountRepository(db)
	accountSrv := services.NewAccountService(accountRepo)
	accountHandler := NewAccountHandler(accountSrv)

	accountsGroup := v1.Group("/accounts")
	accountsGroup.POST("/", accountHandler.Create)
	accountsGroup.GET("/", accountHandler.List)
	accountsGroup.GET("/:id", accountHandler.FindByID)

	// Transactions
	operationTypeRepo := repository.NewOperationTypeRepository(db)
	operationTypeService := services.NewOperationTypeService(operationTypeRepo)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionSrv := services.NewTransactionService(transactionRepo, operationTypeService)
	transactionHandler := NewTransacationHandler(transactionSrv)

	transactionsGroup := v1.Group("/transactions")
	transactionsGroup.POST("/", transactionHandler.Create)
	transactionsGroup.GET("/", transactionHandler.List)
}

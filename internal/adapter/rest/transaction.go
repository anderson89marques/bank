package rest

import (
	"net/http"

	"github.com/anderson89marques/bank/internal/core/domain"
	"github.com/anderson89marques/bank/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	srv ports.TransactionService
}

func NewTransacationHandler(srv ports.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		srv: srv,
	}
}

type TransactionInputSchema struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type TransactionOutputSchema struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func TransactionOutputFromDomain(transaction *domain.Transaction) TransactionOutputSchema {
	return TransactionOutputSchema{
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          transaction.Amount,
	}
}

type TransactionList struct {
	Data []TransactionOutputSchema `json:"data"`
}

func TransactionListFromListDomain(transactions []*domain.Transaction) TransactionList {
	var output []TransactionOutputSchema
	for _, a := range transactions {
		transactionOut := TransactionOutputFromDomain(a)
		output = append(output, transactionOut)
	}
	return TransactionList{
		Data: output,
	}
}

// @Summary Banking
// @Description create bank transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param TransactionInputSchema body TransactionInputSchema true "Create Transaction"
// @Success 200 {object} TransactionInputSchema
// @Router /api/v1/transactions [post]
func (h *TransactionHandler) Create(ctx *gin.Context) {
	var transactionPayload TransactionInputSchema
	if err := ctx.ShouldBindJSON(&transactionPayload); err != nil {
		ctx.PureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := domain.Transaction{
		AccountID:       transactionPayload.AccountID,
		OperationTypeID: transactionPayload.OperationTypeID,
		Amount:          transactionPayload.Amount,
	}

	_, err := h.srv.Create(ctx, &transaction)
	if err != nil {
		ctx.PureJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	transactionResponsePayload := TransactionOutputFromDomain(&transaction)
	ctx.PureJSON(http.StatusCreated, transactionResponsePayload)
}

// @Summary Banking
// @Description List Transactions
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {object} TransactionList
// @Router /api/v1/transactions [get]
func (h *TransactionHandler) List(ctx *gin.Context) {
	accounts, err := h.srv.List(ctx)
	if err != nil {
		ctx.PureJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	transactionResponsePayload := TransactionListFromListDomain(accounts)
	ctx.PureJSON(http.StatusOK, transactionResponsePayload)
}

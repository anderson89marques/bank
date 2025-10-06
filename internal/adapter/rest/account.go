package rest

import (
	"net/http"

	"github.com/anderson89marques/bank/internal/core/domain"
	"github.com/anderson89marques/bank/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	srv ports.AccountService
}

func NewAccountHandler(srv ports.AccountService) *AccountHandler {
	return &AccountHandler{
		srv: srv,
	}
}

type AccountInputSchema struct {
	Document string `json:"document"`
}

type AccountFindByIDSchema struct {
	ID int `uri:"id"`
}

type AccountOutputSchema struct {
	ID       int    `json:"id"`
	Document string `json:"document"`
}

func AccountOutputFromDomain(account *domain.Account) AccountOutputSchema {
	return AccountOutputSchema{
		ID:       account.ID,
		Document: account.Document,
	}
}

type AccountList struct {
	Data []AccountOutputSchema `json:"data"`
}

func AccountListFromListDomain(accounts []*domain.Account) AccountList {
	var output []AccountOutputSchema
	for _, a := range accounts {
		accountOut := AccountOutputFromDomain(a)
		output = append(output, accountOut)
	}
	return AccountList{
		Data: output,
	}
}

// @Summary Banking
// @Description create bank account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param AccountInputSchema body AccountInputSchema true "Create Account"
// @Success 200 {object} AccountOutputSchema
// @Router /api/v1/accounts [post]
func (h *AccountHandler) Create(ctx *gin.Context) {
	var accountPayload AccountInputSchema
	if err := ctx.ShouldBindJSON(&accountPayload); err != nil {
		ctx.PureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := domain.Account{
		Document: accountPayload.Document,
	}

	_, err := h.srv.Create(ctx, &account)
	if err != nil {
		ctx.PureJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	accountResponsePayload := AccountOutputFromDomain(&account)
	ctx.PureJSON(http.StatusCreated, accountResponsePayload)
}

// @Summary Banking
// @Description get account by id
// @Tags Accounts
// @Accept json
// @Produce json
// @Param id path uint64 true "Find Account by Id"
// @Success 200 {object} AccountOutputSchema
// @Router /api/v1/accounts/{id} [get]
func (h *AccountHandler) FindByID(ctx *gin.Context) {
	var accountFindPayload AccountFindByIDSchema
	if err := ctx.ShouldBindUri(&accountFindPayload); err != nil {
		ctx.PureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := h.srv.FindByID(ctx, accountFindPayload.ID)
	if err != nil {
		ctx.PureJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	accountResponsePayload := AccountOutputFromDomain(account)
	ctx.PureJSON(http.StatusOK, accountResponsePayload)
}

// @Summary Banking
// @Description List Accounts
// @Tags Accounts
// @Accept json
// @Produce json
// @Success 200 {object} AccountList
// @Router /api/v1/accounts [get]
func (h *AccountHandler) List(ctx *gin.Context) {
	accounts, err := h.srv.List(ctx)
	if err != nil {
		ctx.PureJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	accountResponsePayload := AccountListFromListDomain(accounts)
	ctx.PureJSON(http.StatusOK, accountResponsePayload)
}

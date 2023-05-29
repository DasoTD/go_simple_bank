package api

import (
	"net/http"

	db "github.com/dasotd/go_simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createEntryRequest struct {
	Amount    int64 `json:"amount" binding:"required"`
	AccountID int64 `json:"accountId" binding:"required"`
}

func (server *Server) createEntry (ctx *gin.Context) {
	var req createEntryRequest
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEntryParams{
		Amount:    req.Amount,
		AccountID: req.AccountID,
	}

	entry, err := server.store.CreateEntry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, entry)
}
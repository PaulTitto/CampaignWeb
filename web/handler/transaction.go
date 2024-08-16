package handler

import (
	"campaignweb/transaction"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTranscationHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService: transactionService}
}

func (h *transactionHandler) Index(c *gin.Context) {
	transactions, err := h.transactionService.GetAllTransaction()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "transcation_index.html", gin.H{"transactions": transactions})

}

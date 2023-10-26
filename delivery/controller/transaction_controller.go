package controller

import (
	"net/http"
	"payment-apps-backend/delivery/middleware"
	"payment-apps-backend/model"
	"payment-apps-backend/usecase"
	"payment-apps-backend/utils/common"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	router     *gin.Engine
	transactionUC usecase.TransactionUseCase
}

func (t *TransactionController) createHandler(gc *gin.Context) {
	var transaction model.Transaction
	if err := gc.ShouldBindJSON(&transaction); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	transaction.ID = common.GenerateID()
	if err := t.transactionUC.RegisterTransaction(transaction); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	userResponse := map[string]any{
		"id":    transaction.ID,
		"customer_id":  transaction.CustomerID,
		"merchant_id": transaction.MerchantID,
		"amount": transaction.Amount,

	}

	gc.JSON(http.StatusOK, userResponse)
}

func (t *TransactionController) listHandler(gc *gin.Context) {
	transaction, err := t.transactionUC.FindAllTransaction()
	if err != nil {
		gc.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	gc.JSON(200, gin.H{
		"status": status,
		"data":   transaction,
	})
}

func (t *TransactionController) getHandler(gc *gin.Context) {
	id := gc.Param("id")
	transaction, err := t.transactionUC.FindByIdMerch(id)
	if err != nil {
		gc.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get By Id Data Successfully",
	}
	gc.JSON(200, gin.H{
		"status": status,
		"data":   transaction,
	})
}

func NewTransactionController(router *gin.Engine, usecase usecase.TransactionUseCase) *TransactionController {
	controller := TransactionController{
		router:     router,
		transactionUC: usecase,
	}

	rg := router.Group("/api/v1")

	rg.POST("/transactions", middleware.AuthMiddleware("1"), controller.createHandler)
	rg.GET("/transactions/", middleware.AuthMiddleware("1"), controller.listHandler)
	rg.GET("/transactions/:id", middleware.AuthMiddleware("1"), controller.getHandler)

	return &controller
}

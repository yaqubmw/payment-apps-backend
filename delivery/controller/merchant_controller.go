package controller

import (
	"net/http"
	"payment-apps-backend/model"
	"payment-apps-backend/usecase"
	"payment-apps-backend/utils/common"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	router     *gin.Engine
	merchantUC usecase.MerchantUseCase
}

func (m *MerchantController) createHandler(gc *gin.Context) {
	var merchant model.Merchant
	if err := gc.ShouldBindJSON(&merchant); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	merchant.ID = common.GenerateID()
	if err := m.merchantUC.RegisterMerchant(merchant); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	userResponse := map[string]any{
		"id":    merchant.ID,
		"name":  merchant.Name,
		"email": merchant.Email,
	}

	gc.JSON(http.StatusOK, userResponse)
}

func (m *MerchantController) getHandler(gc *gin.Context) {
	id := gc.Param("id")
	merchant, err := m.merchantUC.FindByIdMerch(id)
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
		"data":   merchant,
	})
}

func NewMerchantController(router *gin.Engine, usecase usecase.MerchantUseCase) *MerchantController {
	controller := MerchantController{
		router:     router,
		merchantUC: usecase,
	}

	rg := router.Group("/api/v1")

	rg.POST("/merchants", controller.createHandler)
	rg.GET("/merchants/:id", controller.getHandler)

	return &controller
}

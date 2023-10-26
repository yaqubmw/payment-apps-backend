package controller

import (
	"net/http"
	"payment-apps-backend/delivery/middleware"
	"payment-apps-backend/model"
	"payment-apps-backend/usecase"
	"payment-apps-backend/utils/common"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router     *gin.Engine
	customerUC usecase.CustomerUseCase
}

func (c *CustomerController) createHandler(gc *gin.Context) {
	var customer model.Customer
	if err := gc.ShouldBindJSON(&customer); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	customer.ID = common.GenerateID()
	if err := c.customerUC.RegisterCustomer(customer); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	userResponse := map[string]any{
		"id":       customer.ID,
		"name":     customer.Name,
		"username": customer.Username,
		"role_id":   customer.RoleID,
	}

	gc.JSON(http.StatusOK, userResponse)
}

func (c *CustomerController) updateHandler(gc *gin.Context) {
	var customer model.Customer
	if err := gc.ShouldBindJSON(&customer); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := c.customerUC.UpdateCustomer(customer); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	userResponse := map[string]any{
		"id":       customer.ID,
		"name":     customer.Name,
		"username": customer.Username,
		"roleid":   customer.RoleID,
	}

	gc.JSON(http.StatusOK, userResponse)
}

func NewCustomerController(router *gin.Engine, usecase usecase.CustomerUseCase) *CustomerController {
	controller := CustomerController{
		router:     router,
		customerUC: usecase,
	}

	rg := router.Group("/api/v1")

	rg.POST("/customers", controller.createHandler)
	rg.PUT("/dashboard/customers", middleware.AuthMiddleware("1"), controller.updateHandler)

	return &controller
}

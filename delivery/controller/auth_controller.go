package controller

import (
	"net/http"
	"payment-apps-backend/model"
	"payment-apps-backend/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router  *gin.Engine
	usecase usecase.AuthUseCase
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload model.Customer
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	token, err := a.usecase.Login(payload.Username, payload.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func NewAuthController(router *gin.Engine, usecase usecase.AuthUseCase) *AuthController {
	controller := AuthController{
		router:  router,
		usecase: usecase,
	}
	rg := router.Group("/api/v1")
	rg.POST("/customers/login", controller.loginHandler)
	return &controller
}

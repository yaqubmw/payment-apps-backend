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
	var sessions = make(map[string]bool)

	rg := router.Group("/api/v1")
	rg.POST("/customers/login", controller.loginHandler)

	rg.POST("/customers/logout", func(c *gin.Context) {
		sessionToken := c.GetHeader("Authorization")
		if _, exists := sessions[sessionToken]; exists {
			delete(sessions, sessionToken)
			c.JSON(200, gin.H{"message": "Logged out successfully"})
		} else {
			c.JSON(401, gin.H{"message": "Invalid session or token"})
		}
	})

	return &controller
}

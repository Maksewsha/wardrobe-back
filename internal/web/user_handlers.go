package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maksewsha/wardrobe-back/internal/model"
	"github.com/maksewsha/wardrobe-back/internal/usecases"
)

type userHandlers struct {
	userUseCase usecases.User
}

func newUserHandlers(supergroup gin.RouterGroup, u usecases.User) {
	handler := userHandlers{u}
	userGroup := supergroup.Group("/user")
	{
		userGroup.GET("/:login", handler.viewUser)
		userGroup.POST("/register", handler.registerUser)
	}
}

func (uh *userHandlers) viewUser(c *gin.Context) {
	resp, err := uh.userUseCase.GetUserByLogin(context.Background(), c.Param("login"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (uh *userHandlers) registerUser(c *gin.Context) {
	var body model.UserAuthRequest
	if c.BindJSON(&body) != nil {
		err := fmt.Errorf("Unable to get valid body")
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	resp, err := uh.userUseCase.CreateNewUser(context.Background(), body)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully registered!",
		"id":      resp,
	})
}

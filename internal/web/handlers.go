package web

import (
	"github.com/gin-gonic/gin"
	"github.com/maksewsha/wardrobe-back/internal/usecases"
)

type Handlers struct {
	userHandlers
}

func NewHandlers(u usecases.User) *Handlers {
	return &Handlers{
		userHandlers: userHandlers{
			userUseCase: u,
		},
	}
}

func (h *Handlers) SetRouter(r *gin.Engine) {
	r.Use(gin.Recovery())
	supergroup := r.Group("/api")
	{
		newUserHandlers(*supergroup, h.userUseCase)
	}
}

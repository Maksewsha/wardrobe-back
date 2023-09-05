package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maksewsha/wardrobe-back/internal/model"
)

func AddCloth(ctx *gin.Context) {
	var reqEntity model.ClothDTO
	if error := ctx.ShouldBindJSON(&reqEntity); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": error.Error(),
		})
	} else {
		ctx.Status(http.StatusOK)
	}
}

func GetClothes(ctx *gin.Context) {

}

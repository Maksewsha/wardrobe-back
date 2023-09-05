package web

import (
	"github.com/gin-gonic/gin"
)

func RunRouter() (err error) {
	r := gin.Default()
	r.POST("/clothes", AddCloth)
	if error := r.Run(); error != nil {
		return err
	}
	return
}

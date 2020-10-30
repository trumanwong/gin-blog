package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {

}

func (this *Controller) Response(context *gin.Context, data interface{}, code int, message string) {
	context.JSON(code, gin.H{
		"message": message,
		"data": data,
	})
}
package util

import (
	"gin-blog/pkg/settings"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(context *gin.Context) int {
	result := 0
	page, _ := com.StrTo(context.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * settings.PageSize
	}
	return result
}
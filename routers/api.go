package routers

import (
	"gin-blog/app/http/controllers"
	"gin-blog/pkg/settings"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(settings.RunMode)

	router.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "test",
		})
	})

	api := router.Group("/api")
	{
		tag := api.Group("/tags")
		{
			tagController := controllers.Tag{}
			//获取标签列表
			tag.GET("/", tagController.Get)
			//新建标签
			tag.POST("/", tagController.Add)
			//更新指定标签
			tag.PUT("/:id", tagController.Update)
			//删除指定标签
			tag.DELETE("/:id", tagController.Delete)
		}
	}
	return router
}
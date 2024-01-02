package router

import (
	"Demo/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	publicAdminRouter := Router.Group("admin")
	{
		publicAdminRouter.POST("login", services.Login)
		publicAdminRouter.POST("countType", services.CountByType)
		publicAdminRouter.POST("countGrade", services.CountByGrade)
		publicAdminRouter.POST("countDate", services.CountDate)
		publicAdminRouter.POST("showImage", services.ShowImage)
		publicAdminRouter.POST("checkType", services.CheckType)
		publicAdminRouter.POST("upload", services.UploadCSV)
	}
	return Router
}

package routes

import (
	"todo-app/pkg/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Serve Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		v1.POST("/todo", controllers.CreateATodo)
		v1.GET("/todo/:id", controllers.GetATodo)
		v1.PUT("/todo/:id", controllers.UpdateATodo)
		v1.DELETE("/todo/:id", controllers.DeleteATodo)
	}

	return r
}

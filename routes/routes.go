
package routes

import (
	"gatorshare/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		posts := v1.Group("/posts"); {
			posts.GET("getAll/:userId", controllers.Listpost)
			posts.GET("getOne/:id", controllers.GetOnepost)
			posts.POST("create", controllers.AddNewpost)
			posts.PATCH("update/:id", controllers.UpdatePost)
			posts.DELETE("delete/:id", controllers.Deletepost)
		}
	}
}
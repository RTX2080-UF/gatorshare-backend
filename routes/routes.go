
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
		};
		comments := v1.Group("/comments"); {
			comments.GET("getAll/:userId", controllers.GetAllcomment)
			comments.GET("getOne/:id", controllers.GetOnecomment)
			comments.POST("create", controllers.AddNewcomment)
			comments.DELETE("delete/:id", controllers.Deletecomment)
			comments.PATCH("update/:id", controllers.Updatecomment)
		}
	}
}
package routes

import (
	"gatorshare/controllers"
	"gorm.io/gorm"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func InitializeRoutes(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	router.Use(gin.Logger())

	router.GET("/", func(responseCtx *gin.Context) {
	  responseCtx.JSON(http.StatusOK, gin.H{"data": "Welcome to Gatorshare made with the help of Go and Gin!"})    
	})

	router.Use(cors.Default())

	api := controllers.Controller{DB: db}

	v1 := router.Group("/v1")
	{
		posts := v1.Group("/posts"); {
			posts.GET("getAll/:userId", api.Listpost)
			posts.GET("getOne/:id", api.GetOnepost)
			posts.POST("create", api.AddNewpost)
			posts.PATCH("update/:id", api.UpdatePost)
			posts.DELETE("delete/:id", api.Deletepost)
		};
		comments := v1.Group("/comments"); {
			comments.GET("getAll/:postId", api.GetAllcomment)
			comments.GET("getOne/:id", api.GetOnecomment)
			comments.POST("create", api.AddNewcomment)
			comments.DELETE("delete/:id", api.Deletecomment)
			comments.PATCH("update/:id", api.Updatecomment)
		};
		users := v1.Group("/users"); {
			users.POST("register", api.Register)
			users.POST("login", api.Login)
			users.GET("getProfile/:id", api.GetProfile)
			users.DELETE("deleteProfile/:id", api.DeleteUser)
			users.PATCH("updateProfile/:id", api.UpdateProfile)
		}
	}

	return router
}
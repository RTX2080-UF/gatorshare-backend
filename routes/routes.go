package routes

import (
	"gatorshare/controllers"
	"gorm.io/gorm"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func InitializeRoutes(db *gorm.DB, envSrc bool) *gin.Engine {

	router := gin.Default()

	router.Use(gin.Logger())

	router.GET("/", func(responseCtx *gin.Context) {
	  responseCtx.JSON(http.StatusOK, gin.H{"data": "Welcome to Gatorshare made with the help of Go and Gin!"})    
	})

	router.Use(cors.Default())

	api := controllers.Controller{DB: db}

	v1 := router.Group("/v1")
	{	
		users := v1.Group("/users"); {
			users.POST("register", api.Register)
			users.POST("login", api.Login)
			users.GET("refreshToken", api.RefreshToken)
			users.GET("getProfile", api.GetProfile)
			users.GET("getUserProfile/:id", api.GetProfileGeneric)
			users.DELETE("deleteProfile", api.DeleteUser)
			users.PATCH("updateProfile", api.UpdateProfile)
			users.POST("follow/:userId", api.FollowUser)
			users.GET("listFollowers/:userId", api.GetFollowers)
			users.GET("resetPassword", api.ResetPassword)
			users.POST("updatePassword", api.UpdatePassword)
		};
		posts := v1.Group("/posts"); {
			posts.GET("getAll", api.Listpost)
			posts.GET("getOne/:id", api.GetOnepost)
			posts.POST("create", api.AddNewpost)
			posts.PATCH("update/:id", api.UpdatePost)
			posts.DELETE("delete/:id", api.Deletepost)
		};
		comments := v1.Group("/comments"); {
			comments.GET("getAll/:postId", api.GetAllComment)
			comments.GET("getOne/:id", api.GetOneComment)
			comments.POST("create", api.AddNewComment)
			comments.DELETE("delete/:id", api.DeleteComment)
			comments.PATCH("update/:id", api.UpdateComment)
		};
		tags := v1.Group("/tags"); {
			tags.GET("getOne/:tagId", api.GetTag)
			tags.POST("create", api.AddTag)
			tags.DELETE("delete/:id", api.DeleteTag)
			tags.PATCH("update/:id", api.UpdateTag)
			tags.POST("follow/:tagId", api.FollowTagsByUser)
			tags.GET("popularTags/:count",api.PopularTags)
			tags.POST("selectTags",api.SelectTags)
		};
	}

	return router
}
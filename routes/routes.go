package routes

import (
	"gatorshare/controllers"
	"gatorshare/middleware"
	"net/http"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(db *gorm.DB, envSrc bool) *gin.Engine {

	router := gin.Default()

	router.Use(gin.Logger())

	router.GET("/", func(responseCtx *gin.Context) {
	  responseCtx.JSON(http.StatusOK, gin.H{"data": "Welcome to Gatorshare made with the help of Go and Gin!"})    
	})

	// router.Use(cors.Default())
	router.Use(middleware.CORS())

	api := controllers.Controller{DB: db}

	v1 := router.Group("/v1")
	{	
		home := v1.Group("/home"); { 
			home.GET("/user", api.GetUserHome);
			home.GET("/latest", api.GetLatestPost);
		};
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
			users.POST("addFeedback",api.AddFeedback)
			users.GET("getFeedback/:userId",api.GetFeedback)
		};
		posts := v1.Group("/posts"); {
			posts.GET("getAll", api.ListPost)
			posts.GET("getOne/:id", api.GetOnePost)
			posts.GET("getReactions/:postId", api.GetPostReaction)
			posts.POST("create", api.AddNewPost)
			posts.PATCH("update/:id", api.UpdatePost)
			posts.DELETE("delete/:id", api.DeletePost)
			posts.POST("reactToPost", api.ReactToPost)
			posts.POST("searchPost",api.SearchPost)
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
		notifications := v1.Group("/notifications"); {
			notifications.GET("getNew", api.GetNewNotifications)
			notifications.GET("updateStatus", api.UpdateNotifications)
		};
	}

	return router
}
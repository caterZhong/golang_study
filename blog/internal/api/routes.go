package api

import (
	"golang_study/blog/config"
	"golang_study/blog/internal/controller"
	"golang_study/blog/internal/middleware"
	"golang_study/blog/internal/model"
	"golang_study/blog/internal/repo"
	"golang_study/blog/internal/service"
	"golang_study/blog/internal/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	userDao := repo.NewUserRepository(config.DB)
	userService := service.NewUserService(userDao, utils.JWT)
	userController := controller.NewUserController(userService)

	config.DB.AutoMigrate(&model.Post{})
	config.DB.AutoMigrate(&model.Comment{})

	postDao := repo.NewPostRepository(config.DB)
	postService := service.NewPostService(postDao)
	postController := controller.NewPostController(postService)

	commentDao := repo.NewCommentRepository(config.DB)
	commentService := service.NewCommentService(commentDao)
	commentController := controller.NewCommentController(commentService)

	// 在这里注册路由
	pubGroup := r.Group("/pub")
	{
		pubGroup.POST("/user/login", userController.Login)
		pubGroup.GET("/user/getUserInfo", userController.Login)
	}

	{
		pubGroup.POST("/post/getPost", postController.GetPost)
		pubGroup.GET("/post/pagePost", postController.PagePost)
		pubGroup.GET("/comment/getPostComment", commentController.GetPostComment)
	}

	priGroup := r.Group("/pri")
	priGroup.Use(middleware.JWTAuth())
	{
		priGroup.POST("post/createPost", postController.CreatePost)
		priGroup.POST("post/updatePost", postController.UpdatePost)
		priGroup.POST("post/delPost", postController.DeletePost)
		priGroup.POST("comment/createComment", commentController.CreateComment)
	}
}

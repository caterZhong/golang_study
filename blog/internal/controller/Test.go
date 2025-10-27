package controller

// "github.com/gin-gonic/gin"

// type UserController struct{}

// func (uc *UserController) GetUser(c *gin.Context) {
// 	c.JSON(200, gin.H{"method": "GET", "path": c.Request.URL.Path})
// }

// func (uc *UserController) CreateUser(c *gin.Context) {
// 	c.JSON(201, gin.H{"method": "POST", "path": c.Request.URL.Path})
// }

// func main() {
// 	r := gin.Default()

// 	// 直接绑定结构体方法（可能每次请求创建新实例）
// 	api := r.Group("/api/v1")
// 	{
// 		api.GET("/users", UserController.GetUser)
// 		api.POST("/users", UserController.CreateUser)
// 	}

// 	r.Run(":8080")
// }

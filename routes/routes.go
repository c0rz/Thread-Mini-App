package routes

import (
	"Blog/controllers"
	"Blog/middleware"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	routes := gin.New()

	api := routes.Group("/api")
	// Auth
	api.POST("auth/register", controllers.Register)
	api.POST("auth/login", controllers.Login)
	//  Blog
	api.GET("/thread", controllers.GetAllThread)
	api.GET("/thread/:id", controllers.GetThread) // Beluman

	blog := routes.Group("/api")
	blog.Use(middleware.JwtAuthMiddlewareUser())
	blog.POST("thread/send", controllers.SubmitThread)
	blog.DELETE("thread/:id", controllers.DeleteThread) // Note
	blog.PUT("thread/:id", controllers.UpdateThread)

	// Category
	admin := routes.Group("/admin")
	admin.Use(middleware.JwtAuthMiddlewareAdmin())
	admin.POST("category/add", controllers.CreateCategory)
	admin.DELETE("category/:id", controllers.DeleteCategory)

	// Comment
	blog.POST("/thread/:id", controllers.AddComment)

	return routes
}

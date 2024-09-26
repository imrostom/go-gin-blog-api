package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imrostom/go-blog-api/controllers"
	"github.com/imrostom/go-blog-api/middleware"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/", controllers.DefaultHandler)
	router.GET("/test", controllers.TestHandler)

	public := router.Group("/api")
	{
		// Home Page Route
		public.GET("/", controllers.DefaultHandler)

		// Auth Route
		public.POST("/login", controllers.LoginHandler)
		public.POST("/register", controllers.RegisterHandler)

		//Category Route
		public.GET("/categories", controllers.GetCategoryHandler)
		public.GET("/categories/:id", controllers.ShowCategoryHandler)

		//User Route
		public.GET("/users", controllers.GetUserHandler)
		public.GET("/users/:id", controllers.ShowUserHandler)


		//Post Route
		public.GET("/posts", controllers.GetPostHandler)
		public.GET("/posts/:id", controllers.ShowPostHandler)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		//Category Route
		protected.POST("/categories", controllers.CreateCategoryHandler)
		protected.PUT("/categories/:id", controllers.UpdateCategoryHandler)
		protected.DELETE("/categories/:id", controllers.DeleteCategoryHandler)

		//User Route
		protected.POST("/users", controllers.CreateUserHandler)
		protected.PUT("/users/:id", controllers.UpdateUserHandler)
		protected.DELETE("/users/:id", controllers.DeleteUserHandler)


		//Post Route
		protected.POST("/posts", controllers.CreatePostHandler)
		protected.PUT("/posts/:id", controllers.UpdatePostHandler)
		protected.DELETE("/posts/:id", controllers.DeletePostHandler)
	}

}

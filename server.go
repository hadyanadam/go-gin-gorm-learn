package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hadyanadam/go-gorm-gin/config"
	"github.com/hadyanadam/go-gorm-gin/controller"
	"github.com/hadyanadam/go-gorm-gin/middleware"
	"github.com/hadyanadam/go-gorm-gin/repository"
	"github.com/hadyanadam/go-gorm-gin/service"
	"gorm.io/gorm"
)



var (
	db 				*gorm.DB 					= config.SetupDatabaseConnection()
	userRepository 	repository.UserRepository 	= repository.NewUserRepository(db)
	bookRepository	repository.BookRepository	= repository.NewBookRepository(db)
	jwtService 		service.JWTService 			= service.NewJWTService()
	userService		service.UserService			= service.NewUserService(userRepository)
	authService 	service.AuthService 		= service.NewAuthService(userRepository)
	bookService		service.BookService			= service.NewBookService(bookRepository)
	authController 	controller.AuthController 	= controller.NewAuthController(authService, jwtService)
	userController	controller.UserController	= controller.NewUserController(userService, jwtService)
	bookController	controller.BookController	= controller.NewBookController(bookService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	bookRoutes := r.Group("api/book", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	r.Run(":" + port)
}
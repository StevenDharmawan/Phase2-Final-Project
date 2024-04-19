package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"phase2-final-project/config"
	"phase2-final-project/controller"
	"phase2-final-project/middlewares"
	"phase2-final-project/repository"
	"phase2-final-project/service"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Failed open env ", err)
		return
	}
}

// @title Final Project Phase 2 - Hotel
// @version 1.0
// @description API for hotel.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host https://phase2-final-project-d73478d4951c.herokuapp.com/:8080
// @BasePath /api/v1
func main() {
	db, err := config.ConnectDB()
	if err != nil {
		return
	}
	topUpHistoryRepository := repository.NewTopUpHistoryRepository(db)
	topUpHistoryService := service.NewTopUpHistoryService(topUpHistoryRepository)
	topUpHistoryController := controller.NewTopUpHistoryController(topUpHistoryService)

	walletRepository := repository.NewWalletRepository(db)
	walletService := service.NewWalletService(walletRepository, topUpHistoryRepository)
	walletController := controller.NewWalletController(walletService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, walletRepository)
	userController := controller.NewUserController(userService)

	roomRepository := repository.NewRoomRepository(db)
	roomService := service.NewRoomService(roomRepository)
	roomController := controller.NewRoomController(roomService)

	bookingRepository := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepository, roomRepository, walletRepository, userRepository)
	bookingController := controller.NewBookingController(bookingService)
	url := ginSwagger.URL("https://phase2-final-project-d73478d4951c.herokuapp.com/swagger/doc.json") // The url pointing to API definition
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	v1 := r.Group("/api/v1")
	{
		v1.POST("/users/register/admin", userController.RegisterAdmin)
		v1.POST("/users/register", userController.RegisterUser)
		v1.POST("/users/login", userController.LoginUser)
		v1.GET("/users", middlewares.CustomerAuthMiddleware(), userController.GetUser)
		v1.PUT("/users", middlewares.CustomerAuthMiddleware(), userController.UpdateUser)
		v1.GET("/users/topuphistories", middlewares.CustomerAuthMiddleware(), topUpHistoryController.GetAllHistoryById)

		v1.POST("/wallets", middlewares.CustomerAuthMiddleware(), walletController.TopUpWallet)

		v1.GET("/topuphistories", middlewares.AdminAuthMiddleware(), topUpHistoryController.GetAllHistory)

		v1.GET("/rooms", middlewares.AdminAuthMiddleware(), roomController.GetAllRooms)
		v1.POST("/rooms", middlewares.AdminAuthMiddleware(), roomController.CreateRoom)
		v1.GET("/rooms/availability", roomController.GetAllAvailableRooms)

		v1.POST("/bookings", middlewares.CustomerAuthMiddleware(), bookingController.CreateBooking)
	}
	r.Run(":" + config.GetEnv("PORT"))
}

package app

import (
	reminderController "backend/controllers/reminder"
	userController "backend/controllers/user"
	middleware "backend/middleware"
	service "backend/services/implementations"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func init() {
	engine = gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Especifica tu origen aquí
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func StartApp() {
	// Inicializar servicios
	userServiceImplementation := service.UserServiceImplementation{}
	reminderServiceImplementation := service.ReminderServiceImplementation{}

	// Inicializar controladores
	userControllerImplementation := userController.UserControllerImplementation{
		UserService: userServiceImplementation,
	}
	reminderControllerImplementation := reminderController.ReminderControllerImplementation{
		ReminderService: reminderServiceImplementation,
	}

	// Rutas de usuario
	engine.POST("/api/Login", userControllerImplementation.Login)
	engine.POST("/api/Register", userControllerImplementation.RegisterUser)

	// Rutas protegidas con middleware de autenticación
	authorized := engine.Group("")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/api/Reminders/Add", reminderControllerImplementation.AddReminder)
		authorized.PUT("/api/Reminders/Update", reminderControllerImplementation.UpdateReminder)
		authorized.GET("/api/Reminders", reminderControllerImplementation.GetRemindersByUserId)
		authorized.DELETE("/api/Reminders/Delete/:reminderId", reminderControllerImplementation.DeleteReminder)
	}

	// Iniciar el servidor
	engine.Run(":8090")
}

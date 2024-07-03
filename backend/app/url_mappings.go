package app

import (
	reminderController "backend/controllers/reminder"
	userController "backend/controllers/user"
	middleware "backend/middleware"
	service "backend/services/implementations"

	"github.com/gin-gonic/gin"
)

func mapUrls(engine *gin.Engine) {
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
}

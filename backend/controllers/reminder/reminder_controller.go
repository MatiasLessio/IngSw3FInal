package reminders

import (
	dto "backend/dtos/reminders"
	"backend/middleware"
	services "backend/services/contracts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReminderControllerContract interface {
	AddReminder(c *gin.Context)
	GetRemindersByUserId(c *gin.Context)
	DeleteReminder(c *gin.Context)
	UpdateReminder(c *gin.Context)
}

type ReminderControllerImplementation struct {
	ReminderService services.ReminderServiceContract
}

func (reminderController ReminderControllerImplementation) AddReminder(context *gin.Context) {
	userId, er := middleware.GetUserIdByToken(context)
	if er != nil {
		context.JSON(er.Status(), er)
		return
	}
	var request dto.ReminderDto
	request.UserId = userId
	context.BindJSON(&request)
	response, err := reminderController.ReminderService.AddReminder(request)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusCreated, response)
}

func (reminderController ReminderControllerImplementation) GetRemindersByUserId(c *gin.Context) {
	userId, er := middleware.GetUserIdByToken(c)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	response, err := reminderController.ReminderService.GetRemindersByUserId(userId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, response)
}
func (reminderController ReminderControllerImplementation) DeleteReminder(c *gin.Context) {
	reminderId := c.Param("reminderId")
	id, er := strconv.Atoi(reminderId)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reminder ID"})
		return
	}
	err := reminderController.ReminderService.DeleteReminder(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, "Reminder deleted successfully")
}
func (reminderController ReminderControllerImplementation) UpdateReminder(c *gin.Context) {
	userId, er := middleware.GetUserIdByToken(c)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	var reminderRequest dto.ReminderDto
	reminderRequest.UserId = userId
	c.BindJSON(&reminderRequest)
	response, err := reminderController.ReminderService.UpdateReminder(reminderRequest)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, response)
}

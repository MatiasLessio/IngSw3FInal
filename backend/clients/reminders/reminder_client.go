package reminders

import (
	"backend/models"
	e "backend/utilities"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func AddReminder(reminder models.Reminder) (models.Reminder, e.ApiError) {
	result := Db.Create(&reminder)
	if result.Error != nil {
		return reminder, e.NewBadRequestApiError("error while adding a reminder")
	}
	return reminder, nil
}

func GetRemindersByUserId(userId int) (models.Reminders, e.ApiError) {
	var reminders models.Reminders
	result := Db.Where("user_id = ?", userId).Find(&reminders)
	if result.Error != nil {
		return reminders, e.NewBadRequestApiError("error while listing reminders")
	}
	return reminders, nil
}

func DeleteReminder(reminderId int) e.ApiError {
	var reminder models.Reminder
	reminder.ReminderId = reminderId
	result := Db.Where("reminder_id = ?", reminder.ReminderId).Delete(&reminder)
	if result.Error != nil {
		return e.NewBadRequestApiError("error while deleting reminder")
	}
	if result.RowsAffected == 0 {
		return e.NewNotFoundApiError("reminder not found")
	}
	return nil
}

func UpdateReminder(reminder models.Reminder) (models.Reminder, e.ApiError) {
	result := Db.Model(&reminder).Where("reminder_id = ?", reminder.ReminderId).Updates(models.Reminder{Title: reminder.Title, Description: reminder.Description})
	if result.Error != nil {

		return reminder, e.NewBadRequestApiError("error while updating reminder")
	}
	return reminder, nil
}

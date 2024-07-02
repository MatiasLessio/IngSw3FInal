package services

import (
	client "backend/clients/reminders"
	dto "backend/dtos/reminders"
	"backend/models"
	e "backend/utilities"
)

type ReminderServiceImplementation struct{}

func NewReminderServiceImpl() ReminderServiceImplementation {
	return ReminderServiceImplementation{}
}

func (ReminderServiceImplementation) AddReminder(reminder dto.ReminderDto) (dto.ReminderDto, e.ApiError) {
	var newReminder models.Reminder

	newReminder.Title = reminder.Title
	newReminder.Description = reminder.Description
	newReminder.UserId = reminder.UserId

	newReminder, er := client.AddReminder(newReminder)
	if er != nil {
		return reminder, e.NewBadRequestApiError("error while adding reminder")
	}
	return reminder, nil
}

func (ReminderServiceImplementation) GetRemindersByUserId(userId int) (dto.RemindersDto, e.ApiError) {
	var reminders models.Reminders
	reminders, err := client.GetRemindersByUserId(userId)
	if err != nil {
		return nil, e.NewBadRequestApiError("error while searching reminders")
	}
	var remindersDto dto.RemindersDto
	for _, reminder := range reminders {
		remindersDto = append(remindersDto, dto.ReminderDto{
			ReminderId:  reminder.ReminderId,
			Description: reminder.Description,
			UserId:      reminder.UserId,
			Title:       reminder.Title,
		})
	}
	return remindersDto, nil
}

func (ReminderServiceImplementation) DeleteReminder(reminderId int) e.ApiError {
	err := client.DeleteReminder(reminderId)
	if err != nil {
		return e.NewBadRequestApiError("error while deleting reminder")
	}
	return nil
}
func (ReminderServiceImplementation) UpdateReminder(request dto.ReminderDto) (dto.ReminderDto, e.ApiError) {
	var reminder models.Reminder

	reminder.Description = request.Description
	reminder.ReminderId = request.ReminderId
	reminder.Title = request.Title
	reminder.UserId = request.UserId

	reminder, err := client.UpdateReminder(reminder)
	if err != nil {
		return request, e.NewBadRequestApiError("error while updating reminder")
	}
	return request, nil
}

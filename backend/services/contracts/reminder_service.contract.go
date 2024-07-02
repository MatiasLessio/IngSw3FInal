package services

import (
	dto "backend/dtos/reminders"
	e "backend/utilities"
)

type ReminderServiceContract interface {
	AddReminder(reminder dto.ReminderDto) (dto.ReminderDto, e.ApiError)
	GetRemindersByUserId(userId int) (dto.RemindersDto, e.ApiError)
	DeleteReminder(reminderiD int) e.ApiError
	UpdateReminder(reminder dto.ReminderDto) (dto.ReminderDto, e.ApiError)
}

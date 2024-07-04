package test

import (
	dto "backend/dtos/reminders"
	e "backend/utilities"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

type reminderService struct {
	reminderClient *MockClient
}

func (m *MockClient) GetRemindersByUserId(userId int) dto.RemindersDto {
	if userId == 7 {
		return dto.RemindersDto{
			{
				ReminderId:  44,
				Title:       "new test",
				Description: "123",
				UserId:      7,
			},
			{
				ReminderId:  45,
				Title:       "second",
				Description: "second test",
				UserId:      7,
			},
		}
	} else {
		return dto.RemindersDto{}
	}
}

func (m *MockClient) AddReminder(request dto.ReminderDto) (dto.ReminderDto, e.ApiError) {
	return request, nil
}

func (m *MockClient) UpdateReminder(request dto.ReminderDto) (dto.ReminderDto, e.ApiError) {
	return request, nil
}

func (m *MockClient) DeleteReminder(reminderId int) e.ApiError {
	return nil
}

func TestGetRemindersByUserId(t *testing.T) {
	service := &reminderService{reminderClient: NewMockClient()}

	var reminders dto.RemindersDto = service.reminderClient.GetRemindersByUserId(7)
	expectedReminders := dto.RemindersDto{
		{
			ReminderId:  44,
			Title:       "new test",
			Description: "123",
			UserId:      7,
		},
		{
			ReminderId:  45,
			Title:       "second",
			Description: "second test",
			UserId:      7,
		},
	}
	assert.Equal(t, expectedReminders, reminders, "reminders not equal")
}

func TestAddReminder(t *testing.T) {
	service := &reminderService{reminderClient: NewMockClient()}

	reminder := dto.ReminderDto{
		Title:       "title test",
		Description: "description test",
		UserId:      1,
	}

	_, err := service.reminderClient.AddReminder(reminder)
	assert.NoError(t, err, "error while adding a reminder")
}

func TestUpdateReminder(t *testing.T) {
	service := &reminderService{reminderClient: NewMockClient()}

	reminder := dto.ReminderDto{
		Title:       "title test",
		Description: "description test",
		UserId:      1,
		ReminderId:  45,
	}

	_, err := service.reminderClient.UpdateReminder(reminder)
	assert.NoError(t, err, "error while updating a reminder")
}
func TestDeleteReminder(t *testing.T) {
	service := &reminderService{reminderClient: NewMockClient()}

	reminderId := 40

	response := service.reminderClient.DeleteReminder(reminderId)
	assert.NoError(t, response, "error while deleting a reminder")
}

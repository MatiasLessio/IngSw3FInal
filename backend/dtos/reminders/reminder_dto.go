package dto

type ReminderDto struct {
	ReminderId  int    `json:"reminderId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`
}

type RemindersDto []ReminderDto

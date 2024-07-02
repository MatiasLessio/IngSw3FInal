package models

type Reminder struct {
	ReminderId  int    `gorm:"primaryKey;"`
	Title       string `gorm:"type:varchar(256);not null"`
	UserId      int    `gorm:"not null"`
	Description string `gorm:"type:varchar(256);not null"`
}

type Reminders []Reminder

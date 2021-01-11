package models

// Message is a model data for message
type Message struct {
	ID          int    `json:"-" gorm:"column:id"`
	Name        string `json:"name" gorm:"name"`
	Email       string `json:"email" gorm:"email"`
	PhoneNumber string `json:"phoneNumber" gorm:"phone_number"`
	Subject     string `json:"subject" gorm:"subject"`
	Message     string `json:"message" gorm:"message"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type MUser struct {
	ID           int64          `gorm:"primaryKey"`
	Username     string         `gorm:"type:varchar(100)"`
	Password     string         `gorm:"type:varchar(255)"`
	Name         string         `gorm:"type:varchar(255)"`
	Email        string         `gorm:"type:varchar(255)"`
	PhoneNumber  string         `gorm:"type:varchar(15)"`
	SessionLogin string         `gorm:"type:varchar(255)"`
	Gender       string         `gorm:"type:varchar(50)"`
	Photo        string         `gorm:"type:varchar(255)"`
	BirthPlace   string         `gorm:"type:varchar(255)"`
	EmployeeID   string         `gorm:"type:varchar(255)"`
	EmployeeType string         `gorm:"type:varchar(255)"`
	BirthDate    string         `gorm:"type:varchar(255)"`
	CreatedAt    time.Time      `gorm:"type:timestamp with time zone"`
	CreatedBy    string         `gorm:"type:varchar(100)"`
	UpdatedBy    string         `gorm:"type:varchar(100)"`
	UpdatedAt    time.Time      `gorm:"type:timestamp with time zone"`
	DeletedBy    string         `gorm:"type:varchar(100)"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamp with time zone"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MyProfile struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	SessionLogin string `json:"session_login"`
	Gender       string `json:"gender"`
	Photo        string `json:"photo"`
	BirthPlace   string `json:"birth_place"`
	EmployeeID   string `json:"employee_id"`
	EmployeeType string `json:"employee_type"`
	BirthDate    string `json:"birth_date"`
}

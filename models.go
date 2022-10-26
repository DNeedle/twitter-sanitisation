package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-" sql:"index"`
	CompanyName  string         `json:"companyName"`
	FirstName    string         `json:"firstName"`
	LastName     string         `json:"lastName"`
	Email        string         `json:"email"`
	Password     string         `json:"-"`
	Country      string         `json:"country"`
	Phone        string         `json:"phone"`
	DateOfBirth  time.Time      `json:"dateOfBirth"`
	Rank         string         `json:"rank"`
	Gender       string         `json:"gender"`
	Plan         string         `json:"plan"`
	ProfileImage string         `json:"profileImage"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id uint32 `gorm:"primary_key;auto_increment" json:"id"`
	//Required: true
	Username string `gorm:"size:255;not null;unique" json:"username"`
	//Required: true
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (user *User) Create(db *gorm.DB) (*User, error) {
	var err error
	err = db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, err
}

func (user *User) Read(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Find(&user).Error
	if err != nil {
		return &[]User{}, err
	}

	return &users, err
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Mail     string `gorm:"type:varchar(40);unique" json:"mail,omitempty"`
	Password string `gorm:"size:255" json:"password,omitempty"`
}

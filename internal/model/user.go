package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       int64  `gorm:"primarykey"`
	Name     string `gorm:"type:varchar;not null"`
	Email    string `gorm:"type:varchar;not null;index:idx_email,unique"`
	Password string `gorm:"type:varchar; not null"`

	Followers  []User     `gorm:"foreignKey:ID"`
	Photos     []Photo    `gorm:"foreignKey:UserID"`
	Activities []Activity `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
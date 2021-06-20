package model

import "time"

type UserAccount struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
	UID       string `gorm:"unique"`
}

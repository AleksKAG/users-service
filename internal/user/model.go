package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint32 `gorm:"primaryKey"`
	Email string
}

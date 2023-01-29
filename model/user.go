package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"` //int64  `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

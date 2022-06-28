package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model `json:"-"`
	ID         uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name       string    `gorm:"size:60;not null; unique" json:"name,omitempty"`
	Email      string    `gorm:"size:60;not null; unique" json:"email,omitempty"`
	Password   string    `gorm:"size:100;not null" json:"password,omitempty"`
	Admin      bool      `gorm:"default:false" json:"admin,omitempty"`
	Enable     bool      `gorm:"default:true" json:"enable,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type UserResume struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"nick,omitempty"`
}

func (u *Users) ValidateOnCreate() (err error) {
	if u.Name == "" {
		return errors.New("Name cant be empty")
	}

	if u.Email == "" {
		return errors.New("Email cant be empty")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("E-mail invalid")
	}

	if u.Password == "" {
		return errors.New("Password cant be empty")
	}
	return

}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	return
}

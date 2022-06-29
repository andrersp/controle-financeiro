package models

import (
	"errors"
	"strings"
	"time"

	"github.com/andrersp/controle-financeiro/src/core"
	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model `json:"-"`
	ID         uint           `gorm:"primaryKey" json:"id,omitempty"`
	Name       string         `gorm:"size:60;not null; unique" json:"name,omitempty"`
	Email      string         `gorm:"size:60;not null; unique" json:"email,omitempty"`
	Password   string         `gorm:"size:100;not null" json:"password,omitempty"`
	Admin      bool           `gorm:"default:false" json:"admin,omitempty"`
	Enable     bool           `gorm:"default:true" json:"enable,omitempty"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type UserResume struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type UserPassword struct {
	New string `json:"new"`
	Old string `json:"old"`
}

func (u *Users) validateFields() (err error) {
	if u.Name == "" {
		return errors.New("Name cant be empty")
	}

	if u.Email == "" {
		return errors.New("Email cant be empty")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("E-mail invalid")
	}
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)

	return

}

func (u *Users) searchDuplicates(tx *gorm.DB) (err error) {
	result := UserResume{}

	err = tx.Unscoped().Model(&Users{}).Select("id", "name", "email").Where("name = ? OR email = ?", u.Name, u.Email).Find(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if result.Email == u.Email && result.ID != u.ID {
		err = errors.New("Duplicate Email")
	}

	if result.Name == u.Name && result.ID != u.ID {
		err = errors.New("Duplicate Name")
	}
	return err

}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {

	if err = u.validateFields(); err != nil {
		return err
	}
	if u.Password == "" {
		return errors.New("Password cant be empty")
	}

	if err = u.searchDuplicates(tx); err != nil {
		return err
	}

	hashedPassword, err := core.HashGenerator(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return
}

func (u *Users) BeforeUpdate(tx *gorm.DB) (err error) {

	if err := u.validateFields(); err != nil {
		return err
	}

	if err := u.searchDuplicates(tx); err != nil {
		return err
	}

	return

}

package crud

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andrersp/controle-financeiro/src/database"
	"github.com/andrersp/controle-financeiro/src/models"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewCrudUser(db *gorm.DB) *User {
	return &User{db}
}

// Create New User
func (u User) Create(user models.Users) (database.SinglePage, error) {

	var result database.SinglePage

	var userResume models.UserResume

	if err := u.db.Scopes(database.SingleResult(&models.Users{}, &result, u.db)).
		Model(&models.Users{}).Create(&user).Find(&userResume).Error; err != nil {
		return database.SinglePage{}, err
	}
	result.Data = userResume

	return result, nil
}

// SearchUsers from name or email
func (u User) SearchUsers(nameOrEmail string, req *http.Request) (database.Pagination, error) {

	results := database.Pagination{}

	var users []models.UserResume
	nameOrEmail = fmt.Sprintf("%%%s%%", nameOrEmail)

	query := u.db.
		Scopes(database.Paginator([]models.Users{}, &results, u.db, req)).
		Model(&models.Users{}).
		Select("id", "name", "email").Where("name LIKE ? OR email LIKE ?", nameOrEmail, nameOrEmail).Find(&users)

	if err := query.Error; err != nil {
		return results, err
	}
	results.Data = users
	return results, nil

}

func (u User) SelectUser(userID uint) (database.SinglePage, error) {
	var user models.Users
	var result database.SinglePage

	query := u.db.Scopes(database.SingleResult(&user, &result, u.db)).
		Model(&models.Users{}).
		Select("id", "name", "email", "enable", "admin", "created_at").
		First(&user, userID)
	if err := query.Error; err != nil {
		return database.SinglePage{}, err
	}
	result.Data = user

	return result, nil
}

func (u User) UpdateUser(userID uint, user models.Users) (err error) {

	var result models.Users
	if err := u.db.First(&models.Users{}, userID).Find(&result).Error; err != nil {

		return err
	}

	err = u.db.Model(&result).
		Omit("password", "created_at").
		Updates(user).Error

	if err != nil {
		return err
	}

	return err
}

func (u User) DeleteUser(userID uint) (err error) {

	if err = u.db.
		First(&models.Users{}, userID).Error; err != nil {
		return
	}

	err = u.db.Delete(&models.Users{}, userID).Error
	return
}

func (u User) GetUserPassword(userID uint) (userPassword string, err error) {

	var result models.Users

	err = u.db.First(&result, userID).Error
	userPassword = result.Password

	return
}

func (u User) UpdateUserPassword(userID uint, password string) (err error) {

	err = u.db.Session(&gorm.Session{SkipHooks: true}).Model(&models.Users{}).
		Where("id = ?", userID).
		Update("password", password).Error
	return
}

func (u User) SearchByEmail(email string) (user models.Users, err error) {
	query := u.db.Where("email = ?", email).First(&user)

	if query.Error != nil {
		err = errors.New("Email or Passowrd Invalid!")

	}
	return

}

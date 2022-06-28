package crud

import (
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
func (u User) Create(user models.Users) (models.UserResume, error) {

	var userResume models.UserResume

	if err := u.db.Model(&models.Users{}).Create(&user).Find(&userResume).Error; err != nil {
		return userResume, err
	}

	return userResume, nil
}

// SearchUsers from name or email
func (u User) SearchUsers(nameOrEmail string, req *http.Request) (database.Pagination, error) {

	results := database.Pagination{}

	var users []models.Users
	nameOrEmail = fmt.Sprintf("%%%s%%", nameOrEmail)

	query := u.db.Scopes(database.Paginator([]models.Users{}, &results, u.db, req)).
		Select("id", "name").Where("name LIKE ? OR email LIKE ?", nameOrEmail, nameOrEmail).Find(&users)

	if err := query.Error; err != nil {
		return results, err
	}
	return results, nil

}

package auth

import (
	"github.com/izhujiang/appengine/module/core"
	"github.com/izhujiang/appengine/module/db"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(user User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := db.Instance.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id core.UUID) (User, error) {
	var user User
	user.ID = id
	err := db.Instance.Debug().Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

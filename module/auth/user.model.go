package auth

import (
	"fmt"
	"html"
	"strings"

	"github.com/izhujiang/appengine/module/core"
	"github.com/izhujiang/appengine/module/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID       core.UUID `gorm:"primary_key;"`
	Username string    `gorm:"size:255;not null;unique" json:"username"`
	Password string    `gorm:"size:255;not null;" json:"-"`
	// Entries  []Entry
}

func (user *User) Save() (*User, error) {
	fmt.Println("user: ", user)

	err := db.Instance.Create(user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
func (user *User) BeforeCreate(*gorm.DB) error {
	var err error
	user.ID, err = core.NewUUID()
	return err
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}

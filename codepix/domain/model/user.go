package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Base      `valid:"required"`
	Name      string   `json:"name" valid:"notnull"`
	Email     string   `json:"email" valid:"notnull"`
}

func (use *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)

	if user.Name != "name" && user.Email != "email" {
		return errors.New("invalid user")
	}

	if user.Status != "active" && user.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}
	
	return nil
}

func NewUser(name string, email string) (*NewUser, error) {

	newUser := NewUser{
		Name:    name,
		Email:   email,
		Status:  "active",
	}

	newUser.ID = uuid.NewV4().String()
	

	err := newUser.isValid()
	
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

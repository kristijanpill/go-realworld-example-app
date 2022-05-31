package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID uuid.UUID `gorm:"primaryKey; unique; type:uuid"`
	Email    string `gorm:"unique"`
	Password string
	Active   bool
}

func NewUser(email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID: uuid.New(),
		Email: email,
		Password: string(hashedPassword),
		Active: true,
	}, nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
package model

import (
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (u *User) ComparePassword(password string) error {
	hashedPassword, err := base64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return fmt.Errorf("failed to decode hashed password: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}

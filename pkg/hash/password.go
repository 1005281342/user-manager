package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Password(s string) (string, error) {
	if len(s) < 6 {
		return "", fmt.Errorf("密码过短")
	}

	var password, err = bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(password), nil
}

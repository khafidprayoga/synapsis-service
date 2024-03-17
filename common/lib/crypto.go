package commonLib

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPasssword(pass string) (hashedPass string, err error) {
	out, errHash := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if errHash != nil {
		return "", fmt.Errorf("error hashing password: %v", errHash)
	}
	return string(out), nil
}

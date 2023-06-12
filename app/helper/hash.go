package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10);
	if err != nil {
		return nil, err
	}
	return hashed, nil
}

func CompareHashedPassword(hashed string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	}
	return true
}
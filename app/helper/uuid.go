package helper

import "github.com/google/uuid"

func GenerateNewId() (newId string) {
	id := uuid.New()
	return id.String()
}
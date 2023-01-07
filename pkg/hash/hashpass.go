package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePass(hashedPass string, plainPass []byte) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), plainPass); err != nil {
		log.Printf("Error while comparing passwords: %v\n", err)
		return false
	}

	return true
}

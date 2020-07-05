package crypto_utils

import (
	"fmt"
	"github.com/lelinu/api_utils/utils/random_utils"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromString(input string) (string, error) {

	hashedInput, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedInput), nil
}

func CompareHashWithClearPassword(hash string, clearPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(clearPassword))
	if err != nil {
		return err
	}

	return nil
}

func GenerateHashAndSaltKeyFromInput(input string) (string, string, error) {

	newSalt, err := random_utils.NewUUID()
	if err != nil {
		return "", "", err
	}
	hashedInput, err := GenerateHashFromString(fmt.Sprintf("%s:%s", input, newSalt))
	if err != nil {
		return "", "", err
	}

	return hashedInput, newSalt, nil
}

func CompareHashWithClearPasswordAndSalt(hash string, clearPassword string, salt string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(fmt.Sprintf("%s:%s", clearPassword, salt)))
	if err != nil {
		return err
	}

	return nil
}

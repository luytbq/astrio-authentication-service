package auth

import (
	"errors"
	"fmt"
	"log"
	"net/mail"
	"strings"

	"github.com/luytbq/astrio-authentication-service/pkg/auth"
)

const (
	KEY_AUTH_TOKEN = "Astrio-Auth-Token"
)

func validateRegisterPayload(payload auth.RegisterPayload) error {
	_, err := mail.ParseAddress(payload.Email)
	if err != nil {
		return errors.New("invalid email")
	}

	if len(payload.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if strings.Compare(payload.Password, payload.RepeatPassword) != 0 {
		return errors.New("repeat password doesn't match")
	}

	return nil
}

func nomarlizeEmail(email string) string {
	arr := strings.Split(email, "@")
	if len(arr) != 2 {
		return email
	}

	local := strings.ToLower(arr[0])
	domain := strings.ToLower(arr[1])

	replaceDots := strings.ReplaceAll(local, ".", "")

	result := fmt.Sprintf("%s@%s", replaceDots, domain)
	if strings.Compare(result, email) != 0 {
		log.Printf("email normalized from %s to %s", email, result)
	}

	return result
}

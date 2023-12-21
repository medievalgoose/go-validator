package util

import "net/mail"

func ValidateEmail(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}

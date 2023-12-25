package util

import (
	"fmt"
)

func Rot13Encode(plaintext string) string {
	encodedMessage := ""

	for _, char := range plaintext {
		if byte(char) >= 'a' && byte(char) <= 'z' {
			actualNum := byte(char) + 13
			if byte(actualNum) > 'z' {
				actualNum = actualNum - 'z' + 'a' - 1
			}
			encodedMessage += fmt.Sprintf("%c", actualNum)
		}
	}

	return encodedMessage
}

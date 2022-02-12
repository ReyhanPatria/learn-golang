package module

import (
	"errors"
	"fmt"
	"strings"
)

func GenerateHelloMessage(name string) (string, error) {
	if strings.ReplaceAll(name, " ", "") == "" {
		return "", errors.New("blank name")
	}
	return fmt.Sprintf("Hello %s!", name), nil
}

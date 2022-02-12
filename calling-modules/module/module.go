package module

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateHelloMessage(name string) (string, error) {
	if strings.ReplaceAll(name, " ", "") == "" {
		return "", errors.New("blank name")
	}

	format := randomHelloFormat()

	return fmt.Sprintf(format, name), nil
}

func randomHelloFormat() string {
	formats := []string{
		"Hello %s!",
		"Top of the morning to you, %s!",
		"How's it going %s?",
		"Great to see you, %s!",
	}

	index := rand.Intn(len(formats))

	return formats[index]
}

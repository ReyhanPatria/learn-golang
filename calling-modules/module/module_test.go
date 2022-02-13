package module

import (
	"regexp"
	"testing"
)

func TestGenerateHelloMessage_ShouldSucceed_WhenNameIsNotBlank(t *testing.T) {
	// given
	name := "Reyhan"

	// when
	message, err := GenerateHelloMessage(name)

	// then
	expected := regexp.MustCompile(`\b` + name + `\b`)

	if !expected.MatchString(message) || err != nil {
		t.Fatalf("GenerateHelloMessage(%s) failed with error '%s'", name, err)
	}
}

func TestGenerateHelloMessage_ShouldFail_WhenNameIsBlank(t *testing.T) {
	// given
	name := "      "

	// when
	message, err := GenerateHelloMessage(name)

	// then
	expected := regexp.MustCompile(`\b` + name + `\b`)

	if expected.MatchString(message) || err == nil {
		t.Fatalf("GenerateHelloMessage(%s) failed with error '%s'", name, err)
	}
}

func TestGenerateMultipleHelloMessage_ShouldSucced_WhenAllNamesIsNotBlank(t *testing.T) {
	// given
	names := []string{
		"Steve",
		"Johnson",
		"Kevin",
	}

	// when
	messages, err := GenerateMultipleHelloMessage(names)

	// then
	for name, message := range messages {
		expected := regexp.MustCompile(`\b` + name + `\b`)

		if !expected.MatchString(message) || err != nil {
			t.Fatalf("GenerateMultipleHelloMessage(%s) failed with error '%s'", names, err)
		}
	}
}

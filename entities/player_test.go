package entities

import (
	"fmt"
	"testing"
)

func TestValidatePlayer(t *testing.T) {
	player := Player{
		Email: "test@example.com",
		Name:  "Test",
	}

	err := player.Validate()

	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}
}

func TestValidatePlayerRequiresEmail(t *testing.T) {
	player := Player{
		Name: "Test",
	}

	err := player.Validate()

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	expected := "email is required"
	msg := fmt.Sprintf("%s", err)
	if expected != msg {
		t.Errorf("Expected error '%s', but got '%s'", expected, msg)
	}
}

func TestValidatePlayerRequiresEmailChecksForWhitespace(t *testing.T) {
	player := Player{
		Email: " \t\r\n",
		Name:  "Test",
	}

	err := player.Validate()

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	expected := "email is required"
	msg := fmt.Sprintf("%s", err)
	if expected != msg {
		t.Errorf("Expected error '%s', but got '%s'", expected, msg)
	}
}

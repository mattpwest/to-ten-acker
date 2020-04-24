package entities

import (
	"errors"
	"strings"
)

type Player struct {
	Email string
	Name  string
}

func (player *Player) Validate() error {
	if strings.TrimSpace(player.Email) == "" {
		return errors.New("email is required")
	}

	return nil
}

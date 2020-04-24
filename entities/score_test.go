package entities

import (
	"testing"
	"time"
)

func TestAddScore(t *testing.T) {
	player := Player{
		Name: "Bob",
	}
	score := Score{
		Player:    player,
		Score:     100,
		Timestamp: time.Now(),
	}
	list := ScoreList{}

	list.AddScore(score)

	if len(list.Scores) != 1 {
		t.Errorf("Expected 1 score in list but got %d", len(list.Scores))
	}
}

func TestAddScoreWithInvalidPlayerFails(t *testing.T) {
	player := Player{
		Name: "Bob",
	}
	score := Score{
		Player:    player,
		Score:     100,
		Timestamp: time.Now(),
	}
	list := ScoreList{}

	list.AddScore(score)

	if len(list.Scores) != 1 {
		t.Errorf("Expected 1 score in list but got %d", len(list.Scores))
	}
}

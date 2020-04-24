package entities

import "time"

type Score struct {
	Player    Player
	Score     int
	Timestamp time.Time
}

type ScoreList struct {
	Scores []Score
}

func (list *ScoreList) AddScore(score Score) {
	list.Scores = append(list.Scores, score)
}

package entity

import "github.com/google/uuid"

type GameAction struct {
	ID         string
	PlayerID   string
	PlayerName string
	TeamID     string
	Minute     int
	Action     string
	Score      int
}

func NewGameAction(playerID, action string, minute, score int) *GameAction {
	return &GameAction{
		ID:       uuid.New().String(),
		PlayerID: playerID,
		Minute:   minute,
		Action:   action,
		Score:    score,
	}
}

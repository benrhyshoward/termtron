package model

//GameState : enum to track the current state of the game
type GameState int

const (
	Menu GameState = iota
	Countdown
	Playing
	Over
)

//Game : main object to track the current state of the game
type Game struct {
	State     GameState
	Players   []*Player
	Countdown int
}

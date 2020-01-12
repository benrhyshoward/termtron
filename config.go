package main

import (
	"time"

	"github.com/benrhyshoward/termtron/model"

	"github.com/nsf/termbox-go"
)

//GameStepLength : shorter step -> faster game
var GameStepLength = time.Duration(50) * time.Millisecond

//AverageAIDirectionChangePeriod :
//Decrease/increase value to make AI direction changes more/less frequent
//More frequent changes tend to make the AI perform worse
//Values < 1 will cause a panic
var AverageAIDirectionChangePeriod = 20

//CountdownLength : number of seconds to countdown before the start of each game
var CountdownLength = 3

//TerminalWidth : updated in game loop
var TerminalWidth int

//TerminalHeight : updated in game loop
var TerminalHeight int

//ControlPresets : mappings between key presses and directions
var ControlPresets = []model.ControlScheme{
	model.ControlScheme{
		Name: "Arrow Keys",
		Controls: map[termbox.Event]model.Direction{
			termbox.Event{Type: termbox.EventKey, Key: termbox.KeyArrowUp}:    model.Up,
			termbox.Event{Type: termbox.EventKey, Key: termbox.KeyArrowDown}:  model.Down,
			termbox.Event{Type: termbox.EventKey, Key: termbox.KeyArrowLeft}:  model.Left,
			termbox.Event{Type: termbox.EventKey, Key: termbox.KeyArrowRight}: model.Right,
		},
	},
	model.ControlScheme{
		Name: "[W][S][A][D]",
		Controls: map[termbox.Event]model.Direction{
			termbox.Event{Type: termbox.EventKey, Ch: 'w'}: model.Up,
			termbox.Event{Type: termbox.EventKey, Ch: 's'}: model.Down,
			termbox.Event{Type: termbox.EventKey, Ch: 'a'}: model.Left,
			termbox.Event{Type: termbox.EventKey, Ch: 'd'}: model.Right,
		},
	},
	model.ControlScheme{
		Name: "[I][J][K][L]",
		Controls: map[termbox.Event]model.Direction{
			termbox.Event{Type: termbox.EventKey, Ch: 'i'}: model.Up,
			termbox.Event{Type: termbox.EventKey, Ch: 'k'}: model.Down,
			termbox.Event{Type: termbox.EventKey, Ch: 'j'}: model.Left,
			termbox.Event{Type: termbox.EventKey, Ch: 'l'}: model.Right,
		},
	},
	model.ControlScheme{
		Name: "[G][V][B][N]",
		Controls: map[termbox.Event]model.Direction{
			termbox.Event{Type: termbox.EventKey, Ch: 'g'}: model.Up,
			termbox.Event{Type: termbox.EventKey, Ch: 'b'}: model.Down,
			termbox.Event{Type: termbox.EventKey, Ch: 'v'}: model.Left,
			termbox.Event{Type: termbox.EventKey, Ch: 'n'}: model.Right,
		},
	},
	model.ControlScheme{
		Name: "AI",
	},
}

//DefaultPlayers : initial player objects to populate game state
var DefaultPlayers = []*model.Player{
	&model.Player{
		Id:            1,
		Color:         termbox.ColorRed,
		Alive:         true,
		ControlScheme: ControlPresets[0],
	},
	&model.Player{
		Id:            2,
		Color:         termbox.ColorBlue,
		Alive:         true,
		ControlScheme: ControlPresets[1],
	},
	&model.Player{
		Id:            3,
		Color:         termbox.ColorGreen,
		Alive:         true,
		ControlScheme: ControlPresets[2],
	},
	&model.Player{
		Id:            4,
		Color:         termbox.ColorYellow,
		Alive:         true,
		ControlScheme: ControlPresets[3],
	},
}

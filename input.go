package main

import (
	"math/rand"

	"github.com/benrhyshoward/termtron/model"

	"github.com/nsf/termbox-go"
)

func handleInput(game *model.Game, quit chan bool) {
	//blocks until the next event
	ev := termbox.PollEvent()

	switch game.State {
	case model.Menu:

		if ev.Type == termbox.EventKey {
			switch {
			case ev.Key == termbox.KeyArrowUp:
				if len(game.Players) < 4 {
					addPlayer(game)
				}
			case ev.Key == termbox.KeyArrowDown:
				if len(game.Players) > 2 {
					removePlayer(game)
				}
			case ev.Ch == '1':
				cycleControlScheme(game.Players, 0)
			case ev.Ch == '2':
				cycleControlScheme(game.Players, 1)
			case ev.Ch == '3':
				if len(game.Players) > 2 {
					cycleControlScheme(game.Players, 2)
				}
			case ev.Ch == '4':
				if len(game.Players) > 3 {
					cycleControlScheme(game.Players, 3)
				}
			case ev.Key == termbox.KeySpace:
				populateStartingPositions(game.Players)
				randomiseDirections(game.Players)
				game.State = model.Countdown
			}
		}
	case model.Playing:
		//check if any player controls have been pressed and update player directions accordingly
		for _, player := range game.Players {
			for control, direction := range player.ControlScheme.Controls {
				if ev.Type == control.Type && ev.Ch == control.Ch && ev.Key == control.Key {

					//stop a player from turning back into themselves
					if (player.Direction == model.Up && direction == model.Down) ||
						(player.Direction == model.Down && direction == model.Up) ||
						(player.Direction == model.Left && direction == model.Right) ||
						(player.Direction == model.Right && direction == model.Left) {
						break
					}

					player.Direction = direction
				}
			}
		}
	case model.Over:
		if ev.Type == termbox.EventKey {
			switch {
			//restart game
			case ev.Key == termbox.KeySpace:
				for _, player := range game.Players {
					player.Alive = true
				}
				populateStartingPositions(game.Players)
				randomiseDirections(game.Players)
				game.State = model.Countdown
			//back to main menu
			case ev.Ch == 'm':
				for _, player := range game.Players {
					player.Alive = true
				}
				populateStartingPositions(game.Players)
				randomiseDirections(game.Players)
				game.State = model.Menu
			}
		}
	}

	//can quit at any time
	if ev.Type == termbox.EventKey {
		switch {
		case ev.Ch == 'q' || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
			quit <- true
		}
	}
}

func addPlayer(game *model.Game) {
	game.Players = append(game.Players, DefaultPlayers[len(game.Players)])
}

func removePlayer(game *model.Game) {
	game.Players = game.Players[:len(game.Players)-1]
}

func populateStartingPositions(players []*model.Player) {
	startingPositions := generateStartingPositions(len(players))
	for i, player := range players {
		player.Body = startingPositions[i]
	}
}

func randomiseDirections(players []*model.Player) {
	for _, player := range players {
		player.Direction = model.Direction(rand.Intn(4))
	}
}

func cycleControlScheme(players []*model.Player, playerIndex int) {
	var controlPresetIndex int

	for i, controlScheme := range ControlPresets {
		if players[playerIndex].ControlScheme.Name == controlScheme.Name {
			controlPresetIndex = i
		}
	}

	if controlPresetIndex+1 == len(ControlPresets) {
		controlPresetIndex = 0
	} else {
		controlPresetIndex++
	}

	players[playerIndex].ControlScheme = ControlPresets[controlPresetIndex]

}

func generateStartingPositions(numberOfPlayers int) [][]model.Point {
	switch numberOfPlayers {
	case 2:
		return [][]model.Point{
			[]model.Point{
				model.Point{
					X: TerminalWidth / 4,
					Y: TerminalHeight / 2,
				},
			},
			[]model.Point{
				model.Point{
					X: 3 * TerminalWidth / 4,
					Y: TerminalHeight / 2,
				},
			},
		}
	case 3:
		return [][]model.Point{
			[]model.Point{
				model.Point{
					X: TerminalWidth / 4,
					Y: TerminalHeight / 4,
				},
			},
			[]model.Point{
				model.Point{
					X: 3 * TerminalWidth / 4,
					Y: TerminalHeight / 4,
				},
			},
			[]model.Point{
				model.Point{
					X: TerminalWidth / 2,
					Y: 3 * TerminalHeight / 4,
				},
			},
		}
	case 4:
		return [][]model.Point{
			[]model.Point{
				model.Point{
					X: TerminalWidth / 4,
					Y: TerminalHeight / 4,
				},
			},
			[]model.Point{
				model.Point{
					X: 3 * TerminalWidth / 4,
					Y: TerminalHeight / 4,
				},
			},
			[]model.Point{
				model.Point{
					X: TerminalWidth / 4,
					Y: 3 * TerminalHeight / 4,
				},
			},
			[]model.Point{
				model.Point{
					X: 3 * TerminalWidth / 4,
					Y: 3 * TerminalHeight / 4,
				},
			},
		}
	default:
		return nil
	}
}

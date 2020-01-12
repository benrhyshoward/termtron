package main

import (
	"strconv"

	"github.com/benrhyshoward/termtron/model"

	"github.com/nsf/termbox-go"
)

func draw(game *model.Game) {
	TerminalWidth, TerminalHeight = termbox.Size()
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	switch game.State {
	case model.Menu:

		drawMainMenu(game.Players)

	case model.Countdown:

		drawPlayers(game.Players)
		drawCountdownTimer(game)
		drawPlayerDirections(game.Players)

	case model.Playing:

		drawPlayers(game.Players)

	case model.Over:

		drawPlayers(game.Players)
		drawResults(game.Players, 60, 15, termbox.ColorWhite, termbox.ColorBlack)
	}

	termbox.Flush()
}

func drawMainMenu(players []*model.Player) {
	textColour := termbox.ColorBlack
	backgroundColour := termbox.ColorWhite
	padding := 3
	width := TerminalWidth - 2*padding
	height := TerminalHeight - 2*padding
	drawRectangle(TerminalWidth/2-width/2, TerminalHeight/2-height/2, width, height, backgroundColour)

	drawCenteredText("TermTron.",
		2*padding, textColour, backgroundColour)
	drawCenteredText("Don't crash into edges, other players, or yourself.",
		3*padding, textColour, backgroundColour)
	drawCenteredText("Last one standing is the winner.",
		3*padding+1, textColour, backgroundColour)
	drawCenteredText("Change the number of players with [Up]/[Down].",
		4*padding, textColour, backgroundColour)
	drawCenteredText("Toggle a player's control scheme with their number [1],[2],[3],[4].",
		4*padding+1, textColour, backgroundColour)
	drawCenteredText("Resize the terminal window for more or less game space.",
		4*padding+2, textColour, backgroundColour)

	//Drawing current player configuration
	for i, player := range players {
		x := (TerminalWidth / (len(players) + 1)) * (i + 1)

		termbox.SetCell(x-1, 6*padding, ' ', termbox.ColorWhite, player.Color)
		termbox.SetCell(x, 6*padding, rune(strconv.Itoa(player.Id)[0]), termbox.ColorWhite, player.Color)
		termbox.SetCell(x+1, 6*padding, ' ', termbox.ColorWhite, player.Color)

		drawText(player.ControlScheme.Name, x, 6*padding+2, termbox.ColorBlack, termbox.ColorWhite)

	}

	drawCenteredText("Press [Space] to start, or [q] to quit.",
		8*padding, textColour, backgroundColour)

}

func drawPlayers(players []*model.Player) {

	for _, player := range players {
		for i, v := range player.Body {
			char := ' '

			//head of player
			if i == len(player.Body)-1 {
				char = 'o'
			}

			//dead players body populated with 'X's
			if player.Alive == false {
				char = 'X'
			}
			termbox.SetCell(v.X, v.Y, char, termbox.ColorBlack, player.Color)
		}
	}
}

func drawCountdownTimer(game *model.Game) {
	termbox.SetCell(TerminalWidth/2, TerminalHeight/2, rune(strconv.Itoa(game.Countdown)[0]), termbox.ColorWhite, termbox.ColorBlack)
}

func drawPlayerDirections(players []*model.Player) {
	for _, player := range players {
		var directionCell model.Point
		var directionChar rune
		switch player.Direction {
		case model.Up:
			directionCell = model.Point{
				X: player.Head().X,
				Y: player.Head().Y - 1,
			}
			directionChar = '↑'
		case model.Down:
			directionCell = model.Point{
				X: player.Head().X,
				Y: player.Head().Y + 1,
			}
			directionChar = '↓'
		case model.Right:
			directionCell = model.Point{
				X: player.Head().X + 1,
				Y: player.Head().Y,
			}
			directionChar = '→'
		case model.Left:
			directionCell = model.Point{
				X: player.Head().X - 1,
				Y: player.Head().Y,
			}
			directionChar = '←'
		}
		termbox.SetCell(directionCell.X, directionCell.Y, directionChar, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func drawResults(players []*model.Player, width int, height int, boxColour termbox.Attribute, textColour termbox.Attribute) {
	padding := 2
	topOfBox := TerminalHeight/2 - height/2
	drawRectangle(TerminalWidth/2-width/2, topOfBox, width, height, boxColour)

	drawCenteredText("Game Over", topOfBox+padding, textColour, boxColour)

	var survivors []*model.Player
	for _, player := range players {
		if player.Alive {
			survivors = append(survivors, player)
		}
	}
	if len(survivors) == 1 {
		drawCenteredText("Player "+strconv.Itoa(survivors[0].Id)+" won!", topOfBox+2*padding, textColour, boxColour)
		drawCenteredText("                         o", topOfBox+3*padding, termbox.ColorBlack, survivors[0].Color)
	} else {
		drawCenteredText("It's a tie!", topOfBox+2*padding+1, textColour, boxColour)
	}

	drawCenteredText("[Space] to restart", topOfBox+4*padding, textColour, boxColour)
	drawCenteredText("[m] to go back to main menu", topOfBox+5*padding, textColour, boxColour)
	drawCenteredText("[q] to quit", topOfBox+6*padding, textColour, boxColour)

}

func drawCenteredText(text string, y int, textColour termbox.Attribute, backgroundColour termbox.Attribute) {
	drawText(text, TerminalWidth/2, y, textColour, backgroundColour)
}

func drawText(text string, x int, y int, textColour termbox.Attribute, backgroundColour termbox.Attribute) {
	for i := 0; i < len(text); i++ {
		termbox.SetCell(x-len(text)/2+i, y, rune(text[i]), textColour, backgroundColour)
	}
}

func drawRectangle(x int, y int, width int, height int, color termbox.Attribute) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			termbox.SetCell(x+i, y+j, ' ', color, color)
		}
	}
}

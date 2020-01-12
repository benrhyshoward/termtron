package main

import (
	"math/rand"
	"time"

	"github.com/benrhyshoward/termtron/model"

	"github.com/nsf/termbox-go"
)

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	rand.Seed(time.Now().UnixNano())

	//initialising object to hold the state of the game
	//this object is mutated by the input and update loops, and is read by the draw loop
	game := &model.Game{
		State: model.Menu,
		Players: []*model.Player{
			DefaultPlayers[0],
			DefaultPlayers[1],
			DefaultPlayers[2],
			DefaultPlayers[3],
		},
	}

	quit := make(chan bool)

	//update game state based on user input
	go inputLoop(game, quit)

	//iterate game state
	go updateLoop(game)

	//render current game state to the screen
	go drawLoop(game)

	<-quit
}

func inputLoop(game *model.Game, quit chan bool) {
	for {
		handleInput(game, quit)
	}
}

func updateLoop(game *model.Game) {
	for {
		update(game)

		//length of sleep determines the speed of the game. shorter step -> faster game
		time.Sleep(GameStepLength)
	}
}

func drawLoop(game *model.Game) {
	for {
		draw(game)
	}
}

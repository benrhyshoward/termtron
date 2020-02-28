package main

import (
	"math/rand"
	"time"

	"github.com/benrhyshoward/termtron/model"
)

func update(game *model.Game) {
	switch game.State {
	case model.Menu:
	case model.Countdown:
		for i := CountdownLength; i > 0; i-- {
			game.Countdown = i
			time.Sleep(time.Second)
		}
		game.State = model.Playing
	case model.Playing:
		aiControl(game.Players)
		iteratePlayerLocations(game.Players)
		checkForWinner(game)
	case model.Over:
	}
}

func iteratePlayerLocations(players []*model.Player) {
	for _, player := range players {
		if player.Alive {
			player.Step()
		}
	}

	for _, player := range players {
		if player.Alive == true {
			head := player.Head()
			tail := player.Tail()

			if pointOutOfBounds(head) {
				player.Alive = false
			}

			for _, otherPlayer := range players {
				if otherPlayer.Id != player.Id {
					//check if head intersects with any other players' body
					for _, point := range otherPlayer.Body {
						if head == point {
							player.Alive = false
							break
						}
					}
				} else {
					//check if head intersects with any part of its own tail
					for _, point := range tail {
						if head == point {
							player.Alive = false
							break
						}
					}
				}

				if player.Alive == false {
					break
				}
			}
		}
	}
}

func checkForWinner(game *model.Game) {
	survivors := 0
	for _, player := range game.Players {
		if player.Alive {
			survivors++
		}
	}
	if survivors <= 1 {
		game.State = model.Over
	}
}

//aiControl: AI deciding whether to make any direction changes this step
func aiControl(players []*model.Player) {
	for _, player := range players {

		if player.ControlScheme.Name == "AI" {

			nextLocation := player.NextLocation()

			//if we are about to hit a wall or another player, attempt to change direction
			if !isPointSafeForPlayer(nextLocation, player, players) {
				if player.Direction == model.Up || player.Direction == model.Down {
					leftSafe := isPointSafeForPlayer(model.Point{
						X: player.Head().X - 1,
						Y: player.Head().Y,
					}, player, players)
					rightSafe := isPointSafeForPlayer(model.Point{
						X: player.Head().X + 1,
						Y: player.Head().Y,
					}, player, players)
					if leftSafe && rightSafe {
						//randomly go Left or Right
						player.Direction = model.Direction(rand.Intn(2) + 2)
					} else if rightSafe {
						player.Direction = model.Right
					} else if leftSafe {
						player.Direction = model.Left
					} else {
						//nowhere is safe, do nothing
					}
				} else {
					upSafe := isPointSafeForPlayer(model.Point{
						X: player.Head().X,
						Y: player.Head().Y - 1,
					}, player, players)
					downSafe := isPointSafeForPlayer(model.Point{
						X: player.Head().X,
						Y: player.Head().Y + 1,
					}, player, players)
					if upSafe && downSafe {
						//randomly go Up or Down
						player.Direction = model.Direction(rand.Intn(2))
					} else if upSafe {
						player.Direction = model.Up
					} else if downSafe {
						player.Direction = model.Down
					} else {
						//nowhere is safe, do nothing
					}
				}
			} else {
				//occasionally attempt to make a random direction change, helps stop the AI from just hugging walls
				//only makes the direction change if it's safe

				if rand.Intn(AverageAIDirectionChangePeriod) == 0 {
					newDirection := model.Direction(rand.Intn(4))
					switch newDirection {
					case model.Up:
						if isPointSafeForPlayer(model.Point{
							X: player.Head().X,
							Y: player.Head().Y - 1,
						}, player, players) {
							player.Direction = newDirection
						}
					case model.Down:
						if isPointSafeForPlayer(model.Point{
							X: player.Head().X,
							Y: player.Head().Y + 1,
						}, player, players) {
							player.Direction = newDirection
						}
					case model.Left:
						if isPointSafeForPlayer(model.Point{
							X: player.Head().X - 1,
							Y: player.Head().Y,
						}, player, players) {
							player.Direction = newDirection
						}
					case model.Right:
						if isPointSafeForPlayer(model.Point{
							X: player.Head().X + 1,
							Y: player.Head().Y,
						}, player, players) {
							player.Direction = newDirection
						}
					}
				}
			}
		}
	}
}

func isPointSafeForPlayer(point model.Point, player *model.Player, players []*model.Player) bool {
	if pointOutOfBounds(point) {
		return false
	}

	for _, otherPlayer := range players {
		//check if point intersects with any player's body (including itself)
		for _, bodyPart := range otherPlayer.Body {
			if bodyPart == point {
				return false
			}
		}

		//check if point intersects with any other player's next location
		if otherPlayer.Id != player.Id && point == otherPlayer.NextLocation() {
			return false
		}
	}
	return true
}

func pointOutOfBounds(point model.Point) bool {
	return point.X < 0 || point.X >= ScaledTerminalWidth || point.Y < 0 || point.Y >= ScaledTerminalHeight
}

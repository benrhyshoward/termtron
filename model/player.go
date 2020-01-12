package model

import (
	"github.com/nsf/termbox-go"
)

// Direction : enum for each of the four directions
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// Player : object to store the state of a particular player and how this player is controlled
type Player struct {
	Id            int
	Direction     Direction
	Body          []Point
	Color         termbox.Attribute
	Alive         bool
	ControlScheme ControlScheme
}

type ControlScheme struct {
	Name     string
	Controls map[termbox.Event]Direction
}

func (p *Player) Head() Point {
	return p.Body[len(p.Body)-1]
}

//Tail : returns array of points in a player's body excluding their head
func (p *Player) Tail() []Point {
	return p.Body[:len(p.Body)-1]
}

//Step : adding the player's next position onto their body
func (p *Player) Step() {
	p.Body = append(p.Body, p.NextLocation())
}

func (p *Player) NextLocation() Point {

	head := p.Head()

	switch p.Direction {
	case Up:
		head.Y--
	case Down:
		head.Y++
	case Right:
		head.X++
	case Left:
		head.X--
	}

	return head
}

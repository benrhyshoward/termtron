package main

import (
	"github.com/nsf/termbox-go"
)

// SetCellScaled : draw each point as a set of cells based on scaling factors
func SetCellScaled(x, y int, ch rune, fg, bg termbox.Attribute) {
	for i := 0; i < XScalingFactor; i++ {
		for j := 0; j < YScalingFactor; j++ {
			termbox.SetCell(XScalingFactor*x+i, YScalingFactor*y+j, ch, fg, bg)
		}
	}
}

// TermSizeScaled : return the termial width and height accouting for scaling factors
func TermSizeScaled() (width int, height int) {
	termWidth, termHeight := termbox.Size()

	return termWidth / XScalingFactor, termHeight / YScalingFactor
}

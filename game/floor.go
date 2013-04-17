package game

import (
	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/cotta/display"
)

type Floor interface {
	display.Drawable
	Copy() Floor
}

type ColorFloor struct {
	Color allegro.Color
}

func (cf ColorFloor) GetCharacter() string {
	return "."
}

func (cf ColorFloor) GetBGColor() allegro.Color {
	return cf.Color
}

func (cf ColorFloor) GetFGColor() allegro.Color {
	return cf.Color
}

func (cf ColorFloor) Copy() Floor {
	return Floor(ColorFloor{cf.Color})
}
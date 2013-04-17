package display

import "github.com/bluepeppers/allegro"

type Drawable interface {
	GetGlyph() string
	GetFGColor() allegro.Color
	GetBGColor() allegro.Color
}
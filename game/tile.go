package game

import "allegro"

type Tile struct {
	occupier *Entity
	floor    Floor

	cp *Tile
}

type Entity interface {
	Ticker
	GetCharacter() string
	GetFGColor() allegro.Color

	Copy() *Entity
}

func CreateTile(occupier *Entity, floor Floor) Tile {
	return Tile{occupier, floor, nil}
}

func (t *Tile) Copy() *Tile {
	if t.cp != nil {
		return t.cp
	}

	dest := new(Tile)
	t.cp = dest
	if t.occupier != nil {
		dest.occupier = (*t.occupier).Copy()
	}
	dest.floor = t.floor.Copy()
	return dest
}

func (t *Tile) GetCharacter() string {
	if t.occupier != nil {
		return (*t.occupier).GetCharacter()
	}
	return t.floor.GetCharacter()
}

func (t *Tile) GetFGColor() allegro.Color {
	if t.occupier != nil {
		return (*t.occupier).GetFGColor()
	}
	return t.floor.GetFGColor()
}

func (t *Tile) GetBGColor() allegro.Color {
	return t.floor.GetBGColor()
}

func (t *Tile) Tick(g *GameState) {
	if t.occupier != nil {
		go (*t.occupier).Tick(g)
	}
}

package game

import (
	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/danckelmann/resources"
)

type Tile struct {
	occupier *Entity
	floor    Entity
}

type Entity interface {
	Ticker

	GetSprites(*resources.ResourceManager) []*allegro.Bitmap
}

func CreateTile(occupier *Entity, floor Entity) Tile {
	return Tile{occupier, floor}
}

func (t *Tile) Tick(tick int) {
	if t.occupier != nil {
		go (*t.occupier).Tick(tick)
	}
}

func (t *Tile) GetSprites(rm *resources.ResourceManager) []*allegro.Bitmap {
	sprites := t.floor.GetSprites(rm)
	if t.occupier != nil {
		sprites = append(sprites, (*t.occupier).GetSprites(rm)...)
	}
	return sprites
}

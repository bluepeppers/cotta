package tile

import (
	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/danckelmann/resources"

	"github.com/bluepeppers/cotta/game/walker"
)

const (
	EMPTY_TILE_FLOOR = "roadTiles.grass"
)

type Tile interface {
	GetSprites(*resources.ResourceManager) []*allegro.Bitmap

	AdjacentWalker(walker.Walker)
	Tick(int)
}

type EmptyTile struct {
	
}

func CreateEmptyTile() *EmptyTile {
	return &EmptyTile{}
}

func (et *EmptyTile) GetSprites(rm *resources.ResourceManager) []*allegro.Bitmap {
	floor := rm.GetTileOrDefault(EMPTY_TILE_FLOOR)
	return []*allegro.Bitmap{floor}
}

func (et *EmptyTile) AdjacentWalker(w walker.Walker) {}

func (et *EmptyTile) Tick(i int) {}
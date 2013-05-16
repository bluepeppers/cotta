package tile

import (
	"github.com/bluepeppers/danckelmann/resources"

	"github.com/bluepeppers/cotta/game/walker"
)

const (
	EMPTY_TILE_FLOOR = "roadTiles.grass"
)

type Tile interface {
	GetSprites(*resources.ResourceManager) []*resources.Bitmap

	AdjacentWalker(walker.Walker)
	Tick(int)
}

type EmptyTile struct {
	
}

func CreateEmptyTile() *EmptyTile {
	return &EmptyTile{}
}

func (et *EmptyTile) GetSprites(rm *resources.ResourceManager) []*resources.Bitmap {
	floor := rm.GetTileOrDefault(EMPTY_TILE_FLOOR)
	return []*resources.Bitmap{floor}
}

func (et *EmptyTile) AdjacentWalker(w walker.Walker) {}

func (et *EmptyTile) Tick(i int) {}
package game

import (
	"github.com/bluepeppers/danckelmann/resources"
	"github.com/bluepeppers/cotta/game/walker"
)

type TileFloor struct {
	spriteName string
}

func CreateTileFloor(spriteName string) TileFloor {
	return TileFloor{spriteName}
}

func (tf TileFloor) Tick(tick int) {
	return
}

func (tf TileFloor) GetSprites(rm *resources.ResourceManager) []*resources.Bitmap {
	return []*resources.Bitmap{rm.GetTileOrDefault(tf.spriteName)}
}

func (tf TileFloor) AdjacentWalker(w walker.Walker) {}
package game

import (
	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/danckelmann/resources"
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

func (tf TileFloor) GetSprites(rm *resources.ResourceManager) []*allegro.Bitmap {
	return []*allegro.Bitmap{rm.GetTileOrDefault(tf.spriteName)}
}

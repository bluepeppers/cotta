package game

import (
	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/danckelmann/config"
)

type WorldMap struct {
	width, height int
	tiles         []Tile
}

func CreateWorldMap(conf *allegro.Config) *WorldMap {
	var wm WorldMap
	wm.width = config.GetInt(conf, "map", "width", 10)
	wm.height = config.GetInt(conf, "map", "height", 10)
	wm.tiles = make([]Tile, wm.width*wm.height)
	for i, _ := range wm.tiles {
		f := CreateTileFloor("grass")
		wm.tiles[i] = CreateTile(nil, f)
	}
	return &wm
}

func (wm *WorldMap) Tick(tick int) {
	for i, _ := range wm.tiles {
		wm.tiles[i].Tick(tick)
	}
}

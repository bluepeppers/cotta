package game

import (
	"sync"
	
	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/danckelmann/config"

	"github.com/bluepeppers/cotta/game/tile"
	"github.com/bluepeppers/cotta/game/roads"
)

type WorldMap struct {
	// Lock when modifying the tiles variable. Do not need to lock when
	// modifying underlying tiles
	tileLock sync.RWMutex
	width, height int
	tiles         [][]tile.Tile

	roadNetwork *roads.RoadNetwork
}

func CreateWorldMap(conf *allegro.Config) *WorldMap {
	var wm WorldMap
	wm.width = config.GetInt(conf, "map", "width", 20)
	wm.height = config.GetInt(conf, "map", "height", 20)
	tilesParent := make([]tile.Tile, wm.width*wm.height)
	wm.tiles = make([][]tile.Tile, wm.width)
	for x := 0; x < wm.width; x++ {
		wm.tiles[x] = tilesParent[:wm.height]
		tilesParent = tilesParent[wm.height:]
	}
	for x := 0; x < wm.width; x++ {
		for y := 0; y < wm.height; y++ {
			wm.tiles[x][y] = tile.CreateEmptyTile()
		}
	}

	wm.roadNetwork = roads.CreateRoadNetwork(&wm)

	return &wm
}

func (wm *WorldMap) Tick(tick int) {
	for x := 0; x < wm.width; x++ {
		for y := 0; y < wm.height; y++ {
			wm.tiles[x][y].Tick(tick)
		}
	}
}

func (wm *WorldMap) SetTile(x, y int, tile tile.Tile) bool {
	if x < 0 || x >= wm.width || y < 0 || y >= wm.height {
		return false
	}
	wm.tileLock.Lock()
	wm.tiles[x][y] = tile
	wm.tileLock.Unlock()
	return true
}

func (wm *WorldMap) GetTile(x, y int) (tile.Tile, bool) {
	if x < 0 || x >= wm.width || y < 0 || y >= wm.height {
		return nil, false
	}
	wm.tileLock.RLock()
	defer wm.tileLock.RUnlock()
	return wm.tiles[x][y], true
}

func (wm *WorldMap) GetDimensions() (int, int) {
	// Don't think we need to lock as stuff never gets resized.
	return wm.width, wm.height
}
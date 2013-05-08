package roads

import (
	"sync"
	
	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/danckelmann/resources"

	"github.com/bluepeppers/cotta/game/walker"
)

const (
	EAST = 1 << iota
	WEST
	NORTH
	SOUTH
	)

var adjacencyToResource = map[byte]string{
	// Default to north south
0: "roadTiles.roadNorth",
	
	// Road ends
EAST: "roadTiles.roadEndEast",
WEST: "roadTiles.roadEndWest",
SOUTH: "roadTiles.roadEndSouth",
NORTH: "roadTiles.roadEndNorth",

	// Straight roads
	EAST | WEST: "roadTiles.roadEast",
	NORTH | SOUTH: "roadTiles.roadNorth",

	// Corners
	SOUTH | EAST: "roadTiles.roadCornerES",
	EAST | NORTH: "roadTiles.roadCornerNE",
	WEST | NORTH: "roadTiles.roadCornerNW",
	SOUTH | WEST: "roadTiles.roadCornerWS",

	// T junctions
	SOUTH | EAST | NORTH: "roadTiles.roadTEast",
	EAST | NORTH | WEST: "roadTiles.roadTNorth",
	NORTH | WEST | SOUTH: "roadTiles.roadTWest",
	WEST | SOUTH | EAST: "roadTiles.roadTSouth",

	// Crossroads
	WEST | EAST | NORTH | SOUTH: "roadTiles.crossroad"}

type RoadTile struct {
	roadNetwork *RoadNetwork

	adjLock sync.RWMutex
	adjacentRoads byte
	floorName string
}

func CreateRoadTile(rn *RoadNetwork, adjacent byte) *RoadTile {
	var tile RoadTile
	tile.roadNetwork = rn
	tile.adjacentRoads = adjacent
	tile.floorName = adjacencyToResource[tile.adjacentRoads]
	return &tile
}

func (rt *RoadTile) GetSprites(rm *resources.ResourceManager) []*allegro.Bitmap {
	rt.adjLock.RLock()
	floorBmp := rm.GetTileOrDefault(rt.floorName)
	rt.adjLock.RUnlock()
	return []*allegro.Bitmap{floorBmp}
}

func (rt *RoadTile) RoadAdjacent(direction byte) {
	rt.adjLock.Lock()
	defer rt.adjLock.Unlock()
	rt.adjacentRoads = rt.adjacentRoads | direction
	rt.floorName = adjacencyToResource[rt.adjacentRoads]
}

func (rt *RoadTile) RoadNotAdjacent(direction byte) {
	rt.adjLock.Lock()
	defer rt.adjLock.Unlock()
	rt.adjacentRoads = rt.adjacentRoads & ^direction
	rt.floorName = adjacencyToResource[rt.adjacentRoads]
}

func (rt *RoadTile) AdjacentWalker(w walker.Walker) {
}

func (rt *RoadTile) Tick(i int) {}
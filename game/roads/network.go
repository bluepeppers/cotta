package roads

import (
	"github.com/bluepeppers/cotta/game/tile"
)

type WorldMap interface {
	SetTile(int, int, tile.Tile) bool
	GetTile(int, int) (tile.Tile, bool)
	GetDimensions() (int, int)
}

type RoadNetwork struct {
	worldMap WorldMap

	modifiers chan func(*RoadNetwork)
	
	network [][]*RoadTile
}

func CreateRoadNetwork(wm WorldMap) *RoadNetwork {
	var network RoadNetwork
	network.worldMap = wm

	w, h := wm.GetDimensions()
	networkParent := make([]*RoadTile, w * h)
	network.network = make([][]*RoadTile, w)
	for x := 0; x < w; x++ {
		network.network[x] = networkParent[:w]
		networkParent = networkParent[w:]
	}
	network.modifiers = make(chan func(*RoadNetwork))
	go network.runModifier()
	
	return &network
}

func (rn *RoadNetwork) runModifier() {
	for mod := range rn.modifiers {
		mod(rn)
	}
}

func (rn_ *RoadNetwork) AddRoad(x, y int) bool {
	retVal := false
	rn_.modifiers <- func (rn *RoadNetwork) {
		w, h := rn.worldMap.GetDimensions()
		if !(0 <= x && x < w) ||
			!(0 <= y && y < h) {
			return
		}
		currentTile, ok := rn.worldMap.GetTile(x, y)
		if !ok {
			return
		}
		_, empty := currentTile.(*tile.EmptyTile)
		if !empty {
			return
		}

		rx := x + w
		ry := y + w
		var adjacent byte

		if rn.network[x][(ry+1)%h] != nil {
			adjacent = adjacent | SOUTH
			rn.network[x][(ry+1)%h].RoadAdjacent(NORTH)
		}
		if rn.network[x][(ry-1)%h] != nil {
			adjacent = adjacent | NORTH
			rn.network[x][(ry-1)%h].RoadAdjacent(SOUTH)
		}
		if rn.network[(rx+1)%w][y] != nil {
			adjacent = adjacent | WEST
			rn.network[(rx+1)%w][y].RoadAdjacent(EAST)
		}
		if rn.network[(rx-1)%w][y] != nil {
			adjacent = adjacent | EAST
			rn.network[(rx-1)%w][y].RoadAdjacent(WEST)
		}
		newTile := CreateRoadTile(rn, adjacent)
		retVal = rn.worldMap.SetTile(x, y, newTile)
		if retVal {
			rn.network[x][y] = newTile
		}
	}
	return retVal
}
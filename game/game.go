package game

import "allegro"

import "github.com/bluepeppers/cotta/config"

const (
	LAG_THREASHHOLD = 0.5
	MAXIMUM_SPEED   = 1 / 60.
)

type Game struct {
	ticks   *allegro.Timer
	mutator *Mutator
}

type GameState struct {
	world   *WorldMap
	stopped bool
	stats   Stats

	cp *GameState
}

type Stats struct {
	tickN   int
	actionN int
}

type WorldMap struct {
	tiles [][]Tile

	cp *WorldMap
}

type Ticker interface {
	Tick(*GameState)
}

func CreateGame(conf *allegro.Config) *Game {
	g := new(Game)
	tickRate := config.GetInt(conf, "map", "tickRate", 30)
	g.ticks = allegro.CreateTimer(1 / float64(tickRate))

	state := CreateGameState(conf)
	g.mutator = CreateMutator(state)
	return g
}

func CreateGameState(conf *allegro.Config) *GameState {
	st := new(GameState)
	st.stopped = false
	st.world = CreateWorldMap(conf)
	return st
}

func (g *GameState) Copy() *GameState {
	if g.cp != nil {
		return g.cp
	}
	dest := new(GameState)
	g.cp = dest
	dest.world = g.world.Copy()
	dest.stopped = g.stopped
	dest.stats = g.stats
	return dest
}

func CreateWorldMap(conf *allegro.Config) *WorldMap {
	width := config.GetInt(conf, "map", "width", 100)
	height := config.GetInt(conf, "map", "height", 100)
	tiles := make([][]Tile, width)
	for x := 0; x < width; x++ {
		tiles[x] = make([]Tile, height)
		for y := 0; y < height; y++ {
			var f ColorFloor
			tiles[x][y] = CreateTile(nil, f)
		}
	}
	return &WorldMap{tiles: tiles}
}

func (wm *WorldMap) Copy() *WorldMap {
	if wm.cp != nil {
		return wm.cp
	}
	dest := new(WorldMap)
	wm.cp = dest

	dest.tiles = make([][]Tile, len(wm.tiles))
	for i, r := range wm.tiles {
		dest.tiles[i] = make([]Tile, len(r))
		for j, t := range r {
			dest.tiles[i][j] = *(t.Copy())
		}
	}
	return dest
}

func (g *Game) MainLoop() {
	g.mutator.Start()
	defer g.mutator.Stop()

	es := []*allegro.EventSource{g.ticks.GetEventSource()}
	queue := allegro.GetEvents(es)
	stopped := false
	for !stopped {
		g.mutator.stateLock.RLock()
		stopped = g.mutator.lastState.stopped
		g.mutator.stateLock.RUnlock()

		ev := <-queue
		if _, ok := ev.(allegro.TimerEvent); ok {
			g.mutator.Tick()
		}
	}
}

/*
// Returns true if a speed change should be done
func (g *GameState) CheckSpeedChange(ev allegro.TimerEvent) bool {
	diff := g.ticks.GetCount() - int64(ev.Count)
	threshhold := LAG_THREASHHOLD / g.ticks.GetSpeed()

	return diff > int64(threshhold)
}

func (g *GameState) ChangeSpeed(ev allegro.TimerEvent) {
	ns := g.ticks.GetSpeed() * 1.1

	g.ticks.SetSpeed(ns)
}
*/

func (g *GameState) Tick() {
	g.stats.tickN++
	g.world.Tick(g)
}

func (w *WorldMap) Tick(g *GameState) {
	for x := 0; x < len(w.tiles); x++ {
		for y := 0; y < len(w.tiles[x]); y++ {
			go w.tiles[x][y].Tick(g)
		}
	}
}

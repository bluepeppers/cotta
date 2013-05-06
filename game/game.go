package game

import (
	"sync"

	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/danckelmann/config"
	"github.com/bluepeppers/danckelmann/display"
	"github.com/bluepeppers/danckelmann/resources"
)

const (
	LAG_THREASHHOLD     = 0.5
	DEFAULT_TICKRATE    = 60
	DEFAULT_TILE_WIDTH  = 100
	DEFAULT_TILE_HEIGHT = 50
)

type GameEngine struct {
	displayEngine   *display.DisplayEngine
	displayConfig   display.DisplayConfig
	resourceManager *resources.ResourceManager

	world    *WorldMap
	stats    Stats
	tickRate int

	statusLock sync.RWMutex
	stopped    bool
}

type Stats struct {
	tickN   int
	actionN int
}

type Ticker interface {
	Tick(int)
}

func CreateGameEngine(conf *allegro.Config) *GameEngine {
	ge := new(GameEngine)
	ge.stopped = false
	ge.world = CreateWorldMap(conf)
	ge.tickRate = config.GetInt(conf, "game", "tickrate", DEFAULT_TICKRATE)

	tileWidth := config.GetInt(conf, "game", "tileWidth", DEFAULT_TILE_WIDTH)
	tileHeight := config.GetInt(conf, "game", "tileHeight", DEFAULT_TILE_HEIGHT)

	ge.displayConfig = display.DisplayConfig{
		ge.world.width, ge.world.height,
		tileWidth, tileHeight,
		allegro.CreateColorHTML("black")}
	return ge
}

func (ge *GameEngine) RegisterDisplayEngine(de *display.DisplayEngine) {
	ge.displayEngine = de
	ge.resourceManager = de.GetResourceManager()
}

func (ge *GameEngine) GetDisplayConfig() display.DisplayConfig {
	return ge.displayConfig
}

func (ge *GameEngine) GameFinished() {
	ge.statusLock.Lock()
	ge.stopped = true
	ge.statusLock.Unlock()
}

func (ge *GameEngine) MainLoop() {
	timer := allegro.CreateTimer(1 / float64(ge.tickRate))
	es := []*allegro.EventSource{timer.GetEventSource()}
	queue := allegro.GetEvents(es)
	stopped := false

	ge.statusLock.Lock()
	ge.stopped = false
	ge.statusLock.Unlock()
	tick := 0
	for !stopped {
		ev := <-queue
		if _, ok := ev.(allegro.TimerEvent); ok {
			ge.world.Tick(tick)
		}
		tick++

		ge.statusLock.RLock()
		stopped = ge.stopped
		ge.statusLock.RUnlock()
	}
}

func (ge *GameEngine) GetTile(x, y int) []*allegro.Bitmap {
	return ge.world.tiles[x*ge.displayConfig.MapW+y].GetSprites(ge.resourceManager)
}

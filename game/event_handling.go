package game

import (
	"github.com/bluepeppers/allegro"

	"github.com/bluepeppers/cotta/game/tile"
)

const (
	SCROLL_SPEED = 10
)

func (g *GameEngine) eventHandler() {
	src := g.displayEngine.Display.GetEventSource()
	defer src.StopGetEvents()
	es := []*allegro.EventSource{src,
		allegro.GetKeyboardEventSource(),
		allegro.GetMouseEventSource()}
	queue := allegro.GetEvents(es)
	stopped := false
	for !stopped {
		ev := <- queue
		switch tev := ev.(type) {
		case allegro.KeyCharEvent:
			g.handleKeyChar(tev)
		case allegro.MouseButtonDown:
			g.handleMouseDown(tev)
		}
		g.statusLock.RLock()
		stopped = g.stopped
		g.statusLock.RUnlock()
	}
}

func (g *GameEngine) handleKeyChar(ev allegro.KeyCharEvent) {
	var x, y int
	switch ev.Keycode {
	case allegro.KEY_LEFT:
		x = -SCROLL_SPEED
	case allegro.KEY_RIGHT:
		x = SCROLL_SPEED
	case allegro.KEY_UP:
		y = -SCROLL_SPEED
	case allegro.KEY_DOWN:
		y = SCROLL_SPEED
	}
	viewport := g.displayEngine.GetViewport()
	viewport.Move(x, y)
	g.displayEngine.SetViewport(viewport)
}

func (g *GameEngine) handleMouseDown(event allegro.MouseButtonDown) {
	switch event.Button {
	case 3:
		go g.startScrolling(event)
	case 2:
		go g.addBuilding(event)
	case 1:
		go g.addRoad(event)
	}
}

func (g *GameEngine) startScrolling(start allegro.MouseButtonDown) {
	timer := allegro.CreateTimer(float64(1) / 30)
	mouseEvent := allegro.GetMouseEventSource()
	defer mouseEvent.StopGetEvents()
	es := []*allegro.EventSource{mouseEvent,
		timer.GetEventSource()}
	timer.Start()
	defer timer.Destroy()

	x, y := start.X, start.Y
	stopped := false
	for ev := range allegro.GetEvents(es) {
		switch tev := ev.(type) {
		case allegro.MouseButtonUp:
			if tev.Button == start.Button {
				stopped = true
			}
		case allegro.TimerEvent:
			viewport := g.displayEngine.GetViewport()
			viewport.Move((x - start.X) / 30, (y - start.Y) / 30)
			g.displayEngine.SetViewport(viewport)
		case allegro.MouseAxesEvent:
			x, y = tev.X, tev.Y
		}
		g.statusLock.RLock()
		stopped = stopped || g.stopped
		g.statusLock.RUnlock()
		if stopped {
			break
		}
	}
}

func (g *GameEngine) addRoad(event allegro.MouseButtonDown) {
	x, y := g.displayEngine.GetViewport().ScreenCoordinatesToTile(event.X, event.Y, g.displayConfig)
	g.world.roadNetwork.AddRoad(int(x), int(y))
}

func (g *GameEngine) addBuilding(event allegro.MouseButtonDown) {
	x, y := g.displayEngine.GetViewport().ScreenCoordinatesToTile(event.X, event.Y, g.displayConfig)
	current, ok := g.world.GetTile(int(x), int(y))
	if !ok {
		return
	}
	if _, ok = current.(*tile.EmptyTile); ok {
		g.world.SetTile(int(x), int(y), &TileFloor{"buildings.building1"})
	}
}
package game

import (
	"github.com/bluepeppers/allegro"
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
	viewport.X += x
	viewport.Y += y
	g.displayEngine.SetViewport(viewport)
}

func (g *GameEngine) handleMouseDown(event allegro.MouseButtonDown) {
	if event.Button == 3 {
		go g.startScrolling(event)
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
			viewport.X += (x - start.X) / 30
			viewport.Y += (y - start.Y) / 30
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

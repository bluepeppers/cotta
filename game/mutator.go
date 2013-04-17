package game

import "sync"

type Action interface {
	Read(*GameState)
	ReRead(*GameState)
	Apply(*GameState) bool
}

type Mutator struct {
	stateLock sync.RWMutex
	currState *GameState
	lastState *GameState

	starter sync.Once
	input   chan *Action
	tick    chan bool
}

func CreateMutator(g *GameState) *Mutator {
	i := make(chan *Action)
	t := make(chan bool)
	return &Mutator{currState: g, lastState: g.Copy(), input: i, tick: t}
}

func (m *Mutator) AddAction(a *Action) {
	go func() {
		m.stateLock.RLock()
		(*a).Read(m.lastState)
		m.stateLock.RUnlock()

		m.input <- a
	}()
}

func (m *Mutator) Start() {
	m.starter.Do(func() {
		go func() {
			m.run()
			var newonce sync.Once
			m.starter = newonce
		}()
	})
}

func (m *Mutator) run() {
	running := true
	for running {
		select {
		case action, ok := <-m.input:
			if !ok {
				running = false
				break
			}
			applied := (*action).Apply(m.currState)
			if !applied {
				m.stateLock.RLock()
				cp := m.lastState.Copy()
				m.stateLock.RUnlock()

				go func() {
					(*action).ReRead(cp)
					m.input <- action
				}()
			}
		case <-m.tick:
			m.stateLock.Lock()
			m.lastState = m.currState
			m.stateLock.Unlock()

			m.currState = m.lastState.Copy()

			go m.lastState.Tick()
		}
	}
}

func (m *Mutator) Tick() {
	m.tick <- true
}

func (m *Mutator) Stop() {
	close(m.input)
}

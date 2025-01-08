package engine

import "time"

// Engine holds the objects for running the game
// This engine will use a push/pop to transition between scenes
type Engine struct {
	currentState  IGameState
	stateStack    []IGameState
	running       bool
	lastFrameTime time.Time
}

func NewEngine(initialState IGameState) *Engine {
	return &Engine{
		currentState:  initialState,
		running:       true,
		lastFrameTime: time.Now(),
	}
}

func (engine *Engine) Run() {
	for engine.running {
		now := time.Now()
		deltaTime := now.Sub(engine.lastFrameTime).Seconds()
		engine.lastFrameTime = now

		//do inputs or commands

		//update current state
		engine.currentState.Update(deltaTime)

		//render current state
		engine.currentState.Render(deltaTime)

		//mimic 60fps for now
		time.Sleep(16 * time.Millisecond)
	}
}

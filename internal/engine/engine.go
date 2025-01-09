package engine

import (
	"relichunters/internal/gameapi"
	"time"
)

// Engine holds the objects for running the game
// This engine will use a push/pop to transition between scenes
type Engine struct {
	currentScene  gameapi.IGameScene
	stateStack    []gameapi.IGameScene
	running       bool
	lastFrameTime time.Time
	sceneManager  gameapi.ISceneManager
}

func NewEngine(sm gameapi.ISceneManager) *Engine {
	return &Engine{
		running:       true,
		lastFrameTime: time.Now(),
		sceneManager:  sm,
	}
}

func (engine *Engine) Run() {
	for engine.running {
		now := time.Now()
		deltaTime := now.Sub(engine.lastFrameTime).Seconds()
		engine.lastFrameTime = now

		//get current scene from scene manager
		engine.currentScene = engine.sceneManager.CurrentScene()

		//check input handler
		engine.currentScene.HandleInput()

		//update current scene
		engine.currentScene.Update(deltaTime)

		//render current scene
		engine.currentScene.Render(deltaTime)

		//mimic 60fps -- can do more to this later
		time.Sleep(16 * time.Millisecond)
	}
}

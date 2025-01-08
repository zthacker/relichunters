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
	inputHandler  gameapi.IInputHandler
}

func NewEngine(sm gameapi.ISceneManager, inputHandler gameapi.IInputHandler) *Engine {
	return &Engine{
		running:       true,
		lastFrameTime: time.Now(),
		sceneManager:  sm,
		inputHandler:  inputHandler,
	}
}

func (engine *Engine) Run() {
	for engine.running {
		now := time.Now()
		deltaTime := now.Sub(engine.lastFrameTime).Seconds()
		engine.lastFrameTime = now

		//get current scene from scene manager
		engine.currentScene = engine.sceneManager.CurrentScene()

		//handle commands from inputHandler
		cmds := engine.inputHandler.PollCommands(engine.currentScene)
		for _, cmd := range cmds {
			engine.currentScene.HandleInput(cmd)
		}

		//update current scene
		engine.currentScene.Update(deltaTime)

		//render current scene
		engine.currentScene.Render(deltaTime)

		//mimic 60fps -- can do more to this later
		time.Sleep(16 * time.Millisecond)
	}
}

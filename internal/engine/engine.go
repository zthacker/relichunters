package engine

import (
	"log"
	"relichunters/internal/gameapi"
	"time"
)

type Engine struct {
	running       bool
	currentScene  gameapi.IGameScene
	lastFrameTime time.Time
	game          gameapi.IGameApi
	log           *log.Logger
}

func NewEngine(game gameapi.IGameApi, logger *log.Logger) *Engine {
	return &Engine{
		running:       true,
		lastFrameTime: time.Now(),
		game:          game,
		log:           logger,
	}
}

func (engine *Engine) Run() {
	for engine.running {
		now := time.Now()
		deltaTime := now.Sub(engine.lastFrameTime).Seconds()
		engine.lastFrameTime = now

		//get current scene from scene manager
		engine.currentScene = engine.game.CurrentScene()
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

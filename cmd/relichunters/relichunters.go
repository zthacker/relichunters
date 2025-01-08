package main

import (
	"relichunters/internal/engine"
	"relichunters/internal/game"
	"relichunters/internal/scenemanager"
)

func main() {

	//create a new game
	newGame := game.NewGame()

	//create initial scene for the engine
	initialScene := newGame.CreateMenuScene()

	//make the scene manager
	sm := scenemanager.NewDefaultSceneManager(initialScene)

	//create a new game engine
	gameEngine := engine.NewEngine(sm)

	//run the engine
	gameEngine.Run()
}

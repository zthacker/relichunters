package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"relichunters/internal/engine"
	"relichunters/internal/game"
	"relichunters/internal/inputhandlers"
	"relichunters/internal/renderer"
	"relichunters/internal/scenemanager"
)

func main() {

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	if err = screen.Init(); err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	//input handler
	tcellInputHandler := inputhandlers.NewTCellInputHandler(screen)
	tcellRenderer := renderer.NewTCellRenderer(screen)

	//create a new game
	newGame := game.NewGame(tcellInputHandler, tcellRenderer)

	//create initial scene for the engine
	initialScene := newGame.CreateMenuScene()

	//make the scene manager
	sm := scenemanager.NewDefaultSceneManager(initialScene)

	//create a new game engine
	gameEngine := engine.NewEngine(sm)

	//run the engine
	gameEngine.Run()
}

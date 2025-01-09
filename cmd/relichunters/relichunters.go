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

	//game setup
	tcellInputHandler := inputhandlers.NewTCellInputHandler(screen)
	tcellRenderer := renderer.NewTCellRenderer(screen)
	sm := scenemanager.NewDefaultSceneManager(nil)
	newGame := game.NewGame(tcellInputHandler, tcellRenderer, sm)

	initialScene := newGame.CreateMenuScene()
	sm.SetScene(initialScene)

	//create a new game engine
	gameEngine := engine.NewEngine(sm)

	//run the engine
	gameEngine.Run()
}

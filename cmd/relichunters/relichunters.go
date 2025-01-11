package main

import (
	"encoding/json"
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
	"path/filepath"
	"relichunters/internal/engine"
	"relichunters/internal/game"
	"relichunters/internal/inputhandlers"
	"relichunters/internal/models"
	"relichunters/internal/renderer"
	"relichunters/internal/sceneregistry"
)

func main() {

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new logger
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("screen setup")
	screen, err := tcell.NewScreen()
	if err != nil {
		logger.Println(err)
		return
	}

	if err = screen.Init(); err != nil {
		logger.Println(err)
		return
	}
	defer screen.Fini()

	//scene registry from JSON file
	sr := sceneregistry.NewSceneRegistry()
	registerScenesFromJSON(sr, logger)

	//game setup
	logger.Println("game setup")
	tcellInputHandler := inputhandlers.NewTCellInputHandler(screen)
	tcellRenderer := renderer.NewTCellRenderer(screen)

	logger.Println("new game call")
	newGame := game.NewGame(sr, tcellInputHandler, tcellRenderer, logger)
	logger.Println("init call")
	newGame.InitGame()

	//create a new game engine
	logger.Println("new engine call")
	gameEngine := engine.NewEngine(newGame, logger)

	//run the engine
	logger.Println("engine run call")
	gameEngine.Run()
}

func registerScenesFromJSON(sr *sceneregistry.SceneRegistry, logger *log.Logger) {
	var allSceneDefs []models.SceneDefinition
	var sDefs []models.SceneDefinition

	registryData, err := os.ReadDir("sceneregistrydata")
	if err != nil {
		logger.Fatalln(err)
	}

	for _, file := range registryData {
		fBytes, err := os.ReadFile(filepath.Join("sceneregistrydata", file.Name()))
		if err != nil {
			logger.Fatalln(err)
		}
		err = json.Unmarshal(fBytes, &sDefs)
		if err != nil {
			logger.Fatalln(err)
		}
		allSceneDefs = append(allSceneDefs, sDefs...)
	}

	for _, def := range allSceneDefs {
		sr.LoadSceneDef(def)
	}
}

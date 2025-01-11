package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"gopkg.in/yaml.v2"
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
	registerScenes(sr, logger)

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

func registerScenes(sr *sceneregistry.SceneRegistry, logger *log.Logger) {
	logger.Println("Initializing Scene Registry")
	var sceneDefs []models.SceneDefinition
	var sDef models.SceneDefinition

	root := "./data"

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		logger.Println(d.Name())
		if !d.IsDir() {
			logger.Println(path)
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			err = yaml.Unmarshal(content, &sDef)
			if err != nil {
				return err
			}
			logger.Println(fmt.Sprintf("%+v", sDef))
			sceneDefs = append(sceneDefs, sDef)
		}

		return nil
	})

	if err != nil {
		logger.Fatalln(err)
	}

	logger.Println("Looping through All Scenes: %+v", sceneDefs)
	for _, def := range sceneDefs {
		logger.Println(fmt.Printf("%+v", def))
		sr.SetSceneDef(&def)
		scene, err := sr.GetDefinition(def.Key)
		if err != nil {
			logger.Fatalln(err)
		}
		fmt.Println(fmt.Sprintf("GET: %+v", scene))
	}
}

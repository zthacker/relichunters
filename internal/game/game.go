package game

import (
	"fmt"
	"log"
	"relichunters/internal/gameapi"
	"relichunters/internal/models"
	"relichunters/internal/sceneregistry"
	"relichunters/internal/scenes"
)

type Game struct {
	//persistent data
	data         *models.GameData
	input        gameapi.IInputHandler
	renderer     gameapi.Renderer
	registry     *sceneregistry.SceneRegistry
	currentScene gameapi.IGameScene
	stack        []gameapi.IGameScene
	log          *log.Logger
}

func NewGame(sr *sceneregistry.SceneRegistry, ih gameapi.IInputHandler, renderer gameapi.Renderer, logger *log.Logger) *Game {
	return &Game{input: ih, renderer: renderer, registry: sr, log: logger}
}

func (g *Game) InitGame() {
	g.log.Println("Initializing game")
	if err := g.SetScene(models.SceneKeyMainMenu); err != nil {
		g.log.Fatalln(err)
	}
	g.data = &models.GameData{}
	g.log.Println(fmt.Sprintf("Current scene: %v", g.currentScene))
}

func (g *Game) CreateScene(key models.SceneKey) (gameapi.IGameScene, error) {
	g.log.Println(fmt.Sprintf("Creating scene: %v", key))
	def, err := g.registry.GetDefinition(key)
	if err != nil {
		return nil, err
	}
	return g.CreateSceneFromDef(def)
}

func (g *Game) SetScene(key models.SceneKey) error {
	newScene, err := g.CreateScene(key)
	if err != nil {
		return err
	}
	if g.currentScene != nil {
		g.currentScene.OnExit()
	}
	g.currentScene = newScene
	g.currentScene.OnEnter()
	return nil
}

func (g *Game) PushScene(key models.SceneKey) error {
	newScene, err := g.CreateScene(key)
	if err != nil {
		return err
	}

	if g.currentScene != nil {
		g.currentScene.OnExit()
	}
	g.stack = append(g.stack, g.currentScene)
	g.currentScene = newScene
	g.currentScene.OnEnter()
	return nil
}

func (g *Game) PopScene() {
	if len(g.stack) == 0 {
		g.log.Fatalln("no scene to pop")
		return
	}

	g.currentScene.OnExit()
	idx := len(g.stack) - 1
	g.currentScene = g.stack[idx]
	g.stack = g.stack[:idx]
	g.currentScene.OnEnter()
}

func (g *Game) CurrentScene() gameapi.IGameScene {
	return g.currentScene
}

func (g *Game) CreateSceneFromDef(def *models.SceneDefinition) (gameapi.IGameScene, error) {
	g.log.Println(fmt.Sprintf("Creating scene from def: %v", def))
	if def == nil {
		sDef := &models.SceneDefinition{Menu: models.MenuOptions{
			MenuTitle:   "Scene Definition was nil",
			MenuOptions: nil,
		}}
		return scenes.NewMenuScene(g, sDef, g.log), nil
	}
	switch def.Type {
	case models.SceneTypeMenu:
		return scenes.NewMenuScene(g, def, g.log), nil
	case models.SceneTypeCutScene:
		return scenes.NewCutScene(g, def, g.log), nil
	case models.SceneTypeWorld:
		return scenes.NewWorldScene(g, def, g.log), nil
	default:
		sDef := &models.SceneDefinition{Menu: models.MenuOptions{
			MenuTitle:   "Unknown Scene Def Type",
			MenuOptions: nil,
		}}
		return scenes.NewMenuScene(g, sDef, g.log), nil
	}
}

func (g *Game) GetGameData() *models.GameData {
	return g.data
}

func (g *Game) SetGameData(gameData *models.GameData) {
	g.data = gameData
}

func (g *Game) GetInputHandler() gameapi.IInputHandler {
	return g.input
}

func (g *Game) GetRenderer() gameapi.Renderer {
	return g.renderer
}

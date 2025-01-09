package game

import (
	"relichunters/internal/character"
	"relichunters/internal/gameapi"
	"relichunters/internal/models"
	"relichunters/internal/player"
	"relichunters/internal/scenes"
)

type Game struct {
	//persistent data
	data         *models.GameData
	input        gameapi.IInputHandler
	renderer     gameapi.Renderer
	sceneManager gameapi.ISceneManager
}

func NewGame(ih gameapi.IInputHandler, renderer gameapi.Renderer, sm gameapi.ISceneManager) *Game {
	party := []*character.GameCharacter{&character.GameCharacter{
		Name:     "MainCharacter",
		Hp:       100,
		MaxHp:    100,
		Speed:    5,
		MaxSpeed: 5,
	}}
	return &Game{data: &models.GameData{Player: &player.Player{Party: party}}, input: ih, renderer: renderer, sceneManager: sm}
}

func (g *Game) GetData() *models.GameData {
	return g.data
}

func (g *Game) GetInputHandler() gameapi.IInputHandler {
	return g.input
}

func (g *Game) GetRenderer() gameapi.Renderer {
	return g.renderer
}

func (g *Game) GetSceneManager() gameapi.ISceneManager {
	return g.sceneManager
}

func (g *Game) CreateMenuScene() gameapi.IGameScene {
	return scenes.NewMenuScene(g)
}

func (g *Game) CreateWorldScene() gameapi.IGameScene {
	return scenes.NewWorldScene(g)
}

func (g *Game) CreateBattleScene() gameapi.IGameScene {
	//TODO implement me
	panic("implement me")
}

func (g *Game) CreateCutScene() gameapi.IGameScene {
	return scenes.NewCutScene(g)
}

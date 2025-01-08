package gameapi

import "relichunters/internal/models"

type ISceneManager interface {
	CurrentScene() IGameScene
	SetScene(newState IGameScene)
	PushScene(newState IGameScene)
	PopScene()
}

type IGameScene interface {
	Update(delta float64)
	Render(delta float64)
	HandleInput(cmd GameCommand)
	OnEnter()
	OnExit()
}

type IGameApi interface {
	GetData() *models.GameData
	CreateMenuScene() IGameScene
	CreateWorldScene() IGameScene
	CreateBattleScene() IGameScene
	CreateCutScene() IGameScene
}

type IInputHandler interface {
	PollCommands(currentScene IGameScene) []GameCommand
}

type GameCommand interface {
	Execute()
}

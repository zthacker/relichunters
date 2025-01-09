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
	HandleInput()
	OnEnter()
	OnExit()
}

type IGameApi interface {
	GetData() *models.GameData
	GetInputHandler() IInputHandler
	GetRenderer() Renderer
	GetSceneManager() ISceneManager
	CreateMenuScene() IGameScene
	CreateWorldScene() IGameScene
	CreateBattleScene() IGameScene
	CreateCutScene() IGameScene
}

type IInputHandler interface {
	PollCommands() []GameCommand
}

// Renderer will have a TODO to implement other methods as the game grows
type Renderer interface {
	Init() error
	Clear() error
	DrawText(x, y int, text string) error
	DrawTextStyled(x, y int, text string, style *models.Style) error
	DrawBox(x, y, w, h int, style *models.Style) error
	DrawLine(x1, y1, x2, y2 int, style *models.Style) error
	DrawSprites() error
	DrawImage(img string, x, y int) error
	GetSize() (int, int)
	Present() error
	Stop() error
}

type GameCommand interface{}

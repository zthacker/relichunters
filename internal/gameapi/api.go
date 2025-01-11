package gameapi

import (
	"relichunters/internal/models"
)

type ISceneManager interface {
	CreateScene(key string) (IGameScene, error)
	SetScene(key string) error
	PushScene(key string) error
	PopScene()
	CurrentScene() IGameScene
}

type IGameScene interface {
	Update(delta float64)
	Render(delta float64)
	HandleInput()
	OnEnter()
	OnExit()
}

type IGameApi interface {
	ISceneManager
	GetGameData() *models.GameData
	SetGameData(data *models.GameData)
	GetInputHandler() IInputHandler
	GetRenderer() Renderer
}

type IInputHandler interface {
	PollCommands() []GameCommand
}

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

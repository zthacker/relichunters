package scenes

import (
	"relichunters/internal/gameapi"
	"relichunters/internal/models"
)

type WorldScene struct {
	game         gameapi.IGameApi
	gameData     *models.GameData
	renderer     gameapi.Renderer
	inputHandler gameapi.IInputHandler
}

func NewWorldScene(g gameapi.IGameApi) *WorldScene {
	return &WorldScene{game: g}
}

func (w *WorldScene) Update(delta float64) {
	//check for random battle here
	//do AIStrategies -- patrol, etc
	//
}

func (w *WorldScene) Render(delta float64) {
	w.renderer.Clear()
	w.renderer.DrawText(w.gameData.Player.Party[0].X, w.gameData.Player.Party[0].Y, "@")

	w.renderer.Present()
}

func (w *WorldScene) HandleInput() {
	//handle inputs
}

func (w *WorldScene) OnEnter() {
	w.gameData = w.game.GetData()
	w.renderer = w.game.GetRenderer()
	w.inputHandler = w.game.GetInputHandler()
}

func (w *WorldScene) OnExit() {
	//TODO implement me
	panic("implement me")
}

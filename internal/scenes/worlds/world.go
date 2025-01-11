package worlds

import (
	"relichunters/internal/gameapi"
	"relichunters/internal/models"
)

type WorldOneScene struct {
	game         gameapi.IGameApi
	gameData     *models.GameData
	renderer     gameapi.Renderer
	inputHandler gameapi.IInputHandler
}

func NewWorldOneScene(g gameapi.IGameApi) *WorldOneScene {
	return &WorldOneScene{game: g}
}

func (w *WorldOneScene) Update(delta float64) {
	//check for random battle here
	//do AIStrategies -- patrol, etc
	//
}

func (w *WorldOneScene) Render(delta float64) {
	w.renderer.Clear()
	w.renderer.DrawText(w.gameData.Player.Party[0].X, w.gameData.Player.Party[0].Y, "@")

	w.renderer.Present()
}

func (w *WorldOneScene) HandleInput() {
	//handle inputs
}

func (w *WorldOneScene) OnEnter() {
	w.gameData = w.game.GetGameData()
	w.renderer = w.game.GetRenderer()
	w.inputHandler = w.game.GetInputHandler()
}

func (w *WorldOneScene) OnExit() {
	//TODO implement me
	panic("implement me")
}

package scenes

import (
	"fmt"
	"log"
	"relichunters/internal/gameapi"
	"relichunters/internal/models"
)

type WorldScene struct {
	game         gameapi.IGameApi
	sDefs        *models.SceneDefinition
	log          *log.Logger
	gameData     *models.GameData
	renderer     gameapi.Renderer
	inputHandler gameapi.IInputHandler
}

func NewWorldScene(g gameapi.IGameApi, sceneDefs *models.SceneDefinition, logger *log.Logger) *WorldScene {
	return &WorldScene{game: g, sDefs: sceneDefs, log: logger}
}

func (w *WorldScene) Update(delta float64) {
	//check for random battle here
	//do AIStrategies -- patrol, etc
	//
}

func (w *WorldScene) Render(delta float64) {
	w.renderer.Clear()
	w.renderer.DrawText(0, 0, w.sDefs.Description)
	w.renderer.DrawText(w.gameData.Player.Party[0].X+1, w.gameData.Player.Party[0].Y+1, "@")

	w.renderer.Present()
}

func (w *WorldScene) HandleInput() {
	//handle inputs
}

func (w *WorldScene) OnEnter() {
	w.log.Println(fmt.Sprintf("Entering into WorldScene: %s", w.sDefs.Description))
	w.gameData = w.game.GetGameData()
	w.renderer = w.game.GetRenderer()
	w.inputHandler = w.game.GetInputHandler()
}

func (w *WorldScene) OnExit() {
	//do any OnExit() stuff
}

package scenes

import (
	"fmt"
	"log"
	"relichunters/internal/gameapi"
	"relichunters/internal/models"
	"relichunters/internal/player"
)

type CutScene struct {
	game     gameapi.IGameApi
	sDefs    *models.SceneDefinition
	log      *log.Logger
	input    gameapi.IInputHandler
	renderer gameapi.Renderer
	timer    float64
}

func NewCutScene(g gameapi.IGameApi, sceneDefs *models.SceneDefinition, logger *log.Logger) *CutScene {
	return &CutScene{game: g, sDefs: sceneDefs, log: logger}
}

func (n *CutScene) Update(delta float64) {
	//for  now just ticket down the timer
	if n.timer > 0 {
		n.timer -= delta
	}
	if n.timer < 0 {
		//TODO for cutscenes, we can have this in the SceneDefinition on where to go to next
		n.game.SetScene(models.SceneKeyNewGameWorld)
	}
}

func (n *CutScene) Render(delta float64) {
	n.renderer.Clear()
	txt := fmt.Sprintf("Cutscene Length: %f ", n.timer)
	n.renderer.DrawText(0, 0, txt)
	n.renderer.Present()
}

func (n *CutScene) HandleInput() {
	//nothing to do here
}

func (n *CutScene) OnEnter() {
	n.log.Println(fmt.Sprintf("Entered CutScene: %s", n.sDefs.Description))
	n.timer = 2.0
	n.renderer = n.game.GetRenderer()

	//NewGame setup
	newGameData := n.game.GetGameData()
	newGameData.Player = player.NewPlayer()
	n.log.Println(fmt.Sprintf("CutScene: %+v", n))
}

func (n *CutScene) OnExit() {
	//nothing here
}

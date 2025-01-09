package scenes

import (
	"fmt"
	"relichunters/internal/gameapi"
)

type CutScene struct {
	game     gameapi.IGameApi
	renderer gameapi.Renderer
	//later on load in things like images, videos, etc
	//for now we'll just set a timer on a cutscene and when it's done we'll transition
	timer float64
}

func NewCutScene(g gameapi.IGameApi) *CutScene {
	return &CutScene{game: g}
}

func (c *CutScene) Update(delta float64) {
	//for  now just ticket down the timer
	if c.timer > 0 {
		c.timer -= delta
	}
	if c.timer < 0 {
		sm := c.game.GetSceneManager()
		world := c.game.CreateWorldScene()
		sm.SetScene(world)
	}
}

func (c *CutScene) Render(delta float64) {
	c.renderer.Clear()
	txt := fmt.Sprintf("Cutscene Length: %f ", c.timer)
	c.renderer.DrawText(0, 0, txt)
	c.renderer.Present()
}

func (c *CutScene) HandleInput() {
	//nothing to handle
}

func (c *CutScene) OnEnter() {
	c.timer = 2.0
	c.renderer = c.game.GetRenderer()
}

func (c *CutScene) OnExit() {

}

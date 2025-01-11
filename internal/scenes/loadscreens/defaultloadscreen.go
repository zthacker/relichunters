package loadscreens

import (
	"fmt"
	"relichunters/internal/gameapi"
)

type DefaultLoadScreen struct {
	game     gameapi.IGameApi
	renderer gameapi.Renderer
	timer    float64
}

func NewDefaultLoadScreen(g gameapi.IGameApi) *DefaultLoadScreen {
	return &DefaultLoadScreen{game: g}
}

func (c *DefaultLoadScreen) Update(delta float64) {
	//for  now just ticket down the timer
	if c.timer > 0 {
		c.timer -= delta
	}
}

func (c *DefaultLoadScreen) Render(delta float64) {
	c.renderer.Clear()
	if c.timer < 0 {
		txt := fmt.Sprintf("Cutscene Done -- Implement next scene")
		c.renderer.DrawText(0, 0, txt)
	} else {
		txt := fmt.Sprintf("Cutscene Length: %f ", c.timer)
		c.renderer.DrawText(0, 0, txt)
	}

	c.renderer.Present()
}

func (c *DefaultLoadScreen) HandleInput() {
	//nothing to handle
}

func (c *DefaultLoadScreen) OnEnter() {
	c.timer = 2.0
	c.renderer = c.game.GetRenderer()
}

func (c *DefaultLoadScreen) OnExit() {

}

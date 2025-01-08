package scenes

import (
	"fmt"
	"relichunters/internal/gameapi"
)

type MenuScene struct {
	game     gameapi.IGameApi
	timer    float64
	maxTimer float64
}

func NewMenuScene(g gameapi.IGameApi) *MenuScene {
	return &MenuScene{g, 0, 0}
}

func (m *MenuScene) Update(delta float64) {
	if m.timer <= 0 {
		m.timer = m.maxTimer
	}
	m.timer -= delta
}

func (m *MenuScene) Render(delta float64) {
	fmt.Printf("%f / %f\n", m.timer, m.maxTimer)
}

func (m *MenuScene) HandleInput(gameapi.GameCommand) {
	//would handle inputs here
	//we'd use game.CreateSomeScene from the things we select
}

func (m *MenuScene) OnEnter() {
	m.maxTimer = 10.0
	m.timer = m.maxTimer
}

func (m *MenuScene) OnExit() {
	fmt.Println("leaving MenuScene")
}

package scenes

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"relichunters/internal/commands"
	"relichunters/internal/gameapi"
	"relichunters/internal/models"
)

type MenuScene struct {
	game     gameapi.IGameApi
	input    gameapi.IInputHandler
	renderer gameapi.Renderer
	options  []string
}

func NewMenuScene(g gameapi.IGameApi) *MenuScene {
	return &MenuScene{g, nil, nil, nil}
}

func (m *MenuScene) Update(delta float64) {

}

func (m *MenuScene) Render(delta float64) {
	//TODO catch errors later
	m.renderer.Clear()

	screenW, screenH := m.renderer.GetSize()

	menuWidth := 20
	menuHeight := len(m.options) + 2

	startX := (screenW - menuWidth) / 2
	startY := (screenH - menuHeight) / 2

	boxStyle := &models.Style{
		ForegroundColor: uint64(tcell.ColorYellow),
		BackgroundColor: uint64(tcell.ColorBlack),
		Bold:            false,
	}
	textStyle := &models.Style{
		ForegroundColor: uint64(tcell.ColorPurple),
		BackgroundColor: uint64(tcell.ColorBlack),
		Bold:            false,
	}
	m.renderer.DrawBox(startX, startY, menuWidth, menuHeight, boxStyle)

	for i, option := range m.options {
		lineX := startX + 2
		lineY := startY + 1 + i

		m.renderer.DrawTextStyled(lineX, lineY, option, textStyle)
	}

	m.renderer.Present()
}

func (m *MenuScene) HandleInput() {
	cmds := m.input.PollCommands()
	for _, cmd := range cmds {
		switch cmd.(type) {
		case commands.PauseCommand:

		}
	}

}

func (m *MenuScene) OnEnter() {
	m.options = append(m.options, "New Game", "Load Game", "Save Game", "Quit Game")
	m.input = m.game.GetInputHandler()
	m.renderer = m.game.GetRenderer()
}

func (m *MenuScene) OnExit() {
	fmt.Println("leaving MenuScene")
}

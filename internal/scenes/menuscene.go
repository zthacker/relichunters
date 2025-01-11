package scenes

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"log"
	"relichunters/internal/commands"
	"relichunters/internal/gameapi"
	"relichunters/internal/models"
)

type MenuScene struct {
	game        gameapi.IGameApi
	sDefs       *models.SceneDefinition
	log         *log.Logger
	input       gameapi.IInputHandler
	renderer    gameapi.Renderer
	cursorIndex int
}

func NewMenuScene(g gameapi.IGameApi, sceneDefs *models.SceneDefinition, logger *log.Logger) *MenuScene {
	return &MenuScene{game: g, sDefs: sceneDefs, log: logger, cursorIndex: 0}
}

func (m *MenuScene) Update(delta float64) {

}

func (m *MenuScene) Render(delta float64) {

	m.renderer.Clear()

	screenW, screenH := m.renderer.GetSize()

	menuWidth := 20
	menuHeight := len(m.sDefs.Menu.MenuOptions) + 2

	startX := (screenW - menuWidth) / 2
	startY := (screenH - menuHeight) / 2

	m.renderer.DrawText(1, 1, m.sDefs.Menu.MenuTitle)

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

	for i, option := range m.sDefs.Menu.MenuOptions {
		lineX := startX + 2
		lineY := startY + 1 + i

		if m.cursorIndex == i {
			txt := fmt.Sprintf("-> %s", option.Display)
			m.renderer.DrawTextStyled(lineX, lineY, txt, textStyle)
			continue
		}
		m.renderer.DrawTextStyled(lineX, lineY, option.Display, textStyle)
	}

	m.renderer.Present()
}

func (m *MenuScene) HandleInput() {
	cmds := m.input.PollCommands()
	for _, cmd := range cmds {
		switch c := cmd.(type) {
		case commands.MoveCommand:
			if c.Dy > 0 {
				m.cursorIndex++
				if m.cursorIndex >= len(m.sDefs.Menu.MenuOptions) {
					m.cursorIndex = 0
				}
			}
			if c.Dy < 0 {
				m.cursorIndex--
				if m.cursorIndex < 0 {
					m.cursorIndex = len(m.sDefs.Menu.MenuOptions) - 1
				}
			}
		case commands.SelectCommand:
			opt := m.sDefs.Menu.MenuOptions[m.cursorIndex]
			if err := m.game.SetScene(opt.Key); err != nil {
				m.log.Fatalln(err)
			}
		}
	}

}

func (m *MenuScene) OnEnter() {
	m.log.Println("Entering MenuScene")
	m.input = m.game.GetInputHandler()
	m.renderer = m.game.GetRenderer()
	m.log.Println(fmt.Sprintf("MenuScene: %+v", m))
}

func (m *MenuScene) OnExit() {

}

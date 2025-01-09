package scenes

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"relichunters/internal/commands"
	"relichunters/internal/gameapi"
	"relichunters/internal/models"
)

type MenuSelection string

const (
	NewGame  MenuSelection = "New Game"
	LoadGame MenuSelection = "Load Game"
	SaveGame MenuSelection = "Save Game"
	QuitGame MenuSelection = "Quit"
)

type MenuScene struct {
	game        gameapi.IGameApi
	input       gameapi.IInputHandler
	renderer    gameapi.Renderer
	options     []MenuSelection
	cursorIndex int
}

func NewMenuScene(g gameapi.IGameApi) *MenuScene {
	return &MenuScene{game: g, cursorIndex: 0}
}

func (m *MenuScene) Update(delta float64) {

}

func (m *MenuScene) Render(delta float64) {

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

		if m.cursorIndex == i {
			txt := fmt.Sprintf("-> %s", string(option))
			m.renderer.DrawTextStyled(lineX, lineY, txt, textStyle)
			continue
		}
		m.renderer.DrawTextStyled(lineX, lineY, string(option), textStyle)
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
				if m.cursorIndex >= len(m.options) {
					m.cursorIndex = 0
				}
			}
			if c.Dy < 0 {
				m.cursorIndex--
				if m.cursorIndex < 0 {
					m.cursorIndex = len(m.options) - 1
				}
			}
		case commands.SelectCommand:
			opt := m.options[m.cursorIndex]
			switch opt {
			case NewGame:
				introCutScene := m.game.CreateCutScene()
				sm := m.game.GetSceneManager()
				sm.SetScene(introCutScene)
			case LoadGame:
				//goes to a load game scene to load data
			case SaveGame:
				//saves game data
			case QuitGame:
			}
		}
	}

}

func (m *MenuScene) OnEnter() {
	m.options = append(m.options, NewGame, LoadGame, SaveGame, QuitGame)
	m.input = m.game.GetInputHandler()
	m.renderer = m.game.GetRenderer()
}

func (m *MenuScene) OnExit() {

}

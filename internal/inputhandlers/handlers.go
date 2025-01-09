package inputhandlers

import (
	"github.com/gdamore/tcell/v2"
	"relichunters/internal/commands"
	"relichunters/internal/gameapi"
)

type TCellInputHandler struct {
	screen tcell.Screen
}

func NewTCellInputHandler(screen tcell.Screen) *TCellInputHandler {
	return &TCellInputHandler{screen: screen}
}

func (t *TCellInputHandler) PollCommands() []gameapi.GameCommand {
	var cmds []gameapi.GameCommand

	for t.screen.HasPendingEvent() {
		ev := t.screen.PollEvent()
		if ke, ok := ev.(*tcell.EventKey); ok {
			switch ke.Key() {
			case tcell.KeyUp:
				cmds = append(cmds, commands.MoveCommand{Dy: -1})
			case tcell.KeyDown:
				cmds = append(cmds, commands.MoveCommand{Dy: 1})
			case tcell.KeyLeft:
				cmds = append(cmds, commands.MoveCommand{Dx: -1})
			case tcell.KeyRight:
				cmds = append(cmds, commands.MoveCommand{Dx: 1})
			case tcell.KeyEnter:
				cmds = append(cmds, commands.SelectCommand{})
			case tcell.KeyEscape:
				cmds = append(cmds, commands.PauseCommand{})
			case tcell.KeyCtrlX:
				cmds = append(cmds, commands.CancelCommand{})
			case tcell.KeyCtrlZ:
				cmds = append(cmds, commands.BackCommand{})
			}
		}
	}
	return cmds
}

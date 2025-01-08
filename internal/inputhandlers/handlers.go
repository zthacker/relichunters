package inputhandlers

import (
	"relichunters/internal/gameapi"
	"relichunters/internal/scenes"
)

type TCellInputHandler struct {
	screen tcell.Screen
}

func (T TCellInputHandler) PollCommands(currentScene gameapi.IGameScene) []gameapi.GameCommand {
	switch currentScene.(type) {
	case *scenes.MenuScene:

	}
}

package scenemanager

import (
	"fmt"
	"relichunters/internal/gameapi"
)

type DefaultSceneManager struct {
	current gameapi.IGameScene
	stack   []gameapi.IGameScene
}

func NewDefaultSceneManager(initialScene gameapi.IGameScene) *DefaultSceneManager {
	mgr := &DefaultSceneManager{
		current: initialScene,
	}

	mgr.current.OnEnter()
	return mgr
}

func (d *DefaultSceneManager) CurrentScene() gameapi.IGameScene {
	return d.current
}

// SetScene calls OnExit of the current scene and sets current scene to new scene passed and calls OnEnter()
func (d *DefaultSceneManager) SetScene(newState gameapi.IGameScene) {
	if d.current != nil {
		d.current.OnExit()
	}
	d.current = newState
	d.current.OnEnter()
}

// PushScene pushes a scene to the top of the stack; calls OnExit() of current and sets current to new and calls OnEnter()
func (d *DefaultSceneManager) PushScene(newState gameapi.IGameScene) {
	//pushing to back of stack
	d.stack = append(d.stack, newState)
	d.current.OnExit()
	d.current = newState
	d.current.OnEnter()
}

// PopScene removes the top scene from the stack
func (d *DefaultSceneManager) PopScene() {
	if len(d.stack) == 0 {
		fmt.Println("Scene Stack is empty")
		return
	}
	d.current.OnExit()

	//get the back of the stack since we push to back
	idx := len(d.stack) - 1

	//set current
	d.current = d.stack[idx]

	//remove back of stack
	d.stack = d.stack[:idx]

	//do OnEnter() from current
	d.current.OnEnter()

}

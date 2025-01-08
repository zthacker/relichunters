package engine

// IGameState represents the interface for GameStates to implement
type IGameState interface {
	Update(delta float64)
	Render(delta float64)
	HandleInput()
	OnEnter()
	OnExit()
}

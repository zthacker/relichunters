package commands

type BackCommand struct{}
type MoveCommand struct {
	Dx, Dy int
}
type SelectCommand struct{}
type CancelCommand struct{}
type PauseCommand struct{}

package player

import "relichunters/internal/character"

type Player struct {
	Party []*character.GameCharacter
}

func NewPlayer() *Player {
	party := []*character.GameCharacter{&character.GameCharacter{
		Name:     "Ghost",
		Hp:       100,
		MaxHp:    100,
		Speed:    5,
		MaxSpeed: 5,
		X:        0,
		Y:        0,
	}}
	return &Player{Party: party}
}

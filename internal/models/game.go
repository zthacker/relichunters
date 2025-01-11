package models

import "relichunters/internal/player"

type GameData struct {
	Player           *player.Player
	CurrentGameScene string
}

// SceneDefinition holds data about how to build a scene
// You can expand this with fields for dialogues, music, cutscene scripts, etc.
type SceneDefinition struct {
	Key       string    `yaml:"key"`
	Type      SceneType `yaml:"type"`
	NextScene string    `yaml:"next_scene"`

	//General
	Description string `yaml:"description"`
	IsOverlay   bool   `yaml:"is_overlay"`

	// Menu
	Menu MenuOptions `yaml:"menu"`

	// Map
	MapID string `yaml:"map_id,omitempty"`

	//Same thing as Enemies, need a way to look up NPCs to create
	NPCs []string `yaml:"npcs,omitempty"`

	// Battle -- lookups for factory -- expand to a Enemy struct
	//for more spawning
	EnemyNames []string `yaml:"enemy_names,omitempty"`
	//Player Participants
	//The idea here for "special" battles that you only want to load certain
	//player controlled characters instead of the whole party

	//Dialogue loading
	//Dialogue table look up
	//

}

type MenuOptions struct {
	MenuTitle string           `yaml:"menu_title"`
	Options   []MenuSelections `yaml:"options"`
}

type MenuSelections struct {
	SceneKey string `yaml:"scene_key"`
	Display  string `yaml:"display"`
}

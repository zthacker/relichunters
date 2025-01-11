package models

import "relichunters/internal/player"

type GameData struct {
	Player           *player.Player
	CurrentGameScene SceneKey
}

// SceneDefinition holds data about how to build a scene
// You can expand this with fields for dialogues, music, cutscene scripts, etc.
type SceneDefinition struct {
	Key  SceneKey  `json:"key"`
	Type SceneType `json:"type"`

	//General
	Description string `json:"description"`
	IsOverlay   bool   `json:"is_overlay"`

	// Menu
	Menu MenuOptions `json:"menu"`

	// Map
	MapID string `json:"map_id,omitempty"`

	//Same thing as Enemies, need a way to look up NPCs to create
	NPCs []string `json:"npcs,omitempty"`

	// Battle -- lookups for factory -- expand to a Enemy struct
	//for more spawning
	EnemyNames []string `json:"enemy_names,omitempty"`
	//Player Participants
	//The idea here for "special" battles that you only want to load certain
	//player controlled characters instead of the whole party

	//Dialogue loading
	//Dialogue table look up
	//

}

type MenuOptions struct {
	MenuTitle   string           `json:"menu_title"`
	MenuOptions []MenuSelections `json:"display_options"`
}

type MenuSelections struct {
	Key     SceneKey `json:"key"`
	Display string   `json:"display"`
}

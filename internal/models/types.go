package models

type SceneKey string // definitions in according files
type SceneType string

const (
	SceneTypeMenu      SceneType = "MENU"
	SceneTypeOverworld SceneType = "OVERWORLD"
	SceneTypeBattle    SceneType = "BATTLE"
	SceneTypeInventory SceneType = "INVENTORY"
	SceneTypeCutScene  SceneType = "CUT_SCENE"
	SceneTypeLoadScene SceneType = "LOAD_SCENE"
)

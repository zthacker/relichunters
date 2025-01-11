package sceneregistry

import (
	"fmt"
	"relichunters/internal/models"
)

type SceneRegistry struct {
	defs map[models.SceneKey]*models.SceneDefinition
}

func NewSceneRegistry() *SceneRegistry {
	return &SceneRegistry{
		defs: make(map[models.SceneKey]*models.SceneDefinition),
	}
}

// Register a definition. In real code, you'd parse JSON or YAML in a loop.
func (r *SceneRegistry) LoadSceneDef(def models.SceneDefinition) {
	r.defs[def.Key] = &def
}

func (r *SceneRegistry) GetDefinition(key models.SceneKey) (*models.SceneDefinition, error) {
	def, ok := r.defs[key]
	if !ok {
		return nil, fmt.Errorf("no scene definition for key: %s", key)
	}
	return def, nil
}

func (r *SceneRegistry) SaveSceneDef(def *models.SceneDefinition) {
	r.defs[def.Key] = def
}

package sceneregistry

import (
	"fmt"
	"relichunters/internal/models"
)

type SceneRegistry struct {
	defs map[string]*models.SceneDefinition
}

func NewSceneRegistry() *SceneRegistry {
	return &SceneRegistry{
		defs: make(map[string]*models.SceneDefinition),
	}
}

func (r *SceneRegistry) SetSceneDef(def *models.SceneDefinition) {
	r.defs[def.Key] = def
}

func (r *SceneRegistry) GetDefinition(key string) (*models.SceneDefinition, error) {
	def, ok := r.defs[key]
	if !ok {
		return nil, fmt.Errorf("no scene definition for key: %s", key)
	}
	return def, nil
}

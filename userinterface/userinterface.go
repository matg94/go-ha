package userinterface

import "autobubble/api"

func InitialModel(api api.HaApi) model {
	allEntities := api.GetEntityStates()
	var categories []string
	for _, e := range allEntities {
		var exists = false
		for _, cat := range categories {
			if cat == e.Category {
				exists = true
				break
			}
		}
		if !exists {
			categories = append(categories, e.Category)
		}
	}
	return model{
		categories:     categories,
		selected:       false,
		categoryCursor: 0,
		entityCursor:   0,
		entities:       allEntities,
	}
}

func (m *model) UpdateEntities(updatedEntities []*api.Entity) {
	for _, entity := range m.entities {
		for _, updatedEntity := range updatedEntities {
			if entity.ID != updatedEntity.ID {
				continue
			}
			entity.State = updatedEntity.State
		}
	}
}

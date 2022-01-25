package userinterface

import "autobubble/api"

func InitialModel(allEntities []*api.Entity) model {
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

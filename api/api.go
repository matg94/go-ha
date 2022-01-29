package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

type HaApi struct {
	HAUrl   string
	HAToken string
}

func (h HaApi) GetAPIState() string {
	req := HARequest{
		url:     fmt.Sprintf("http://%s/api/", h.HAUrl),
		token:   h.HAToken,
		reqtype: "GET",
	}
	return req.send()
}

func (h HaApi) GetEntityStates() []*Entity {
	var entities []*HAEntity
	req := HARequest{
		url:     fmt.Sprintf("http://%s/api/states", h.HAUrl),
		token:   h.HAToken,
		reqtype: "GET",
	}

	entitiesJson := req.send()
	json.Unmarshal([]byte(entitiesJson), &entities)

	var convertedEntities []*Entity
	for _, e := range entities {
		convertedEntities = append(convertedEntities, &Entity{
			ID:       strings.Split(e.ID, ".")[1],
			State:    e.State,
			Category: strings.Split(e.ID, ".")[0],
		})
	}

	return convertedEntities
}

func (h HaApi) ActivateService(category string, serviceName string, entityId string) []*Entity {
	var changedEntities []*HAEntity
	req := HARequest{
		url:     fmt.Sprintf("http://%s/api/services/%s/%s", h.HAUrl, category, serviceName),
		token:   h.HAToken,
		body:    map[string]string{"entity_id": fmt.Sprintf("%s.%s", category, entityId)},
		reqtype: "POST",
	}

	entitiesJson := req.send()
	json.Unmarshal([]byte(entitiesJson), &changedEntities)

	var convertedEntities []*Entity
	for _, e := range changedEntities {
		convertedEntities = append(convertedEntities, &Entity{
			ID:       strings.Split(e.ID, ".")[1],
			State:    e.State,
			Category: strings.Split(e.ID, ".")[0],
		})
	}

	return convertedEntities
}

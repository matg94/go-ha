package api

type Entity struct {
	id       string
	state    string
	category string
}

type HAEntity struct {
	ID    string `json:"entity_id"`
	State string `json:"state"`
}

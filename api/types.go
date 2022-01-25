package api

type Entity struct {
	ID       string
	State    string
	Category string
}

type HAEntity struct {
	ID    string `json:"entity_id"`
	State string `json:"state"`
}

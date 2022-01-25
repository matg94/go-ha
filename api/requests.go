package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type HARequest struct {
	token   string
	url     string
	reqtype string
	body    map[string]string
}

func (req *HARequest) send() string {
	body, err := json.Marshal(req.body)

	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	request, err := http.NewRequest(req.reqtype, req.url, bytes.NewBuffer(body))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", req.token))

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(res)
}

func GetAPIState() (string, error) {
	req := HARequest{
		url:     fmt.Sprintf("http://%s/api/", os.Getenv("HA_URL")),
		token:   os.Getenv("HA_TOKEN"),
		reqtype: "GET",
	}
	return req.send(), nil
}

func GetEntityStates() ([]*Entity, error) {
	var entities []*HAEntity
	req := HARequest{
		url:     fmt.Sprintf("http://%s/api/states", os.Getenv("HA_URL")),
		token:   os.Getenv("HA_TOKEN"),
		reqtype: "GET",
	}

	entitiesJson := req.send()
	json.Unmarshal([]byte(entitiesJson), &entities)

	var convertedEntities []*Entity
	for _, e := range entities {
		convertedEntities = append(convertedEntities, &Entity{
			id:       strings.Split(e.ID, ".")[1],
			state:    e.State,
			category: strings.Split(e.ID, ".")[0],
		})
	}
	for _, e := range convertedEntities {
		fmt.Println(e)
	}
	return convertedEntities, nil
}

func GetEntityState(entity Entity) (Entity, error) {

	return Entity{}, nil
}
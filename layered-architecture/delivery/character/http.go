// layered-architecture/delivery/character/http.go

package character

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"layered-architecture/datastore"
	"layered-architecture/entities"
	"net/http"
	"strconv"
)

type CharacterHandler struct {
	datastore datastore.Character
}

func New(character datastore.Character) CharacterHandler {
	return CharacterHandler{datastore: character}
}

func (c CharacterHandler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		c.Get(w, r)
	case http.MethodPost:
		c.Create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Get() makes a GET request and returns a specific character
func (c CharacterHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	cid, err := strconv.Atoi(id)
	if err != nil {
		_, _ = w.Write([]byte("invalid parameter id"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	resp, err := c.datastore.Get(cid)
	if err != nil {
		_, _ = w.Write([]byte("could not retrieve character"))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	body, _ := json.Marshal(resp)
	_, _ = w.Write(body)
}

// Create() makes a POST request and returns a newly created character
func (c CharacterHandler) Create(w http.ResponseWriter, r *http.Request) {
	var character entities.Character

	body, _ := ioutil.ReadAll(r.Body)
	
	err := json.Unmarshal(body, &character)
	if err != nil {
		fmt.Println(err)
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	resp, err := c.datastore.Create(character)
	if err != nil {
		_, _ = w.Write([]byte("could not create character"))
		w.WriteHeader(http.StatusInternalServerError)

		return 
	}

	body, _ = json.Marshal(resp)
	_, _ = w.Write(body)
}
// layered-architecture/use-case/character/logic.go

package character

import (
	"layered-architecture/datastore"
	"layered-architecture/entities"
)

type CharacterLogic struct {
	datastore datastore.Character
}

func New(character datastore.Character) CharacterLogic {
	return CharacterLogic{datastore: character}
}

// Get() returns a character
func (c CharacterLogic) Get(cid int) ([]entities.Character, error) {
	resp, err := c.datastore.Get(cid)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Create() creates a new character
func (c CharacterLogic) Create(character entities.Character) (entities.Character, error) {
	resp, err := c.datastore.Create(character)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

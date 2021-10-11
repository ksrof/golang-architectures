// layered-gin/usecase/character/character.go

package character

import (
	"layered-gin/entity"
	"layered-gin/repository"
	"log"
)

type CharacterUse struct {
	repo repository.Character
}

func New(character repository.Character) CharacterUse {
	return CharacterUse{repo: character}
}

// GetAll() returns all characters
func (c CharacterUse) GetAll(characters []entity.Character) ([]entity.Character, error) {
	resp, err := c.repo.GetAll(characters)
	if err != nil {
		log.Printf("could not get characters")
		return resp, err
	}

	return resp, nil
}

// GetByID() returns a specific character
func (c CharacterUse) GetByID(id uint) (entity.Character, error) {
	resp, err := c.repo.GetByID(id)
	if err != nil {
		log.Printf("could not get character")
		return resp, err
	}

	return resp, nil
}

// Create() creates a character
func (c CharacterUse) Create(character entity.Character) (entity.Character, error) {
	resp, err := c.repo.Create(character)
	if err != nil {
		log.Printf("could create character")
		return resp, err
	}

	return resp, nil
}
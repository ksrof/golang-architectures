// layered-gin/repository/mysql/mysql.go

package mysql

import (
	"layered-gin/entity"
	"log"

	"gorm.io/gorm"
)

type CharacterRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) CharacterRepo {
	return CharacterRepo{DB: db}
}

// GetAll() returns all characters
func (c CharacterRepo) GetAll(characters []entity.Character) ([]entity.Character, error) {
	err := c.DB.Model(entity.Character{}).Limit(100).Find(characters).Error
	if err != nil {
		log.Printf("unable to find characters")
		return []entity.Character{}, err
	}

	return characters, nil
}

// GetByID() returns a specific character
func (c CharacterRepo) GetByID(id uint) (entity.Character, error) {
	var character entity.Character

	err := c.DB.Model(entity.Character{}).Where("id = ?", id).Take(character).Error
	if err != nil {
		log.Printf("unable to find character")
		return entity.Character{}, err
	}

	return character, nil
}

// Create() creates a character
func (c CharacterRepo) Create(character entity.Character) (entity.Character, error) {
	err := c.DB.Create(character).Error
	if err != nil {
		log.Printf("unable to create character")
		return entity.Character{}, err
	}

	return character, nil
}
// layered-architecture/datastore/character/mysql.go

package character

import (
	"database/sql"
	"layered-architecture/entities"
)

type CharacterStorer struct {
	DB *sql.DB
}

func New(db *sql.DB) CharacterStorer {
	return CharacterStorer{DB: db}
}

// Get() returns specific character data
func (c CharacterStorer) Get(id int) ([]entities.Character, error) {
	var (
		rows *sql.Rows
		err error
	)

	if id != 0 {
		rows, err = c.DB.Query("SELECT * FROM characters where id = ?", id)
	} else {
		rows, err = c.DB.Query("SELECT * FROM characters")
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var characters []entities.Character

	for rows.Next() {
		var c entities.Character
		_ = rows.Scan(&c.ID, &c.Name, &c.Level, &c.Gender, &c.Race, &c.Class, &c.Faction)
		characters = append(characters, c)
	}

	return characters, nil
}

// Create() returns a newly created character data
func (c CharacterStorer) Create(character entities.Character) (entities.Character, error) {
	res, err := c.DB.Exec("INSERT INTO characters (name, level, gender, race, class, faction) VALUES(?,?,?,?,?,?)", character.Name, character.Level, character.Gender, character.Race, character.Class, character.Faction)
	if err != nil {
		return entities.Character{}, nil
	}

	id, _ := res.LastInsertId()
	character.ID = int(id)

	return character, nil
}
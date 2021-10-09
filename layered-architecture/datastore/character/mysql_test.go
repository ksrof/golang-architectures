// layered-architecture/datastore/character/mysql_test.go

package character

import (
	"database/sql"
	"layered-architecture/driver"
	"layered-architecture/entities"
	"os"
	"reflect"
	"testing"
)

// MySQLInit() initializes MySQL character database
func MySQLInit(t *testing.T) *sql.DB {
	conf := driver.MySQLConfig{
		Host: os.Getenv("SQL_HOST"),
		User: os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port: os.Getenv("SQL_PORT"),
		DB: os.Getenv("SQL_DB"),
	}

	var err error
	db, err := driver.MySQLConnect(conf)
	if err != nil {
		t.Errorf("could not connect to sql, err: %v", err)
	}

	return db
}

func TestDatastore(t *testing.T) {
	db := MySQLInit(t)
	c := New(db)
	TestCharacterStorer_Get(t, c)
	TestCharacterStorer_Create(t, c)
}

func TestCharacterStorer_Create(t *testing.T, db CharacterStorer) {
	testcases := []struct{
		request		entities.Character
		response 	entities.Character
	}{
		{entities.Character{Name: "Illidan Stormrage", Level: "73", Gender: "Male", Race: "Demon", Class: "Demon hunter", Faction: "Illidari"}, entities.Character{3, "Illidan Stormrage", "73", "Male", "Demon", "Demon hunter", "Illidari"}},
		{entities.Character{Name: "The Lich King", Level: "80", Gender: "Male", Race: "Undead", Class: "Necromancer", Faction: "Scourge"}, entities.Character{4, "The Lich King", "80", "Male", "Undead", "Necromancer", "Scourge"}},
	}

	for i, v := range testcases {
		resp, _ := db.Create(v.request)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.response)
		}
	}
}

func TestCharacterStorer_Get(t *testing.T, db CharacterStorer) {
	testcases := []struct{
		id int
		resp []entities.Character
	}{
		{0, []entities.Character{{1, "Leeroy Jenkins", "70", "Male", "Humanoid", "Paladin", "Alliance"}, {2, "Sylavanas Windrunner", "??(Boss)", "Female", "Forsaken", "Ranger", "Horde"}}},
		{1, []entities.Character{{1, "Leeroy Jenkins", "70", "Male", "Humanoid", "Paladin", "Alliance"}}},
	}

	for i, v := range testcases {
		resp, _ := db.Get(v.id)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.resp)
		}
	}
}
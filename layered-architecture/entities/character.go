// layered-architecture/entities/character.go

package entities

type Character struct {
	ID 			int 		`json:"id"`
	Name 		string 	`json:"name"`
	Level 	string 	`json:"level"`
	Gender 	string 	`json:"gender"`
	Race 		string 	`json:"race"`
	Class 	string 	`json:"class"`
	Faction string 	`json:"faction"`
}
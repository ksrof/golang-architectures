// layered-architecture/datastore/interface.go

package datastore

import "layered-architecture/entities"

type Character interface {
	Get(id int) ([]entities.Character, error)
	Create(entities.Character) (entities.Character, error)
}
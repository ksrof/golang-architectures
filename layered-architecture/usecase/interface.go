// layered-architecture/datastore/interface.go

package usecase

import "layered-architecture/entities"

type Character interface {
	Get(cid int) ([]entities.Character, error)
	Create(entities.Character) (entities.Character, error)
}
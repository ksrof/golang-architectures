// layered-gin/repository/repo.go

package repository

import "layered-gin/entity"

type Character interface {
	Get(id int) ([]entity.Character, error)
	Create(entity.Character) (entity.Character, error)
}
// layered-gin/usecase/use.go

package usecase

import "layered-gin/entity"

type Character interface {
	GetAll([]entity.Character) ([]entity.Character, error)
	GetByID(id uint) (entity.Character, error)
	Create(entity.Character) (entity.Character, error)
}
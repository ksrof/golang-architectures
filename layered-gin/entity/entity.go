// layered-gin/entity/entity.go

package entity

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name 					string `gorm:"size:255;not null;unique" validate:"required,min=3,max=255" json:"name"` 
	Title 				string `gorm:"size:255;not null" validate:"required,min=3,max=255" json:"title"`
	Gender 				string `gorm:"size:100;not null" validate:"required,min=3,max=100" json:"gender"`
	Race 					string `gorm:"size:100;not null" validate:"required,min=3,max=100" json:"race"`
	Class 				string `gorm:"size:100;not null" validate:"required,min=3,max=100" json:"class"`
	Affiliation 	string `gorm:"size:100;not null" validate:"required,min=3,max=100" json:"affiliation"`
	Status 				string `gorm:"size:50;not null" validate:"required,min=3,max=50" json:"status"`
}
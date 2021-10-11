// layered-gin/entity/entity.go

package entity

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name 					string `gorm:"size:255;not null;unique" json:"name"` 
	Title 				string `gorm:"size:255;not null" json:"title"`
	Gender 				string `gorm:"size:100;not null" json:"gender"`
	Race 					string `gorm:"size:100;not null" json:"race"`
	Class 				string `gorm:"size:100;not null" json:"class"`
	Affiliation 	string `gorm:"size:100;not null" json:"affiliation"`
	Status 				string `gorm:"size:50;not null" json:"status"`
}
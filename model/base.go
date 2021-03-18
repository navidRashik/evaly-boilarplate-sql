package model

import (
	"gorm.io/gorm"
)

// Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
// It may be embedded into your model or you may build your own model without it
//    type User struct {
//      gorm.Model
//    }
type BaseModel struct {
	gorm.Model
	CreatedBy string `json:"created_by" gorm:"index"`
	UpdatedBy string `json:"updated_by"`
}

package model

import (
	"gorm.io/gorm"
)

// Status defines status
type Status string

const (
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
)

// BrandInfo defines brand model
type BrandInfo struct {
	gorm.Model
	Name        string
	Approved    bool
	Slug        string   `gorm:"unique"`
	Description string
	BrandType   string
	Categories  string
	ImageURL    string
	Status      Status
	BrandScore  float64
	Version     int64
}

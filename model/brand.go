package model

// Status defines status
type Status string

const (
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
)

// BrandInfo defines brand model
type BrandInfo struct {
	BaseModel
	Name        string   `json:"name,omitempty"`
	Approved    bool     `json:"approved,omitempty"`
	Slug        string   `gorm:"unique" json:"slug,omitempty"`
	Description string   `json:"description,omitempty"`
	BrandType   string   `json:"brand_type,omitempty"`
	Categories  string   `json:"categories,omitempty"`
	ImageURL    string   `json:"image_url,omitempty"`
	Status      Status   `json:"status,omitempty"`
	BrandScore  float64  `json:"brand_score,omitempty"`
	Version     int64    `json:"version"`
}

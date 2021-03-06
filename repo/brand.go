package repo

import (
	"context"
	"database/sql"
	"errors"
	"gorm.io/gorm/clause"
	"log"

	infraSql "go-mysql-boilerplate/infra/sql"
	"go-mysql-boilerplate/model"
)

// BrandRepo returns brand repo
type BrandRepo interface {
	CreateBrand(ctx context.Context, bi *model.BrandInfo) error
	UpdateBrand(ctx context.Context, bi model.BrandInfo) error
	ListBrands(ctx context.Context, selector *model.BrandInfo, skip, limit int) ([]*model.BrandInfo, error)
	GetBrandDetails(ctx context.Context, slug string) (*model.BrandInfo, error)
}

// MyBrand brand repo
type MyBrand struct {
	table string
	db    *infraSql.DB
}

// NewBrand returns new brand repo
func NewBrand(table string, db *infraSql.DB) BrandRepo {
	return &MyBrand{
		table: table,
		db:    db,
	}
}

// Indices returns indices
func (*MyBrand) Indices() []interface{} {
	res := make([]interface{}, 0)
	return res
}

// EnsureIndices ...
func (p *MyBrand) EnsureIndices() error {
	return p.db.EnsureIndices(context.Background(), p.Indices())
}

//// DropIndices ...
//func (p *MyBrand) DropIndices() error {
//	return p.db.DropIndices(context.Background(), p.table, p.Indices())
//}

// Create ...
func (p *MyBrand) CreateBrand(ctx context.Context, bi *model.BrandInfo) error {
	return p.db.Database.Create(bi).Error
}

// ListBrands ...
func (p *MyBrand) ListBrands(ctx context.Context, selector *model.BrandInfo, skip, limit int) ([]*model.BrandInfo, error) {
	brands := make([]*model.BrandInfo, 0)

	//tx := p.db.Database.Preload("x").Preload("y").Limit(limit).Offset(skip).Where(selector)
	//tx := p.db.Database.Where("approved = ?", false).Preload(clause.Associations).Limit(limit).Offset(skip).Where(selector)
	tx := p.db.Database.Preload(clause.Associations).Limit(limit).Offset(skip).Where(selector)

	err := tx.Find(&brands).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return brands, nil
}

// GetBrandDetails ...
func (p *MyBrand) GetBrandDetails(ctx context.Context, slug string) (*model.BrandInfo, error) {
	selector := &model.BrandInfo{Slug: slug}

	brand := &model.BrandInfo{}

	err := p.db.Database.First(brand, selector).Error
	if err != nil || brand.ID == 0 {
		log.Println("brand not found for slug", slug)
		return nil, sql.ErrNoRows
	} else {
		log.Println("brand found:", brand)
	}

	return brand, nil
}

func (p *MyBrand) UpdateBrand(ctx context.Context, param model.BrandInfo) error {
	selector := &model.BrandInfo{Slug: param.Slug}

	brand := &model.BrandInfo{}

	err := p.db.Database.First(brand, selector).Error
	if err != nil || brand.ID == 0 {
		log.Println("brand not found for slug", param.Slug)
		return errors.New("brand not found for slug")
	} else {
		log.Println("brand found:", param)
		newParam := map[string]interface{}{
			"name": param.Name,
			"brand_score": param.BrandScore,
			"approved": param.Approved,
			"status": param.Status}
		if param.Description != "" {
			newParam["description"] = param.Description
		}
		p.db.Database.Model(brand).Updates(newParam)
	}
	return nil
}
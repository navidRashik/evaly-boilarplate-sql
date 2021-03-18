package service

import (
	"context"
	"fmt"
	"errors"

	"go-mysql-boilerplate/logger"
	"go-mysql-boilerplate/model"
	"go-mysql-boilerplate/utils"
)

// SetLogger ...
func (c *Service) SetLogger(l logger.StructLogger) {
	c.log = l
}

// ListBrand ...
func (c *Service) ListBrand(ctx context.Context, pager utils.Pager) ([]*model.BrandInfo, error) {
	tid := utils.GetTracingID(ctx)

	c.log.Println("ListBrands", tid, "listing product brands from database")
	brands, err := c.brandRepo.ListBrands(ctx, &model.BrandInfo{}, int(pager.Skip), int(pager.Limit))
	if err != nil {
		fmt.Println(err)
		c.log.Errorln("ListBrands", tid, err.Error())
		return nil, err
	}

	c.log.Println("ListBrands", tid, "sent response successfully")
	return brands, nil
}

// AddBrand ...
func (c *Service) AddBrand(ctx context.Context, brand *model.BrandInfo) error {
	tid := utils.GetTracingID(ctx)

	c.log.Println("AddBrand", tid, "inserting brands into database")
	err := c.brandRepo.CreateBrand(ctx, brand)
	if err != nil {
		c.log.Errorln("AddBrand", tid, err.Error())
		return err
	}

	c.log.Println("AddBrand", tid, "sent response successfully")
	return nil
}


func (c *Service) UpdateBrand(ctx context.Context, param model.BrandInfo) error {
	tid := utils.GetTracingID(ctx)
	c.log.Printf("UpdateBrand", tid, "param %v", param)
	slug := param.Slug
	if slug == "" {
		return errors.New("slug missing")
	}

	if err := c.brandRepo.UpdateBrand(ctx, param); err != nil {
		c.log.Errorln("UpdateBrand", tid, err.Error())
		return err
	}
	return nil
}

func (c *Service) GetBrand(ctx context.Context, slug string) (*model.BrandInfo, error) {
	tid := utils.GetTracingID(ctx)
	c.log.Println("GetUserBalance", tid, "getting brand from database")
	brand, err := c.brandRepo.GetBrandDetails(ctx, slug)
	if err != nil {
		c.log.Errorln("GetBrand", tid, err.Error())
		return nil, err
	}
	c.log.Println("GetBrand", tid, "sent response successfully")
	return brand, nil
}




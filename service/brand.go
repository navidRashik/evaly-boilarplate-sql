package service

import (
	"context"
	"fmt"

	"go-mysql-boilerplate/logger"
	"go-mysql-boilerplate/model"
	"go-mysql-boilerplate/utils"
)

// SetLogger ...
func (c *Service) SetLogger(l logger.StructLogger) {
	c.log = l
}

// ListBrand ...
func (c *Service) ListBrand(ctx context.Context, pager utils.Pager) ([]model.BrandInfo, error) {
	tid := utils.GetTracingID(ctx)

	c.log.Println("ListBrands", tid, "listing product brands from database")
	brands, err := c.brandRepo.ListBrands(ctx, "", pager.Skip, pager.Limit)
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
	err := c.brandRepo.Create(ctx, brand)
	if err != nil {
		c.log.Errorln("AddBrand", tid, err.Error())
		return err
	}

	c.log.Println("AddBrand", tid, "sent response successfully")
	return nil
}

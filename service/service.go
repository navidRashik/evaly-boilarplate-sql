package service

import (
	"go-mysql-boilerplate/logger"
	"go-mysql-boilerplate/repo"
)

// Service ...
type Service struct {
	log       logger.StructLogger
	brandRepo repo.BrandRepo
}

// New ...
func New(brandRepo repo.BrandRepo, lgr logger.StructLogger) *Service {
	return &Service{
		log:       lgr,
		brandRepo: brandRepo,
	}
}

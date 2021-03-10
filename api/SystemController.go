package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"go-mysql-boilerplate/api/response"
	infraSql "go-mysql-boilerplate/infra/sql"
)

// SystemController ..
type SystemController struct {
	DB *infraSql.DB
}

// NewSystemController ..
func NewSystemController(db *infraSql.DB) *SystemController {
	return &SystemController{
		DB: db,
	}
}

func (s *SystemController) systemCheck(w http.ResponseWriter, r *http.Request) {
	if err := s.connCheck(); err != nil {
		_ = response.ServeJSON(w, http.StatusInternalServerError, nil, nil, err.Error(), nil)
		return
	}
	response.ServeJSONData(w, "ok", http.StatusOK)
	return
}

func (s *SystemController) apiCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("apiCheck")
	if err := s.connCheck(); err != nil {
		_ = response.ServeJSON(w, http.StatusInternalServerError, nil, nil, err.Error(), nil)
		return
	}
	response.ServeJSONData(w, "ok", http.StatusOK)
	return
}

func (s *SystemController) connCheck() error {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	log.Println("DB ping")

	return s.DB.Ping()
}

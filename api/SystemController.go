package api

import (
	"go-mysql-boilerplate/api/response"
	infraSql "go-mysql-boilerplate/infra/sql"
	"log"
	"net/http"
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
	log.Println("DB ping")
	return s.DB.Ping()
}

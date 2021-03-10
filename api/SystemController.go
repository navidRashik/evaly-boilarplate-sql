package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"go-mysql-boilerplate/api/response"
	"go-mysql-boilerplate/infra"
)

// SystemController ..
type SystemController struct {
	db infra.DB
}

// NewSystemController ..
func NewSystemController(db infra.DB) *SystemController {
	return &SystemController{
		db: db,
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
	log.Println("db ping")
	//if err := s.db.Ping(ctx); err != nil {
	//	return fmt.Errorf("mongo conn error: %v", err)
	//}

	return nil
}

package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"

	"go-mysql-boilerplate/api/response"
	"go-mysql-boilerplate/logger"
	"go-mysql-boilerplate/model"
	"go-mysql-boilerplate/service"
	"go-mysql-boilerplate/utils"
)

// BrandsController ...
type BrandsController struct {
	svc *service.Service
	lgr logger.StructLogger
}

// NewBrandsController ...
func NewBrandsController(svc *service.Service) *BrandsController {
	return &BrandsController{
		svc: svc,
	}
}

// SetLogger ...
func (cc *BrandsController) SetLogger(lgr logger.StructLogger) {
	cc.lgr = lgr
}

// ListBrand ...
func (cc *BrandsController) ListBrand(w http.ResponseWriter, r *http.Request) {
	tid := utils.GetTracingID(r.Context())
	pageQ, skipQ, limitQ, err := parseSkipLimit(r, 10, 100)
	if err != nil {
		cc.lgr.Errorln("listBrands", tid, err.Error())
		_ = response.ServeJSON(w, http.StatusBadRequest, nil, nil, err.Error(), nil)
		return
	}
	pager := utils.Pager{Skip: skipQ, Limit: limitQ}

	cc.lgr.Println("listBrands", tid, "getting brands")
	result, err := cc.svc.ListBrand(r.Context(), pager)
	if err != nil {
		cc.lgr.Errorln("listBrands", tid, err.Error())
		_ = response.ServeJSON(w, http.StatusInternalServerError, nil, nil, err.Error(), nil)
		return
	}

	cc.lgr.Println("listBrands", tid, "sending response")
	prev, next := getNextPreviousPager(r.URL.Path, pageQ, limitQ)
	_ = response.ServeJSON(w, http.StatusOK, prev, next, utils.SuccessMessage, result)
	return
}

// AddBrand ...
func (cc *BrandsController) AddBrand(w http.ResponseWriter, r *http.Request) {
	tid := utils.GetTracingID(r.Context())

	var b *model.BrandInfo
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		_ = response.ServeJSON(w, http.StatusBadRequest, nil, nil, utils.RequiredFieldMessage(), nil)
		return
	}

	cc.lgr.Println("AddBrand", tid, "inserting brand")
	err := cc.svc.AddBrand(r.Context(), b)
	if err != nil {
		cc.lgr.Errorln("AddBrand", tid, err.Error())
		_ = response.ServeJSON(w, http.StatusInternalServerError, nil, nil, err.Error(), nil)
		return
	}

	cc.lgr.Println("AddBrand", tid, "sending response")
	_ = response.ServeJSON(w, http.StatusOK, nil, nil, utils.SuccessMessage, nil)
	return
}

// Update UserBalanceTransactions ...
//func (cc *BrandsController) UpdateBrand(w http.ResponseWriter, r *http.Request) {
//	tid := utils.GetTracingID(r.Context())
//	//fmt.Println("method w", w, "r ", r)
//	param := model.BrandInfo{}
//
//	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
//		_ = response.ServeJSON(w, http.StatusBadRequest, nil, nil, utils.RequiredFieldMessage(), nil)
//		return
//	}
//
//	cc.lgr.Println("UpdateUserBalance", tid, "updating UserBalance")
//	err := cc.svc.AddBrand(r.Context(), param)
//	if err != nil {
//		cc.lgr.Errorln("UpdateUserBalance", tid, err.Error())
//		_ = response.ServeJSON(w, http.StatusInternalServerError, nil, nil, err.Error(), nil)
//		return
//	}
//
//	cc.lgr.Println("UpdateUserBalance", tid, "sending response")
//	_ = response.ServeJSON(w, http.StatusOK, nil, nil, utils.SuccessMessage, nil)
//	return
//}

// Get UserBalance ...
func (c *BrandsController) GetBrand(w http.ResponseWriter, r *http.Request) {
	tid := utils.GetTracingID(r.Context())
	c.lgr.Println("getBrandDetails", tid, "Initial")
	brandSlug := chi.URLParam(r, "slug")
	fmt.Println("brandSlug ", chi.URLParam)
	if brandSlug == "" {
		_ = response.ServeJSON(w, http.StatusBadRequest, nil, nil, utils.RequiredFieldMessage("slug"), nil)
		return
	}
	result, err := c.svc.GetBrand(r.Context(), brandSlug)
	if err != nil {
		c.lgr.Errorln("getBrandDetails", tid, err.Error())
		_ = response.ServeJSON(w, http.StatusInternalServerError, nil, nil, err.Error(), nil)
		return
	}

	_ = response.ServeJSON(w, http.StatusOK, nil, nil, utils.SuccessMessage, result)
}



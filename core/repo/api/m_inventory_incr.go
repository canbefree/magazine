package api

import (
	"net/http"

	"github.com/canbefree/magazine/core/repo/repo_dao"
	"github.com/canbefree/magazine/core/repo/repo_model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configMInventoryIncrRouter(router *httprouter.Router) {
	router.GET("/minventoryincr", GetAllMInventoryIncr)
	router.POST("/minventoryincr", AddMInventoryIncr)
	router.GET("/minventoryincr/:argID", GetMInventoryIncr)
	router.PUT("/minventoryincr/:argID", UpdateMInventoryIncr)
	router.DELETE("/minventoryincr/:argID", DeleteMInventoryIncr)
}

func configGinMInventoryIncrRouter(router gin.IRoutes) {
	router.GET("/minventoryincr", ConverHttprouterToGin(GetAllMInventoryIncr))
	router.POST("/minventoryincr", ConverHttprouterToGin(AddMInventoryIncr))
	router.GET("/minventoryincr/:argID", ConverHttprouterToGin(GetMInventoryIncr))
	router.PUT("/minventoryincr/:argID", ConverHttprouterToGin(UpdateMInventoryIncr))
	router.DELETE("/minventoryincr/:argID", ConverHttprouterToGin(DeleteMInventoryIncr))
}

// GetAllMInventoryIncr is a function to get a slice of record(s) from m_inventory_incr table in the magazine database
// @Summary Get list of MInventoryIncr
// @Tags MInventoryIncr
// @Description GetAllMInventoryIncr is a handler to get a slice of record(s) from m_inventory_incr table in the magazine database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]repo_model.MInventoryIncr}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventoryincr [get]
// http "http://localhost:8080/minventoryincr?page=0&pagesize=20" X-Api-User:user123
func GetAllMInventoryIncr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "m_inventory_incr", repo_model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := repo_dao.GetAllMInventoryIncr(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetMInventoryIncr is a function to get a single record from the m_inventory_incr table in the magazine database
// @Summary Get record from table MInventoryIncr by  argID
// @Tags MInventoryIncr
// @ID argID
// @Description GetMInventoryIncr is a function to get a single record from the m_inventory_incr table in the magazine database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} repo_model.MInventoryIncr
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /minventoryincr/{argID} [get]
// http "http://localhost:8080/minventoryincr/1" X-Api-User:user123
func GetMInventoryIncr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_incr", repo_model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := repo_dao.GetMInventoryIncr(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddMInventoryIncr add to add a single record to m_inventory_incr table in the magazine database
// @Summary Add an record to m_inventory_incr table
// @Description add to add a single record to m_inventory_incr table in the magazine database
// @Tags MInventoryIncr
// @Accept  json
// @Produce  json
// @Param MInventoryIncr body repo_model.MInventoryIncr true "Add MInventoryIncr"
// @Success 200 {object} repo_model.MInventoryIncr
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventoryincr [post]
// echo '{"id": 58,"change_qty": "MdfnZUTYsHIreyeWTDWugQuqG","snopshot": "ZpnSKNRrpxptMoZBJZNtQfUbM","tran_id": 38,"incr_type": 8,"created_at": "2173-11-26T10:13:04.173313035+08:00","updated_at": "2252-08-28T22:37:06.910141183+08:00"}' | http POST "http://localhost:8080/minventoryincr" X-Api-User:user123
func AddMInventoryIncr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	minventoryincr := &repo_model.MInventoryIncr{}

	if err := readJSON(r, minventoryincr); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := minventoryincr.BeforeSave(); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
	}

	minventoryincr.Prepare()

	if err := minventoryincr.Validate(repo_model.Create); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_incr", repo_model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	minventoryincr, _, err = repo_dao.AddMInventoryIncr(ctx, minventoryincr)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, minventoryincr)
}

// UpdateMInventoryIncr Update a single record from m_inventory_incr table in the magazine database
// @Summary Update an record in table m_inventory_incr
// @Description Update a single record from m_inventory_incr table in the magazine database
// @Tags MInventoryIncr
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  MInventoryIncr body repo_model.MInventoryIncr true "Update MInventoryIncr record"
// @Success 200 {object} repo_model.MInventoryIncr
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventoryincr/{argID} [put]
// echo '{"id": 58,"change_qty": "MdfnZUTYsHIreyeWTDWugQuqG","snopshot": "ZpnSKNRrpxptMoZBJZNtQfUbM","tran_id": 38,"incr_type": 8,"created_at": "2173-11-26T10:13:04.173313035+08:00","updated_at": "2252-08-28T22:37:06.910141183+08:00"}' | http PUT "http://localhost:8080/minventoryincr/1"  X-Api-User:user123
func UpdateMInventoryIncr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	minventoryincr := &repo_model.MInventoryIncr{}
	if err := readJSON(r, minventoryincr); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := minventoryincr.BeforeSave(); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
	}

	minventoryincr.Prepare()

	if err := minventoryincr.Validate(repo_model.Update); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_incr", repo_model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	minventoryincr, _, err = repo_dao.UpdateMInventoryIncr(ctx,
		argID,
		minventoryincr)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, minventoryincr)
}

// DeleteMInventoryIncr Delete a single record from m_inventory_incr table in the magazine database
// @Summary Delete a record from m_inventory_incr
// @Description Delete a single record from m_inventory_incr table in the magazine database
// @Tags MInventoryIncr
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} repo_model.MInventoryIncr
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /minventoryincr/{argID} [delete]
// http DELETE "http://localhost:8080/minventoryincr/1" X-Api-User:user123
func DeleteMInventoryIncr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_incr", repo_model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := repo_dao.DeleteMInventoryIncr(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}

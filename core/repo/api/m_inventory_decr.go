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

func configMInventoryDecrRouter(router *httprouter.Router) {
	router.GET("/minventorydecr", GetAllMInventoryDecr)
	router.POST("/minventorydecr", AddMInventoryDecr)
	router.GET("/minventorydecr/:argID", GetMInventoryDecr)
	router.PUT("/minventorydecr/:argID", UpdateMInventoryDecr)
	router.DELETE("/minventorydecr/:argID", DeleteMInventoryDecr)
}

func configGinMInventoryDecrRouter(router gin.IRoutes) {
	router.GET("/minventorydecr", ConverHttprouterToGin(GetAllMInventoryDecr))
	router.POST("/minventorydecr", ConverHttprouterToGin(AddMInventoryDecr))
	router.GET("/minventorydecr/:argID", ConverHttprouterToGin(GetMInventoryDecr))
	router.PUT("/minventorydecr/:argID", ConverHttprouterToGin(UpdateMInventoryDecr))
	router.DELETE("/minventorydecr/:argID", ConverHttprouterToGin(DeleteMInventoryDecr))
}

// GetAllMInventoryDecr is a function to get a slice of record(s) from m_inventory_decr table in the magazine database
// @Summary Get list of MInventoryDecr
// @Tags MInventoryDecr
// @Description GetAllMInventoryDecr is a handler to get a slice of record(s) from m_inventory_decr table in the magazine database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]repo_model.MInventoryDecr}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventorydecr [get]
// http "http://localhost:8080/minventorydecr?page=0&pagesize=20" X-Api-User:user123
func GetAllMInventoryDecr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "m_inventory_decr", repo_model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := repo_dao.GetAllMInventoryDecr(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetMInventoryDecr is a function to get a single record from the m_inventory_decr table in the magazine database
// @Summary Get record from table MInventoryDecr by  argID
// @Tags MInventoryDecr
// @ID argID
// @Description GetMInventoryDecr is a function to get a single record from the m_inventory_decr table in the magazine database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} repo_model.MInventoryDecr
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /minventorydecr/{argID} [get]
// http "http://localhost:8080/minventorydecr/1" X-Api-User:user123
func GetMInventoryDecr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_decr", repo_model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := repo_dao.GetMInventoryDecr(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddMInventoryDecr add to add a single record to m_inventory_decr table in the magazine database
// @Summary Add an record to m_inventory_decr table
// @Description add to add a single record to m_inventory_decr table in the magazine database
// @Tags MInventoryDecr
// @Accept  json
// @Produce  json
// @Param MInventoryDecr body repo_model.MInventoryDecr true "Add MInventoryDecr"
// @Success 200 {object} repo_model.MInventoryDecr
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventorydecr [post]
// echo '{"id": 11,"change_qty": "omdkFjFoifgiVLAuFWKljkHUC","snopshot": "oQLauvVgkDGnQgIPHLgbhLfBT","tran_id": 46,"decr_type": 21,"created_at": "2253-07-13T02:55:25.213858309+08:00","updated_at": "2291-08-03T13:49:38.27972354+08:00"}' | http POST "http://localhost:8080/minventorydecr" X-Api-User:user123
func AddMInventoryDecr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	minventorydecr := &repo_model.MInventoryDecr{}

	if err := readJSON(r, minventorydecr); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := minventorydecr.BeforeSave(); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
	}

	minventorydecr.Prepare()

	if err := minventorydecr.Validate(repo_model.Create); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_decr", repo_model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	minventorydecr, _, err = repo_dao.AddMInventoryDecr(ctx, minventorydecr)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, minventorydecr)
}

// UpdateMInventoryDecr Update a single record from m_inventory_decr table in the magazine database
// @Summary Update an record in table m_inventory_decr
// @Description Update a single record from m_inventory_decr table in the magazine database
// @Tags MInventoryDecr
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  MInventoryDecr body repo_model.MInventoryDecr true "Update MInventoryDecr record"
// @Success 200 {object} repo_model.MInventoryDecr
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventorydecr/{argID} [put]
// echo '{"id": 11,"change_qty": "omdkFjFoifgiVLAuFWKljkHUC","snopshot": "oQLauvVgkDGnQgIPHLgbhLfBT","tran_id": 46,"decr_type": 21,"created_at": "2253-07-13T02:55:25.213858309+08:00","updated_at": "2291-08-03T13:49:38.27972354+08:00"}' | http PUT "http://localhost:8080/minventorydecr/1"  X-Api-User:user123
func UpdateMInventoryDecr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	minventorydecr := &repo_model.MInventoryDecr{}
	if err := readJSON(r, minventorydecr); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := minventorydecr.BeforeSave(); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
	}

	minventorydecr.Prepare()

	if err := minventorydecr.Validate(repo_model.Update); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_decr", repo_model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	minventorydecr, _, err = repo_dao.UpdateMInventoryDecr(ctx,
		argID,
		minventorydecr)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, minventorydecr)
}

// DeleteMInventoryDecr Delete a single record from m_inventory_decr table in the magazine database
// @Summary Delete a record from m_inventory_decr
// @Description Delete a single record from m_inventory_decr table in the magazine database
// @Tags MInventoryDecr
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} repo_model.MInventoryDecr
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /minventorydecr/{argID} [delete]
// http DELETE "http://localhost:8080/minventorydecr/1" X-Api-User:user123
func DeleteMInventoryDecr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_decr", repo_model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := repo_dao.DeleteMInventoryDecr(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}

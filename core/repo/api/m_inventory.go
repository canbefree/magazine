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

func configMInventoryRouter(router *httprouter.Router) {
	router.GET("/minventory", GetAllMInventory)
	router.POST("/minventory", AddMInventory)
	router.GET("/minventory/:argID", GetMInventory)
	router.PUT("/minventory/:argID", UpdateMInventory)
	router.DELETE("/minventory/:argID", DeleteMInventory)
}

func configGinMInventoryRouter(router gin.IRoutes) {
	router.GET("/minventory", ConverHttprouterToGin(GetAllMInventory))
	router.POST("/minventory", ConverHttprouterToGin(AddMInventory))
	router.GET("/minventory/:argID", ConverHttprouterToGin(GetMInventory))
	router.PUT("/minventory/:argID", ConverHttprouterToGin(UpdateMInventory))
	router.DELETE("/minventory/:argID", ConverHttprouterToGin(DeleteMInventory))
}

// GetAllMInventory is a function to get a slice of record(s) from m_inventory table in the magazine database
// @Summary Get list of MInventory
// @Tags MInventory
// @Description GetAllMInventory is a handler to get a slice of record(s) from m_inventory table in the magazine database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]repo_model.MInventory}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventory [get]
// http "http://localhost:8080/minventory?page=0&pagesize=20" X-Api-User:user123
func GetAllMInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "m_inventory", repo_model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := repo_dao.GetAllMInventory(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetMInventory is a function to get a single record from the m_inventory table in the magazine database
// @Summary Get record from table MInventory by  argID
// @Tags MInventory
// @ID argID
// @Description GetMInventory is a function to get a single record from the m_inventory table in the magazine database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} repo_model.MInventory
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /minventory/{argID} [get]
// http "http://localhost:8080/minventory/1" X-Api-User:user123
func GetMInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory", repo_model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := repo_dao.GetMInventory(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddMInventory add to add a single record to m_inventory table in the magazine database
// @Summary Add an record to m_inventory table
// @Description add to add a single record to m_inventory table in the magazine database
// @Tags MInventory
// @Accept  json
// @Produce  json
// @Param MInventory body repo_model.MInventory true "Add MInventory"
// @Success 200 {object} repo_model.MInventory
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventory [post]
// echo '{"id": 98,"qty": "pjYtKaywFpvTjnxXpGUjOTsNt","rest_qty": "JQoVAcsrZNotOswMbTiOfNtEn","version": "ZgGABIZRGbEvtJlxgFevmmhKU","created_at": "2146-09-03T22:05:56.619844516+08:00","updated_at": "2156-09-10T01:33:42.35436937+08:00"}' | http POST "http://localhost:8080/minventory" X-Api-User:user123
func AddMInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	minventory := &repo_model.MInventory{}

	if err := readJSON(r, minventory); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := minventory.BeforeSave(); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
	}

	minventory.Prepare()

	if err := minventory.Validate(repo_model.Create); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory", repo_model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	minventory, _, err = repo_dao.AddMInventory(ctx, minventory)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, minventory)
}

// UpdateMInventory Update a single record from m_inventory table in the magazine database
// @Summary Update an record in table m_inventory
// @Description Update a single record from m_inventory table in the magazine database
// @Tags MInventory
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  MInventory body repo_model.MInventory true "Update MInventory record"
// @Success 200 {object} repo_model.MInventory
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventory/{argID} [put]
// echo '{"id": 98,"qty": "pjYtKaywFpvTjnxXpGUjOTsNt","rest_qty": "JQoVAcsrZNotOswMbTiOfNtEn","version": "ZgGABIZRGbEvtJlxgFevmmhKU","created_at": "2146-09-03T22:05:56.619844516+08:00","updated_at": "2156-09-10T01:33:42.35436937+08:00"}' | http PUT "http://localhost:8080/minventory/1"  X-Api-User:user123
func UpdateMInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	minventory := &repo_model.MInventory{}
	if err := readJSON(r, minventory); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := minventory.BeforeSave(); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
	}

	minventory.Prepare()

	if err := minventory.Validate(repo_model.Update); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory", repo_model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	minventory, _, err = repo_dao.UpdateMInventory(ctx,
		argID,
		minventory)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, minventory)
}

// DeleteMInventory Delete a single record from m_inventory table in the magazine database
// @Summary Delete a record from m_inventory
// @Description Delete a single record from m_inventory table in the magazine database
// @Tags MInventory
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} repo_model.MInventory
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /minventory/{argID} [delete]
// http DELETE "http://localhost:8080/minventory/1" X-Api-User:user123
func DeleteMInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory", repo_model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := repo_dao.DeleteMInventory(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}

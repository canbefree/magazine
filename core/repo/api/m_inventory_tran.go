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

func configMInventoryTranRouter(router *httprouter.Router) {
	router.GET("/minventorytran", GetAllMInventoryTran)
	router.POST("/minventorytran", AddMInventoryTran)
	router.GET("/minventorytran/:argID", GetMInventoryTran)
	router.PUT("/minventorytran/:argID", UpdateMInventoryTran)
	router.DELETE("/minventorytran/:argID", DeleteMInventoryTran)
}

func configGinMInventoryTranRouter(router gin.IRoutes) {
	router.GET("/minventorytran", ConverHttprouterToGin(GetAllMInventoryTran))
	router.POST("/minventorytran", ConverHttprouterToGin(AddMInventoryTran))
	router.GET("/minventorytran/:argID", ConverHttprouterToGin(GetMInventoryTran))
	router.PUT("/minventorytran/:argID", ConverHttprouterToGin(UpdateMInventoryTran))
	router.DELETE("/minventorytran/:argID", ConverHttprouterToGin(DeleteMInventoryTran))
}

// GetAllMInventoryTran is a function to get a slice of record(s) from m_inventory_tran table in the magazine database
// @Summary Get list of MInventoryTran
// @Tags MInventoryTran
// @Description GetAllMInventoryTran is a handler to get a slice of record(s) from m_inventory_tran table in the magazine database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]repo_model.MInventoryTran}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventorytran [get]
// http "http://localhost:8080/minventorytran?page=0&pagesize=20" X-Api-User:user123
func GetAllMInventoryTran(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "m_inventory_tran", repo_model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := repo_dao.GetAllMInventoryTran(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetMInventoryTran is a function to get a single record from the m_inventory_tran table in the magazine database
// @Summary Get record from table MInventoryTran by  argID
// @Tags MInventoryTran
// @ID argID
// @Description GetMInventoryTran is a function to get a single record from the m_inventory_tran table in the magazine database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} repo_model.MInventoryTran
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /minventorytran/{argID} [get]
// http "http://localhost:8080/minventorytran/1" X-Api-User:user123
func GetMInventoryTran(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_tran", repo_model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := repo_dao.GetMInventoryTran(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddMInventoryTran add to add a single record to m_inventory_tran table in the magazine database
// @Summary Add an record to m_inventory_tran table
// @Description add to add a single record to m_inventory_tran table in the magazine database
// @Tags MInventoryTran
// @Accept  json
// @Produce  json
// @Param MInventoryTran body repo_model.MInventoryTran true "Add MInventoryTran"
// @Success 200 {object} repo_model.MInventoryTran
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventorytran [post]
// echo '{"id": 94,"changet_type": 36,"change_qty": "LEdkDEJqodwJdcApdDSLLOOJK","tran_id": 36,"tran_status": 81,"created_at": "2242-02-24T04:57:56.658031026+08:00","updated_at": "2219-07-22T04:23:03.261779391+08:00","remark": "CZHCLLqWUdUWAdbDDgvJspnlp"}' | http POST "http://localhost:8080/minventorytran" X-Api-User:user123
func AddMInventoryTran(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	minventorytran := &repo_model.MInventoryTran{}

	if err := readJSON(r, minventorytran); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := minventorytran.BeforeSave(); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
	}

	minventorytran.Prepare()

	if err := minventorytran.Validate(repo_model.Create); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_tran", repo_model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	minventorytran, _, err = repo_dao.AddMInventoryTran(ctx, minventorytran)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, minventorytran)
}

// UpdateMInventoryTran Update a single record from m_inventory_tran table in the magazine database
// @Summary Update an record in table m_inventory_tran
// @Description Update a single record from m_inventory_tran table in the magazine database
// @Tags MInventoryTran
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  MInventoryTran body repo_model.MInventoryTran true "Update MInventoryTran record"
// @Success 200 {object} repo_model.MInventoryTran
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minventorytran/{argID} [put]
// echo '{"id": 94,"changet_type": 36,"change_qty": "LEdkDEJqodwJdcApdDSLLOOJK","tran_id": 36,"tran_status": 81,"created_at": "2242-02-24T04:57:56.658031026+08:00","updated_at": "2219-07-22T04:23:03.261779391+08:00","remark": "CZHCLLqWUdUWAdbDDgvJspnlp"}' | http PUT "http://localhost:8080/minventorytran/1"  X-Api-User:user123
func UpdateMInventoryTran(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	minventorytran := &repo_model.MInventoryTran{}
	if err := readJSON(r, minventorytran); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := minventorytran.BeforeSave(); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
	}

	minventorytran.Prepare()

	if err := minventorytran.Validate(repo_model.Update); err != nil {
		returnError(ctx, w, r, repo_dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_tran", repo_model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	minventorytran, _, err = repo_dao.UpdateMInventoryTran(ctx,
		argID,
		minventorytran)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, minventorytran)
}

// DeleteMInventoryTran Delete a single record from m_inventory_tran table in the magazine database
// @Summary Delete a record from m_inventory_tran
// @Description Delete a single record from m_inventory_tran table in the magazine database
// @Tags MInventoryTran
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} repo_model.MInventoryTran
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /minventorytran/{argID} [delete]
// http DELETE "http://localhost:8080/minventorytran/1" X-Api-User:user123
func DeleteMInventoryTran(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "m_inventory_tran", repo_model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := repo_dao.DeleteMInventoryTran(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}

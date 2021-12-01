package repo_dao

import (
	"context"
	"time"

	"github.com/canbefree/magazine/internel/repo/repo_model"
	"github.com/google/uuid"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllMInventoryTran is a function to get a slice of record(s) from m_inventory_tran table in the magazine database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllMInventoryTran(ctx context.Context, page, pagesize int, order string) (results []*repo_model.MInventoryTran, totalRows int64, err error) {

	resultOrm := DB.Model(&repo_model.MInventoryTran{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetMInventoryTran is a function to get a single record from the m_inventory_tran table in the magazine database
// error - ErrNotFound, db Find error
func GetMInventoryTran(ctx context.Context, argID int64) (record *repo_model.MInventoryTran, err error) {
	record = &repo_model.MInventoryTran{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddMInventoryTran is a function to add a single record to m_inventory_tran table in the magazine database
// error - ErrInsertFailed, db save call failed
func AddMInventoryTran(ctx context.Context, record *repo_model.MInventoryTran) (result *repo_model.MInventoryTran, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateMInventoryTran is a function to update a single record from m_inventory_tran table in the magazine database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateMInventoryTran(ctx context.Context, argID int64, updated *repo_model.MInventoryTran) (result *repo_model.MInventoryTran, RowsAffected int64, err error) {

	result = &repo_model.MInventoryTran{}
	db := DB.First(result, argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteMInventoryTran is a function to delete a single record from m_inventory_tran table in the magazine database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteMInventoryTran(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &repo_model.MInventoryTran{}
	db := DB.First(record, argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}

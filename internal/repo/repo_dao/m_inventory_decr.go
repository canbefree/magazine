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

// GetAllMInventoryDecr is a function to get a slice of record(s) from m_inventory_decr table in the magazine database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllMInventoryDecr(ctx context.Context, page, pagesize int, order string) (results []*repo_model.MInventoryDecr, totalRows int64, err error) {

	resultOrm := DB.Model(&repo_model.MInventoryDecr{})
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

// GetMInventoryDecr is a function to get a single record from the m_inventory_decr table in the magazine database
// error - ErrNotFound, db Find error
func GetMInventoryDecr(ctx context.Context, argID int64) (record *repo_model.MInventoryDecr, err error) {
	record = &repo_model.MInventoryDecr{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddMInventoryDecr is a function to add a single record to m_inventory_decr table in the magazine database
// error - ErrInsertFailed, db save call failed
func AddMInventoryDecr(ctx context.Context, record *repo_model.MInventoryDecr) (result *repo_model.MInventoryDecr, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateMInventoryDecr is a function to update a single record from m_inventory_decr table in the magazine database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateMInventoryDecr(ctx context.Context, argID int64, updated *repo_model.MInventoryDecr) (result *repo_model.MInventoryDecr, RowsAffected int64, err error) {

	result = &repo_model.MInventoryDecr{}
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

// DeleteMInventoryDecr is a function to delete a single record from m_inventory_decr table in the magazine database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteMInventoryDecr(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &repo_model.MInventoryDecr{}
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

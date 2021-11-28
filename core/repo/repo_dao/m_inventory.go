package repo_dao

import (
	"context"
	"time"

	"github.com/canbefree/magazine/core/repo/repo_model"
	"github.com/google/uuid"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllMInventory is a function to get a slice of record(s) from m_inventory table in the magazine database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllMInventory(ctx context.Context, page, pagesize int, order string) (results []*repo_model.MInventory, totalRows int64, err error) {

	resultOrm := DB.Model(&repo_model.MInventory{})
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

// GetMInventory is a function to get a single record from the m_inventory table in the magazine database
// error - ErrNotFound, db Find error
func GetMInventory(ctx context.Context, argID int64) (record *repo_model.MInventory, err error) {
	record = &repo_model.MInventory{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddMInventory is a function to add a single record to m_inventory table in the magazine database
// error - ErrInsertFailed, db save call failed
func AddMInventory(ctx context.Context, record *repo_model.MInventory) (result *repo_model.MInventory, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateMInventory is a function to update a single record from m_inventory table in the magazine database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateMInventory(ctx context.Context, argID int64, updated *repo_model.MInventory) (result *repo_model.MInventory, RowsAffected int64, err error) {

	result = &repo_model.MInventory{}
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

// DeleteMInventory is a function to delete a single record from m_inventory table in the magazine database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteMInventory(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &repo_model.MInventory{}
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

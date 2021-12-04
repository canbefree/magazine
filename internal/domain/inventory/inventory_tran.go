package inventory

import (
	"context"

	"github.com/canbefree/magazine/vars"
)

type InventoryTran struct {
	ID         TYPE_INVENTORY_TRAN_ID
	ChangeType TYPE_CHANGE_TYPE
	ChangeQty  uint32
	Status     TYPE_INVENTORY_TRAN_STATUS
	Remark     string
}

func NewInventoryTran(ctx context.Context, changeType TYPE_CHANGE_TYPE, changeQty uint32) (inventoryTran *InventoryTran, err error) {
	id := vars.Snowflake.GenerateID()
	inventoryTran = &InventoryTran{
		ID:         TYPE_INVENTORY_TRAN_ID(id),
		ChangeType: changeType,
		ChangeQty:  changeQty,
	}
	err = inventoryTran.validate(ctx)
	return
}

func (c *InventoryTran) validate(ctx context.Context) error {
	if c.ID == 0 {
		return ErrInventoryTranID
	}
	switch c.Status {
	case TRANSTATUS_DONE, TRANSTATUS_FAILED, TRANSTATUS_PROCESS:
	default:
		return ErrInventoryTranStatus
	}
	return nil
}

type InventoryIncrDetail struct {
	ID        TYPE_DETAIL_INCR_ID
	ChangeQty uint32
	IncrType  TYPE_INCR
	CommondDetail
}

func NewInventoryIncrDetail(ctx context.Context, changeQty uint32, incrType TYPE_INCR, common CommondDetail) (decrDetail *InventoryIncrDetail, err error) {
	id := vars.Snowflake.GenerateID()
	decrDetail = &InventoryIncrDetail{
		ID: TYPE_DETAIL_INCR_ID(id),
	}
	return
}

type InventoryDecrDetail struct {
	ID        TYPE_DETAIL_DECR_ID
	ChangeQty uint32
	DecrType  TYPE_DECR
	CommondDetail
}

func NewInventoryDecrDetail(ctx context.Context, changeQty uint32, incrType TYPE_INCR, common CommondDetail) (decrDetail *InventoryDecrDetail, err error) {
	id := vars.Snowflake.GenerateID()
	decrDetail = &InventoryDecrDetail{
		ID: TYPE_DETAIL_DECR_ID(id),
	}
	return
}

type CommondDetail struct {
	Snopshot *Inventory             // 库存快照
	TranID   TYPE_INVENTORY_TRAN_ID // 交易明细ID
}

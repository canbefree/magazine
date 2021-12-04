package inventory

import (
	"context"

	"github.com/canbefree/magazine/vars"
)

const (
	CHANGETYPE_INCR = 1 // 添加
	CHANGETYPE_DECR = 2 // 扣减
)

const (
	COINPACKETSTATUS_ENABLE  = 1 // 正常
	COINPACKETSTATUS_DISABLE = 2 // 禁用
)

// 对象
type Inventory struct {
	ID      TYPE_INVENTORY_ID
	RelType TYPE_REL_TYPE
	RelID   TYPE_REL_ID
	Qty     uint64
	RestQty uint64
}

func NewInventory(ctx context.Context, qty uint64, relType TYPE_REL_TYPE, relID TYPE_REL_ID) (inventory *Inventory, err error) {
	id := vars.Snowflake.GenerateID()
	inventory = &Inventory{
		ID:      TYPE_INVENTORY_ID(id),
		RelType: relType,
		RelID:   relID,
		Qty:     qty,
	}
	err = inventory.validate(ctx)
	return
}

func (c *Inventory) validate(ctx context.Context) error {
	if c.ID == 0 {
		return ErrInventoryID
	}
	return nil
}

// fucn ()

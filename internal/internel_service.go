package internal

import (
	"context"

	"github.com/canbefree/magazine/internal/domain/inventory"
)

// go:generate gex
type InventoryServiceIFace interface {
	// GetInventory 获取当前总库存数量
	GetInventory(ctx context.Context, relType inventory.TYPE_REL_TYPE, relID inventory.TYPE_REL_ID) (*inventory.Inventory, error)
	// Incr 添加库存
	Incr(ctx context.Context, relType inventory.TYPE_REL_TYPE, relID inventory.TYPE_REL_ID, incrType inventory.TYPE_INCR, changeQty uint64) (inventory.TYPE_INVENTORY_TRAN_ID, error)
	// Decr 库存扣减
	Decr(ctx context.Context, relType inventory.TYPE_REL_TYPE, relID inventory.TYPE_REL_ID, incrType inventory.TYPE_DECR, changeQty uint64) (*inventory.TYPE_INVENTORY_TRAN_ID, error)
	// TranRollback 根据交易记录回滚
	TranRollback(ctx context.Context, tranID inventory.TYPE_INVENTORY_TRAN_ID) error
}

type DefaultInventoryService struct {
	InventoryRepoIFace
}

func newDefaultIntentoryService(ctx context.Context) *DefaultInventoryService {
	return &DefaultInventoryService{
		InventoryRepoIFace: getInventoryRepo(ctx),
	}
}

// GetInventory 获取当前总库存数量
func (s *DefaultInventoryService) GetInventory(ctx context.Context, relType inventory.TYPE_REL_TYPE, relID inventory.TYPE_REL_ID) (*inventory.Inventory, error) {
	panic("TODO")
}

// Incr 添加库存
func (s *DefaultInventoryService) Incr(ctx context.Context, relType inventory.TYPE_REL_TYPE, relID inventory.TYPE_REL_ID, incrType inventory.TYPE_INCR, changeQty uint64) (inventory.TYPE_INVENTORY_TRAN_ID, error) {
	panic("TODO ")
}

// Decr 库存扣减
func (s *DefaultInventoryService) Decr(ctx context.Context, relType inventory.TYPE_REL_TYPE, relID inventory.TYPE_REL_ID, incrType inventory.TYPE_DECR, changeQty uint64) (*inventory.TYPE_INVENTORY_TRAN_ID, error) {
	panic("TODO ")
}

// TranRollback 回滚明细
func (s *DefaultInventoryService) TranRollback(ctx context.Context, tranID inventory.TYPE_INVENTORY_TRAN_ID) error {
	panic("TODO ")
}

type Factory struct {
}

// InventoryRepo
type InventoryRepoIFace interface {
	//
	GetInventory(ctx context.Context, tranID inventory.TYPE_INVENTORY_TRAN_ID) (*inventory.InventoryTran, error)
}

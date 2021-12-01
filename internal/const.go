package internel

import "errors"

// errors
var (
	// inventory
	ErrInventoryID = errors.New("inventory id empty")

	// inventory tran
	ErrInventoryTranID     = errors.New("inventory tran id empty")
	ErrInventoryTranStatus = errors.New("inventory tran status empty")
)

type TYPE_INVENTORY_ID uint64
type TYPE_REL_TYPE uint64
type TYPE_REL_ID uint64

// TranStatus
type TYPE_INVENTORY_TRAN_ID uint64
type TYPE_CHANGE_TYPE uint32
type TYPE_INVENTORY_TRAN_STATUS uint32
type TYPE_TRANSTATUS uint32

type TYPE_DETAIL_INCR_ID uint64
type TYPE_DETAIL_DECR_ID uint64

const (
	TRANSTATUS_PROCESS TYPE_INVENTORY_TRAN_STATUS = iota + 1
	TRANSTATUS_DONE
	TRANSTATUS_FAILED
)

// TranDetail
type TYPE_INCR uint32

const (
	INCRTYPE_INIT    TYPE_INCR = iota + 1 // 库存初始化
	INCRTYPE_STORAGE                      // 入库
)

type TYPE_DECR uint32

const (
	DECRTYPE_STOCK_OUT TYPE_DECR = iota + 1
)

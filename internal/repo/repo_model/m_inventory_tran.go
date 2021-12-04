package repo_model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `m_inventory_tran` (
  `id` bigint(20) unsigned NOT NULL,
  `changet_type` tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '入库类型(1 出库, 2 入库)',
  `change_qty` int(11) unsigned NOT NULL DEFAULT 0,
  `tran_id` bigint(20) unsigned NOT NULL COMMENT '库存流水明细记录ID',
  `tran_status` tinyint(4) unsigned NOT NULL DEFAULT 1 COMMENT '明细状态 (1 处理中，2 已完成，3 处理失败)',
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  `remark` varchar(256) NOT NULL DEFAULT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存明细申请记录'

JSON Sample
-------------------------------------
{    "id": 11,    "changet_type": 59,    "change_qty": 32,    "tran_id": 97,    "tran_status": 71,    "created_at": "2169-10-19T15:47:30.251331361+08:00",    "updated_at": "2058-08-28T22:12:18.709260626+08:00",    "remark": "LAotCSgbCjqXNnAhrVLkgfdgG"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// MInventoryTran struct is a row record of the m_inventory_tran table in the magazine database
type MInventoryTran struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: false  col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;column:id;type:ubigint;" json:"id"`
	//[ 1] changet_type                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	ChangetType uint32 `gorm:"column:changet_type;type:utinyint;default:0;" json:"changet_type"` // 入库类型(1 出库, 2 入库)
	//[ 2] change_qty                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ChangeQty uint32 `gorm:"column:change_qty;type:uint;default:0;" json:"change_qty"`
	//[ 3] tran_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	TranID uint64 `gorm:"column:tran_id;type:ubigint;" json:"tran_id"` // 库存流水明细记录ID
	//[ 4] tran_status                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	TranStatus uint32 `gorm:"column:tran_status;type:utinyint;default:1;" json:"tran_status"` // 明细状态 (1 处理中，2 已完成，3 处理失败)
	//[ 5] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [current_timestamp()]
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 6] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [0000-00-00 00:00:00]
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:0000-00-00 00:00:00;" json:"updated_at"`
	//[ 7] remark                                         varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: [备注]
	Remark string `gorm:"column:remark;type:varchar;size:256;default:备注;" json:"remark"`
}

var m_inventory_tranTableInfo = &TableInfo{
	Name: "m_inventory_tran",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "changet_type",
			Comment:            `入库类型(1 出库, 2 入库)`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "ChangetType",
			GoFieldType:        "uint32",
			JSONFieldName:      "changet_type",
			ProtobufFieldName:  "changet_type",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "change_qty",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ChangeQty",
			GoFieldType:        "uint32",
			JSONFieldName:      "change_qty",
			ProtobufFieldName:  "change_qty",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "tran_id",
			Comment:            `库存流水明细记录ID`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "TranID",
			GoFieldType:        "uint64",
			JSONFieldName:      "tran_id",
			ProtobufFieldName:  "tran_id",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "tran_status",
			Comment:            `明细状态 (1 处理中，2 已完成，3 处理失败)`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "TranStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "tran_status",
			ProtobufFieldName:  "tran_status",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "remark",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MInventoryTran) TableName() string {
	return "m_inventory_tran"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MInventoryTran) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MInventoryTran) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MInventoryTran) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MInventoryTran) TableInfo() *TableInfo {
	return m_inventory_tranTableInfo
}

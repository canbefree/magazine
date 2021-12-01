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


CREATE TABLE `m_inventory_decr` (
  `id` bigint(20) unsigned NOT NULL,
  `change_qty` int(10) unsigned NOT NULL DEFAULT 0,
  `snopshot` varchar(256) NOT NULL DEFAULT '库存快照',
  `tran_id` bigint(20) unsigned NOT NULL COMMENT '库存流水明细记录ID',
  `incr_type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '出库类型(1 普通库存扣减)',
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存出库表'

JSON Sample
-------------------------------------
{    "id": 54,    "change_qty": 84,    "snopshot": "IjsBFfqkRmhMeasUHweVriiVu",    "tran_id": 54,    "incr_type": 12,    "created_at": "2035-03-05T07:13:32.617748489+08:00",    "updated_at": "2230-03-26T05:53:50.964291956+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 3] column is set for unsigned



*/

// MInventoryDecr struct is a row record of the m_inventory_decr table in the magazine database
type MInventoryDecr struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: false  col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;column:id;type:ubigint;" json:"id"`
	//[ 1] change_qty                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ChangeQty uint32 `gorm:"column:change_qty;type:uint;default:0;" json:"change_qty"`
	//[ 2] snopshot                                       varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: [库存快照]
	Snopshot string `gorm:"column:snopshot;type:varchar;size:256;default:库存快照;" json:"snopshot"`
	//[ 3] tran_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	TranID uint64 `gorm:"column:tran_id;type:ubigint;" json:"tran_id"` // 库存流水明细记录ID
	//[ 4] incr_type                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	IncrType int32 `gorm:"column:incr_type;type:tinyint;default:1;" json:"incr_type"` // 出库类型(1 普通库存扣减)
	//[ 5] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [current_timestamp()]
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 6] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [0000-00-00 00:00:00]
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:0000-00-00 00:00:00;" json:"updated_at"`
}

var m_inventory_decrTableInfo = &TableInfo{
	Name: "m_inventory_decr",
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "snopshot",
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
			GoFieldName:        "Snopshot",
			GoFieldType:        "string",
			JSONFieldName:      "snopshot",
			ProtobufFieldName:  "snopshot",
			ProtobufType:       "string",
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
			Name:               "incr_type",
			Comment:            `出库类型(1 普通库存扣减)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IncrType",
			GoFieldType:        "int32",
			JSONFieldName:      "incr_type",
			ProtobufFieldName:  "incr_type",
			ProtobufType:       "int32",
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
	},
}

// TableName sets the insert table name for this struct type
func (m *MInventoryDecr) TableName() string {
	return "m_inventory_decr"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MInventoryDecr) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MInventoryDecr) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MInventoryDecr) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MInventoryDecr) TableInfo() *TableInfo {
	return m_inventory_decrTableInfo
}

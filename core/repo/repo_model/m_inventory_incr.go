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


CREATE TABLE `m_inventory_incr` (
  `id` bigint(20) NOT NULL,
  `change_qty` varchar(100) NOT NULL DEFAULT '0',
  `snopshot` varchar(256) NOT NULL DEFAULT '库存快照',
  `tran_id` bigint(20) NOT NULL COMMENT '库存流水明细记录ID',
  `incr_type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '入库类型(1 库存初始化, 2 库存返换)',
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存入库表'

JSON Sample
-------------------------------------
{    "id": 58,    "change_qty": "MdfnZUTYsHIreyeWTDWugQuqG",    "snopshot": "ZpnSKNRrpxptMoZBJZNtQfUbM",    "tran_id": 38,    "incr_type": 8,    "created_at": "2173-11-26T10:13:04.173313035+08:00",    "updated_at": "2252-08-28T22:37:06.910141183+08:00"}



*/

// MInventoryIncr struct is a row record of the m_inventory_incr table in the magazine database
type MInventoryIncr struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id"`
	//[ 1] change_qty                                     varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [0]
	ChangeQty string `gorm:"column:change_qty;type:varchar;size:100;default:0;" json:"change_qty"`
	//[ 2] snopshot                                       varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: [库存快照]
	Snopshot string `gorm:"column:snopshot;type:varchar;size:256;default:库存快照;" json:"snopshot"`
	//[ 3] tran_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	TranID int64 `gorm:"column:tran_id;type:bigint;" json:"tran_id"` // 库存流水明细记录ID
	//[ 4] incr_type                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	IncrType int32 `gorm:"column:incr_type;type:tinyint;default:1;" json:"incr_type"` // 入库类型(1 库存初始化, 2 库存返换)
	//[ 5] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [current_timestamp()]
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 6] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [0000-00-00 00:00:00]
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:0000-00-00 00:00:00;" json:"updated_at"`
}

var m_inventory_incrTableInfo = &TableInfo{
	Name: "m_inventory_incr",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "change_qty",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ChangeQty",
			GoFieldType:        "string",
			JSONFieldName:      "change_qty",
			ProtobufFieldName:  "change_qty",
			ProtobufType:       "string",
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
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "TranID",
			GoFieldType:        "int64",
			JSONFieldName:      "tran_id",
			ProtobufFieldName:  "tran_id",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "incr_type",
			Comment:            `入库类型(1 库存初始化, 2 库存返换)`,
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
func (m *MInventoryIncr) TableName() string {
	return "m_inventory_incr"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MInventoryIncr) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MInventoryIncr) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MInventoryIncr) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MInventoryIncr) TableInfo() *TableInfo {
	return m_inventory_incrTableInfo
}

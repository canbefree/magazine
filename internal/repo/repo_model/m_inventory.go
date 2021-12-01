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


CREATE TABLE `m_inventory` (
  `id` bigint(20) unsigned NOT NULL,
  `rel_type` bigint(20) unsigned NOT NULL COMMENT '实体类型',
  `rel_id` bigint(20) unsigned NOT NULL COMMENT '实体',
  `qty` int(10) unsigned NOT NULL DEFAULT 0,
  `rest_qty` int(10) unsigned NOT NULL DEFAULT 0,
  `version` int(10) unsigned NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存总表'

JSON Sample
-------------------------------------
{    "id": 66,    "rel_type": 87,    "rel_id": 85,    "qty": 33,    "rest_qty": 49,    "version": 60,    "created_at": "2113-12-23T07:27:38.081287198+08:00",    "updated_at": "2141-06-13T18:17:20.30602401+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// MInventory struct is a row record of the m_inventory table in the magazine database
type MInventory struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: false  col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;column:id;type:ubigint;" json:"id"`
	//[ 1] rel_type                                       ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	RelType uint64 `gorm:"column:rel_type;type:ubigint;" json:"rel_type"` // 实体类型
	//[ 2] rel_id                                         ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	RelID uint64 `gorm:"column:rel_id;type:ubigint;" json:"rel_id"` // 实体
	//[ 3] qty                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Qty uint32 `gorm:"column:qty;type:uint;default:0;" json:"qty"`
	//[ 4] rest_qty                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	RestQty uint32 `gorm:"column:rest_qty;type:uint;default:0;" json:"rest_qty"`
	//[ 5] version                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Version uint32 `gorm:"column:version;type:uint;default:0;" json:"version"`
	//[ 6] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [current_timestamp()]
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 7] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [0000-00-00 00:00:00]
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:0000-00-00 00:00:00;" json:"updated_at"`
}

var m_inventoryTableInfo = &TableInfo{
	Name: "m_inventory",
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
			Name:               "rel_type",
			Comment:            `实体类型`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "RelType",
			GoFieldType:        "uint64",
			JSONFieldName:      "rel_type",
			ProtobufFieldName:  "rel_type",
			ProtobufType:       "uint64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "rel_id",
			Comment:            `实体`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "RelID",
			GoFieldType:        "uint64",
			JSONFieldName:      "rel_id",
			ProtobufFieldName:  "rel_id",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "qty",
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
			GoFieldName:        "Qty",
			GoFieldType:        "uint32",
			JSONFieldName:      "qty",
			ProtobufFieldName:  "qty",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "rest_qty",
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
			GoFieldName:        "RestQty",
			GoFieldType:        "uint32",
			JSONFieldName:      "rest_qty",
			ProtobufFieldName:  "rest_qty",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "version",
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
			GoFieldName:        "Version",
			GoFieldType:        "uint32",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MInventory) TableName() string {
	return "m_inventory"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MInventory) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MInventory) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MInventory) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MInventory) TableInfo() *TableInfo {
	return m_inventoryTableInfo
}

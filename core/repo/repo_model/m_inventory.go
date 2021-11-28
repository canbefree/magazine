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
  `id` bigint(20) NOT NULL,
  `qty` varchar(100) NOT NULL DEFAULT '0',
  `rest_qty` varchar(100) NOT NULL DEFAULT '0',
  `version` varchar(100) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='库存总表'

JSON Sample
-------------------------------------
{    "id": 98,    "qty": "pjYtKaywFpvTjnxXpGUjOTsNt",    "rest_qty": "JQoVAcsrZNotOswMbTiOfNtEn",    "version": "ZgGABIZRGbEvtJlxgFevmmhKU",    "created_at": "2146-09-03T22:05:56.619844516+08:00",    "updated_at": "2156-09-10T01:33:42.35436937+08:00"}



*/

// MInventory struct is a row record of the m_inventory table in the magazine database
type MInventory struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id"`
	//[ 1] qty                                            varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [0]
	Qty string `gorm:"column:qty;type:varchar;size:100;default:0;" json:"qty"`
	//[ 2] rest_qty                                       varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [0]
	RestQty string `gorm:"column:rest_qty;type:varchar;size:100;default:0;" json:"rest_qty"`
	//[ 3] version                                        varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: [0]
	Version string `gorm:"column:version;type:varchar;size:100;default:0;" json:"version"`
	//[ 4] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [current_timestamp()]
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 5] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [0000-00-00 00:00:00]
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:0000-00-00 00:00:00;" json:"updated_at"`
}

var m_inventoryTableInfo = &TableInfo{
	Name: "m_inventory",
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
			Name:               "qty",
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
			GoFieldName:        "Qty",
			GoFieldType:        "string",
			JSONFieldName:      "qty",
			ProtobufFieldName:  "qty",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "rest_qty",
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
			GoFieldName:        "RestQty",
			GoFieldType:        "string",
			JSONFieldName:      "rest_qty",
			ProtobufFieldName:  "rest_qty",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "version",
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
			GoFieldName:        "Version",
			GoFieldType:        "string",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
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

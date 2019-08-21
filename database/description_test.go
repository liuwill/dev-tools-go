package database

import (
	"strings"
	"testing"
)

const schemaData = `[
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "pro_calshop",
    "TABLE_NAME": "tb_admin",
    "TABLE_TYPE": "BASE TABLE",
    "ENGINE": "InnoDB",
    "VERSION": 10,
    "ROW_FORMAT": "Compact",
    "TABLE_ROWS": 7,
    "AVG_ROW_LENGTH": 2340,
    "DATA_LENGTH": 16384,
    "MAX_DATA_LENGTH": 0,
    "INDEX_LENGTH": 0,
    "DATA_FREE": 0,
    "AUTO_INCREMENT": 8,
    "CREATE_TIME": "2019-01-16 10:28:50",
    "UPDATE_TIME": null,
    "CHECK_TIME": null,
    "TABLE_COLLATION": "utf8_general_ci",
    "CHECKSUM": null,
    "CREATE_OPTIONS": "",
    "TABLE_COMMENT": "管理员信息表",
    "BLOCK_FORMAT": "Original"
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "pro_calshop",
    "TABLE_NAME": "tb_ads_position",
    "TABLE_TYPE": "BASE TABLE",
    "ENGINE": "InnoDB",
    "VERSION": 10,
    "ROW_FORMAT": "Compact",
    "TABLE_ROWS": 11,
    "AVG_ROW_LENGTH": 1489,
    "DATA_LENGTH": 16384,
    "MAX_DATA_LENGTH": 0,
    "INDEX_LENGTH": 32768,
    "DATA_FREE": 0,
    "AUTO_INCREMENT": 12,
    "CREATE_TIME": "2019-07-23 14:43:14",
    "UPDATE_TIME": null,
    "CHECK_TIME": null,
    "TABLE_COLLATION": "utf8_general_ci",
    "CHECKSUM": null,
    "CREATE_OPTIONS": "",
    "TABLE_COMMENT": "系统广告位表",
    "BLOCK_FORMAT": "Original"
  }
]`

func Test_BuildColumnConfig(t *testing.T) {
	tableRows := BuildColumnConfig(RAW_SCHEMA_CONFIG)
	count := strings.Count(RAW_SCHEMA_CONFIG, "column")

	if len(tableRows) != count {
		t.Error("Test BuildColumnConfig Fail", len(tableRows), count)
	}

	t.Log("Test BuildColumnConfig Success")
}

func Test_ParseTableContent(t *testing.T) {
	tableRows := ParseTableContent(RAW_SCHEMA_CONFIG)
	count := strings.Count(RAW_SCHEMA_CONFIG, "column")

	if len(tableRows) != count {
		t.Error("Test BuildColumnConfig Fail", len(tableRows), count)
	}
	t.Log("Test BuildColumnConfig Success")
}

func Test_BuildMarkedTable(t *testing.T) {
	tableConfig := BuildColumnConfig(RAW_SCHEMA_CONFIG)
	tableData := ParseTableContent(schemaData)

	markedLines := BuildMarkedTable(tableConfig, tableData)
	if len(markedLines) != len(tableData)+2 {
		t.Error("Test BuildMarkedTable Fail", len(markedLines), len(tableData))
	}

	t.Log("Test BuildMarkedTable Success")
}

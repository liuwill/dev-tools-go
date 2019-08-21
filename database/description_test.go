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
const RAW_COLUMN_DATA = `[
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "id",
    "ORDINAL_POSITION": 1,
    "COLUMN_DEFAULT": null,
    "IS_NULLABLE": "NO",
    "DATA_TYPE": "int",
    "CHARACTER_MAXIMUM_LENGTH": null,
    "CHARACTER_OCTET_LENGTH": null,
    "NUMERIC_PRECISION": 10,
    "NUMERIC_SCALE": 0,
    "DATETIME_PRECISION": null,
    "CHARACTER_SET_NAME": null,
    "COLLATION_NAME": null,
    "COLUMN_TYPE": "int(10) unsigned",
    "COLUMN_KEY": "PRI",
    "EXTRA": "auto_increment",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": ""
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "task_id",
    "ORDINAL_POSITION": 2,
    "COLUMN_DEFAULT": "",
    "IS_NULLABLE": "NO",
    "DATA_TYPE": "varchar",
    "CHARACTER_MAXIMUM_LENGTH": 64,
    "CHARACTER_OCTET_LENGTH": 192,
    "NUMERIC_PRECISION": null,
    "NUMERIC_SCALE": null,
    "DATETIME_PRECISION": null,
    "CHARACTER_SET_NAME": "utf8",
    "COLLATION_NAME": "utf8_general_ci",
    "COLUMN_TYPE": "varchar(64)",
    "COLUMN_KEY": "MUL",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": "定时任务id"
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_admin",
    "COLUMN_NAME": "id",
    "ORDINAL_POSITION": 1,
    "COLUMN_DEFAULT": null,
    "IS_NULLABLE": "NO",
    "DATA_TYPE": "int",
    "CHARACTER_MAXIMUM_LENGTH": null,
    "CHARACTER_OCTET_LENGTH": null,
    "NUMERIC_PRECISION": 10,
    "NUMERIC_SCALE": 0,
    "DATETIME_PRECISION": null,
    "CHARACTER_SET_NAME": null,
    "COLLATION_NAME": null,
    "COLUMN_TYPE": "int(10) unsigned",
    "COLUMN_KEY": "PRI",
    "EXTRA": "auto_increment",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": ""
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_admin",
    "COLUMN_NAME": "mobile",
    "ORDINAL_POSITION": 2,
    "COLUMN_DEFAULT": null,
    "IS_NULLABLE": "YES",
    "DATA_TYPE": "varchar",
    "CHARACTER_MAXIMUM_LENGTH": 32,
    "CHARACTER_OCTET_LENGTH": 96,
    "NUMERIC_PRECISION": null,
    "NUMERIC_SCALE": null,
    "DATETIME_PRECISION": null,
    "CHARACTER_SET_NAME": "utf8",
    "COLLATION_NAME": "utf8_general_ci",
    "COLUMN_TYPE": "varchar(32)",
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": "管理员手机"
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_admin",
    "COLUMN_NAME": "email",
    "ORDINAL_POSITION": 3,
    "COLUMN_DEFAULT": null,
    "IS_NULLABLE": "YES",
    "DATA_TYPE": "varchar",
    "CHARACTER_MAXIMUM_LENGTH": 512,
    "CHARACTER_OCTET_LENGTH": 1536,
    "NUMERIC_PRECISION": null,
    "NUMERIC_SCALE": null,
    "DATETIME_PRECISION": null,
    "CHARACTER_SET_NAME": "utf8",
    "COLLATION_NAME": "utf8_general_ci",
    "COLUMN_TYPE": "varchar(512)",
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": "管理员邮箱"
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

	header, splitter := BuildMarkedTableHeader(tableConfig)
	headerList := []string{header, splitter}

	markedLines := BuildMarkedTable(tableConfig, tableData)
	if len(markedLines) != len(tableData)+len(headerList) {
		t.Error("Test BuildMarkedTable Fail", len(markedLines), len(tableData))
	}

	for i, headerItem := range headerList {
		if headerItem != markedLines[i] {
			t.Error("Test BuildMarkedTable Header Fail", len(markedLines), len(tableData))
		}
	}

	t.Log("Test BuildMarkedTable Success")
}

func Test_GroupRowByTable(t *testing.T) {
	tableData := ParseTableContent(RAW_COLUMN_DATA)
	targetMap := GroupRowByTable(tableData)

	dataMeta := map[string]int{
		"tb_admin":      3,
		"tb_timer_task": 2,
	}

	if len(targetMap) != len(dataMeta) {
		t.Error("Test GroupRowByTable Len Fail", len(targetMap))
	}

	for k, v := range dataMeta {
		if len(targetMap[k]) != v {
			t.Error("Test GroupRowByTable Fail", len(targetMap))
		}
	}

	t.Log("Test GroupRowByTable Success")
}

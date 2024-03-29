package comment

import (
	"strings"
	"testing"
)

const mockContent = `[
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
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "title",
    "ORDINAL_POSITION": 3,
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
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": "定时任务名称"
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "cron_time",
    "ORDINAL_POSITION": 4,
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
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": "定时时间配置"
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "status",
    "ORDINAL_POSITION": 5,
    "COLUMN_DEFAULT": "1",
    "IS_NULLABLE": "NO",
    "DATA_TYPE": "tinyint",
    "CHARACTER_MAXIMUM_LENGTH": null,
    "CHARACTER_OCTET_LENGTH": null,
    "NUMERIC_PRECISION": 3,
    "NUMERIC_SCALE": 0,
    "DATETIME_PRECISION": null,
    "CHARACTER_SET_NAME": null,
    "COLLATION_NAME": null,
    "COLUMN_TYPE": "tinyint(2)",
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": "状态： 0 不执行 1 执行"
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "start_timestamp",
    "ORDINAL_POSITION": 6,
    "COLUMN_DEFAULT": "0000-00-00 00:00:00",
    "IS_NULLABLE": "NO",
    "DATA_TYPE": "timestamp",
    "CHARACTER_MAXIMUM_LENGTH": null,
    "CHARACTER_OCTET_LENGTH": null,
    "NUMERIC_PRECISION": null,
    "NUMERIC_SCALE": null,
    "DATETIME_PRECISION": 0,
    "CHARACTER_SET_NAME": null,
    "COLLATION_NAME": null,
    "COLUMN_TYPE": "timestamp",
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": "开始执行的时间戳"
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "end_timestamp",
    "ORDINAL_POSITION": 7,
    "COLUMN_DEFAULT": "0000-00-00 00:00:00",
    "IS_NULLABLE": "NO",
    "DATA_TYPE": "timestamp",
    "CHARACTER_MAXIMUM_LENGTH": null,
    "CHARACTER_OCTET_LENGTH": null,
    "NUMERIC_PRECISION": null,
    "NUMERIC_SCALE": null,
    "DATETIME_PRECISION": 0,
    "CHARACTER_SET_NAME": null,
    "COLLATION_NAME": null,
    "COLUMN_TYPE": "timestamp",
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": "结束执行的时间戳"
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "run_duration",
    "ORDINAL_POSITION": 8,
    "COLUMN_DEFAULT": "0",
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
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": "执行时间，毫秒"
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "last_run_time",
    "ORDINAL_POSITION": 9,
    "COLUMN_DEFAULT": "0000-00-00 00:00:00",
    "IS_NULLABLE": "NO",
    "DATA_TYPE": "datetime",
    "CHARACTER_MAXIMUM_LENGTH": null,
    "CHARACTER_OCTET_LENGTH": null,
    "NUMERIC_PRECISION": null,
    "NUMERIC_SCALE": null,
    "DATETIME_PRECISION": 0,
    "CHARACTER_SET_NAME": null,
    "COLLATION_NAME": null,
    "COLUMN_TYPE": "datetime",
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": "上次执行时间"
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "create_time",
    "ORDINAL_POSITION": 10,
    "COLUMN_DEFAULT": "CURRENT_TIMESTAMP",
    "IS_NULLABLE": "NO",
    "DATA_TYPE": "datetime",
    "CHARACTER_MAXIMUM_LENGTH": null,
    "CHARACTER_OCTET_LENGTH": null,
    "NUMERIC_PRECISION": null,
    "NUMERIC_SCALE": null,
    "DATETIME_PRECISION": 0,
    "CHARACTER_SET_NAME": null,
    "COLLATION_NAME": null,
    "COLUMN_TYPE": "datetime",
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": ""
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "update_time",
    "ORDINAL_POSITION": 11,
    "COLUMN_DEFAULT": "CURRENT_TIMESTAMP",
    "IS_NULLABLE": "NO",
    "DATA_TYPE": "timestamp",
    "CHARACTER_MAXIMUM_LENGTH": null,
    "CHARACTER_OCTET_LENGTH": null,
    "NUMERIC_PRECISION": null,
    "NUMERIC_SCALE": null,
    "DATETIME_PRECISION": 0,
    "CHARACTER_SET_NAME": null,
    "COLLATION_NAME": null,
    "COLUMN_TYPE": "timestamp",
    "COLUMN_KEY": "",
    "EXTRA": "on update CURRENT_TIMESTAMP",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": ""
  },
  {
    "TABLE_CATALOG": "def",
    "TABLE_SCHEMA": "calshop",
    "TABLE_NAME": "tb_timer_task",
    "COLUMN_NAME": "deleted",
    "ORDINAL_POSITION": 12,
    "COLUMN_DEFAULT": "0",
    "IS_NULLABLE": "NO",
    "DATA_TYPE": "tinyint",
    "CHARACTER_MAXIMUM_LENGTH": null,
    "CHARACTER_OCTET_LENGTH": null,
    "NUMERIC_PRECISION": 3,
    "NUMERIC_SCALE": 0,
    "DATETIME_PRECISION": null,
    "CHARACTER_SET_NAME": null,
    "COLLATION_NAME": null,
    "COLUMN_TYPE": "tinyint(2) unsigned",
    "COLUMN_KEY": "",
    "EXTRA": "",
    "PRIVILEGES": "select,insert,update,references",
    "COLUMN_COMMENT": ""
  }
]`

func Test_PickTableRow(t *testing.T) {
	tableRows := PickTableRow(mockContent)
	count := strings.Count(mockContent, "TABLE_NAME")

	if len(tableRows) != count {
		t.Error("Test PickTableRow Fail", len(tableRows), count)
	}
	t.Log("Test PickTableRow Success")
}

func Test_GenerateComment(t *testing.T) {
	tableRows := PickTableRow(mockContent)
	apiConfig := ApiConfig{
		Method:   "get",
		Path:     "/assistance/sku/list",
		Name:     "获取一键补货活动配置列表",
		Module:   "Activity",
		Function: "listAssistSkuPage",
		Type:     "assistSkuData",
	}

	target := GenerateComment(apiConfig, tableRows)
	for _, row := range tableRows {
		if !strings.Contains(target, row.ColumnName) {
			t.Error("Test PickTableRow:ColumnName Fail", row.ColumnName)
			break
		}
	}

	t.Log("Test GenerateComment Success")
}

func Test_BuildColumnStructDefine(t *testing.T) {
	header := "TABLE_CATALOG	TABLE_SCHEMA	TABLE_NAME	COLUMN_NAME	ORDINAL_POSITION	COLUMN_DEFAULT	IS_NULLABLE	DATA_TYPE	CHARACTER_MAXIMUM_LENGTH	CHARACTER_OCTET_LENGTH	NUMERIC_PRECISION	NUMERIC_SCALE	DATETIME_PRECISION	CHARACTER_SET_NAME	COLLATION_NAME	COLUMN_TYPE	COLUMN_KEY	EXTRA	PRIVILEGES	COLUMN_COMMENT"
	result := buildColumnStructDefine(header)
	words := strings.Split(header, "\t")
	if strings.Count(result, "json") != len(words) {
		t.Error("Test BuildColumnStructDefine json Fail")
	}
	t.Log("Test BuildColumnStructDefine Success")
}

func Test_ToBigCamelCaseWord(t *testing.T) {
	camelCase := [][]string{
		[]string{"Table", "Table"},
		[]string{"TABLE", "Table"},
		[]string{"1st", "1st"},
		[]string{"", ""},
	}

	for _, caseData := range camelCase {
		camel := toBigCamelCaseWord(caseData[0])
		if camel != caseData[1] {
			t.Error("Test ToBigCamelCaseWord json Fail", camel, caseData)
		}
	}

	t.Log("Test ToBigCamelCaseWord Success")
}

func Test_ToCamelCase(t *testing.T) {
	camelCase := [][]string{
		[]string{"data_Table", "DataTable"},
		[]string{"TABLE_test", "tableTest"},
	}
	bigCase := []bool{
		true, false,
	}

	for i, caseData := range camelCase {
		isBig := bigCase[i]
		word := toCamelCase(caseData[0], isBig)
		if word != caseData[1] {
			t.Error("Test ToCamelCase Fail", word, caseData)
		}
	}

	t.Log("Test ToCamelCase Success")
}

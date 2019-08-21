package database

import (
	"encoding/json"
	"strings"
)

const (
	RAW_SCHEMA_CONFIG = `[
		{ "column": "table_name", "name": "表名", "anchor": true },
		{ "column": "table_comment", "name": "描述" },
		{ "column": "engine", "name": "存储引擎" }
	]`
	RAW_COLUMN_CONFIG = `[
		{ "column": "column_name", "name": "列名" },
		{ "column": "column_comment", "name": "描述" },
		{ "column": "column_type", "name": "类型" },
		{ "column": "is_nullable", "name": "允许为空" },
		{ "column": "column_key", "name": "键类型" },
		{ "column": "column_default", "name": "默认值" },
		{ "column": "extra", "name": "额外属性" }
	]`
	RAW_INDEX_CONFIG = `[
		{ "column": "index_name", "name": "索引名称" },
		{ "column": "seq_in_index", "name": "键位置" },
		{ "column": "column_name", "name": "列名" },
		{ "column": "non_unique", "name": "是否唯一" },
		{ "column": "cardinality", "name": "基数" }
	]`
)

type ColumnConfig struct {
	Column   string `json:"column"`
	Name     string `json:"name"`
	IsAnchor bool   `json:"anchor"`
}

func BuildColumnConfig(jsonStr string) []ColumnConfig {
	columnConfig := []ColumnConfig{}
	json.Unmarshal([]byte(jsonStr), &columnConfig)
	return columnConfig
}

func ParseTableContent(jsonStr string) []map[string]interface{} {
	var tableData []map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &tableData)
	return tableData
}

func JoinMarkedLine(columns []string) string {
	line := strings.Join(columns, " | ")
	return "| " + line + " |"
}

func BuildMarkedTableHeader(tableConfig []ColumnConfig) (string, string) {
	headerMeta := make([]string, len(tableConfig))
	splitterMeta := make([]string, len(tableConfig))

	for i, item := range tableConfig {
		headerMeta[i] = item.Name
		splitterMeta[i] = "--------:"
	}

	header := JoinMarkedLine(headerMeta)
	splitter := JoinMarkedLine(splitterMeta)
	return header, splitter
}

func FetchTableColumnValue(config ColumnConfig, rowData map[string]interface{}) string {
	key := strings.ToUpper(config.Column)
	columnVal := rowData[key].(string)
	if config.IsAnchor {
		columnVal = "[" + columnVal + "](#" + columnVal + ")"
	}
	return columnVal
}

func BuildMarkedTable(tableConfig []ColumnConfig, tableData []map[string]interface{}) []string {
	header, splitter := BuildMarkedTableHeader(tableConfig)
	lines := make([]string, len(tableData)+2)

	pos := 0
	tableHeader := []string{
		header, splitter,
	}
	for _, val := range tableHeader {
		lines[pos] = val
		pos++
	}

	for _, current := range tableData {
		lineItems := make([]string, len(tableConfig))
		for j, config := range tableConfig {
			lineItems[j] = FetchTableColumnValue(config, current)
		}
		lines[pos] = JoinMarkedLine(lineItems)
		pos++
	}

	return lines
}

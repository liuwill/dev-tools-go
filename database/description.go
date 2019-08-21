package database

import "encoding/json"

const (
	RAW_SCHEMA_CONFIG = `[
		{ "column": "table_name", "name": "表名" },
		{ "column": "table_comment", "name": "描述" },
		{ "column": "engine", "name": "存储引擎" }
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

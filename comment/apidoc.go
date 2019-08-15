package comment

import "encoding/json"

/*
 * Row From information_schema.COLUMNS
 * TABLE_CATALOG	TABLE_SCHEMA	TABLE_NAME	COLUMN_NAME	ORDINAL_POSITION	COLUMN_DEFAULT	IS_NULLABLE	DATA_TYPE	CHARACTER_MAXIMUM_LENGTH	CHARACTER_OCTET_LENGTH	NUMERIC_PRECISION	NUMERIC_SCALE	DATETIME_PRECISION	CHARACTER_SET_NAME	COLLATION_NAME	COLUMN_TYPE	COLUMN_KEY	EXTRA	PRIVILEGES	COLUMN_COMMENT
 * const toCamel = (item) => {
   	const list = item.toLowerCase().split('_');
   	return list.reduce((result, item, index) => {
   		result += item[0].toUpperCase() + item.substr(1)
   		return result
   	}, '')
   }
 * str.split('\t').map(item => `${toCamel(item)} string \`json:"${item}"\``)
*/
type ApiConfig struct {
	method   string
	path     string
	name     string
	module   string
	function string
}

type TableRow struct {
	TableCatalog           string `json:"TABLE_CATALOG"`
	TableSchema            string `json:"TABLE_SCHEMA"`
	TableName              string `json:"TABLE_NAME"`
	ColumnName             string `json:"COLUMN_NAME"`
	OrdinalPosition        string `json:"ORDINAL_POSITION"`
	ColumnDefault          string `json:"COLUMN_DEFAULT"`
	IsNullable             string `json:"IS_NULLABLE"`
	DataType               string `json:"DATA_TYPE"`
	CharacterMaximumLength string `json:"CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   string `json:"CHARACTER_OCTET_LENGTH"`
	NumericPrecision       string `json:"NUMERIC_PRECISION"`
	NumericScale           string `json:"NUMERIC_SCALE"`
	DatetimePrecision      string `json:"DATETIME_PRECISION"`
	CharacterSetName       string `json:"CHARACTER_SET_NAME"`
	CollationName          string `json:"COLLATION_NAME"`
	ColumnType             string `json:"COLUMN_TYPE"`
	ColumnKey              string `json:"COLUMN_KEY"`
	Extra                  string `json:"EXTRA"`
	Privileges             string `json:"PRIVILEGES"`
	ColumnComment          string `json:"COLUMN_COMMENT"`
}

func pickTableRow(jsonStr string) []TableRow {
	tableRow := []TableRow{}
	json.Unmarshal([]byte(jsonStr), &tableRow)
	return tableRow
}

func generateComment(apiConfig ApiConfig, dataName string, tableData []TableRow) string {
	return ""
}

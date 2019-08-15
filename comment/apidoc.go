package comment

import (
	"bytes"
	"encoding/json"
	"html/template"
	"strings"
)

/*
 * Row From information_schema.COLUMNS
 * TABLE_CATALOG	TABLE_SCHEMA	TABLE_NAME	COLUMN_NAME	ORDINAL_POSITION	COLUMN_DEFAULT	IS_NULLABLE	DATA_TYPE	CHARACTER_MAXIMUM_LENGTH	CHARACTER_OCTET_LENGTH	NUMERIC_PRECISION	NUMERIC_SCALE	DATETIME_PRECISION	CHARACTER_SET_NAME	COLLATION_NAME	COLUMN_TYPE	COLUMN_KEY	EXTRA	PRIVILEGES	COLUMN_COMMENT
 * const toCamel = (item) => {
   	const list = item.toLowerCase().split("_")
   	return list.reduce((result, item, index) => {
   		result += item[0].toUpperCase() + item.substr(1)
   		return result
   	}, "")
   }
 * str.split("\t").map(item => `${toCamel(item)} string \`json:"${item}"\``)
*/
type ApiConfig struct {
	Method   string
	Path     string
	Name     string
	Module   string
	Function string
	Type     string
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

var CommentTemplate = `
/**
 * @api {{"{"}}{{.Method}}{{"}"}} {{.Path}} {{.Name}}
 * @apiName {{.Function}}
 * @apiGroup {{.Module}}
 *
 * @apiHeader {String} Authentication 登录获得的token.
 *
 * @apiParam (query) {Number} pn 页码.
 * @apiParam (query) {Number} ps 页面尺寸.
 *
 * @apiSuccess {Number} status 服务器状态.
 * @apiSuccess {MissionDataResult} data 返回数据.
 *
 * @apiSuccess ({{.Type}}Result) {{"{"}}{{.Type}}[]{{"}"}} list 任务数据列表.
 * @apiSuccess ({{.Type}}Result) {Number} total 数量.
 *
${commentStr}
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 0,
 *       "data": {
 *          total: 11,
 *          list: [{
${responseStr}
 *         }]
 *       }
 *     }
 */
`

var defaultComment = map[string]string{
	"id":          "主键id",
	"create_time": "创建时间",
	"update_time": "更新时间",
}

var typeConfig = map[string]string{
	"int":       "Number",
	"varchar":   "String",
	"datetime":  "Date",
	"tinyint":   "Number",
	"timestamp": "Date",
}

func toBigCamelCaseWord(word string) string {
	if len(word) <= 0 {
		return word
	}

	word = strings.ToLower(word)
	if word[0] < 97 || word[0] > 122 {
		return word
	}

	firstLetter := string(word[0] - 32)
	return firstLetter + string(word[1:])
}

func toBigCamelCase(sentence string) string {
	words := strings.Split(sentence, "_")
	target := make([]string, len(words))

	for i, v := range words {
		target[i] = toBigCamelCaseWord(v)
	}
	return strings.Join(target, "")
}

func buildColumnStructDefine(header string) string {
	columns := strings.Split(header, "\t")
	lines := make([]string, len(columns))
	for i, v := range columns {
		word := toBigCamelCase(v)
		lines[i] = "\t" + word + " string `json:\"" + v + "\"`"
	}
	lineStr := strings.Join(lines, "\n")
	return "type TableRow struct {\n" + lineStr + "\n}"
}

func PickTableRow(jsonStr string) []TableRow {
	tableRow := []TableRow{}
	json.Unmarshal([]byte(jsonStr), &tableRow)
	return tableRow
}

func buildCommentLine(dataName string, row TableRow) string {
	columnType := "Number"
	if v, ok := typeConfig[row.DataType]; ok {
		columnType = v
	}
	columnName := row.ColumnName
	columnComment := row.ColumnComment
	if v, ok := defaultComment[columnName]; ok && len(columnComment) <= 0 {
		columnComment = v
	}
	return ` * @apiSuccess (` + dataName + `) {` + columnType + `} ` + columnName + ` ` + columnComment + `.`
}

func buildResponseLine(row TableRow) string {
	// columnType := row.DataType
	columnValue := ""
	// if (mockFactory[columnType]) {
	//   columnValue = mockFactory[columnType].apply(chance)
	// }

	columnName := row.ColumnName
	return ` *            "` + columnName + `": "` + columnValue + `",`
}

func GenerateComment(apiConfig ApiConfig, tableData []TableRow) string {
	dataName := apiConfig.Type
	commentList := make([]string, len(tableData))
	responseList := make([]string, len(tableData))
	for i, item := range tableData {
		commentList[i] = buildCommentLine(dataName, item)
		responseList[i] = buildResponseLine(item)
	}
	commentStr := strings.Join(commentList, "\n")
	responseStr := strings.Join(responseList, "\n")

	CommentTemplate = strings.Replace(CommentTemplate, "${commentStr}", commentStr, -1)
	CommentTemplate = strings.Replace(CommentTemplate, "${responseStr}", responseStr, -1)

	runner, _ := template.New("comment").Parse(CommentTemplate) //建立一个模板
	// if err != nil {
	// 	panic(err)
	// }
	buf := new(bytes.Buffer)
	runner.Execute(buf, apiConfig)
	return buf.String()
}

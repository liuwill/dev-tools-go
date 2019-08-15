package comment

type ApiConfig struct {
	method   string
	path     string
	name     string
	module   string
	function string
}

type TableRow struct {
}

func generateComment(apiConfig ApiConfig, dataName string, tableData []TableRow) string {
	return ""
}

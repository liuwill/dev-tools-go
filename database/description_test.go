package database

import (
	"strings"
	"testing"
)

func Test_BuildColumnConfig(t *testing.T) {
	tableRows := BuildColumnConfig(RAW_SCHEMA_CONFIG)
	count := strings.Count(RAW_SCHEMA_CONFIG, "column")

	if len(tableRows) != count {
		t.Error("Test BuildColumnConfig Fail", len(tableRows), count)
	}
	t.Log("Test BuildColumnConfig Success")
}

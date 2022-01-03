package web

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func TestWorkspaceNameUpdate(t *testing.T) {
	if WorkspaceNameUpdate("testing@xenonstack.com", "1") != nil {
		t.Error("test case fail")
	}
}

package web

import "testing"

func TestCheckWorkspaceUser(t *testing.T) {
	err := checkWorkspaceUser("rahul@xenonstack.com", "rahul-172", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQxMjExMTYwLCJpZCI6MTcyLCJuYW1lIjoiUmFodWwgS3VtYXIiLCJvcmlnX2lhdCI6MTY0MTIwOTM2MCwic3lzX3JvbGUiOiJ1c2VyIn0.fikFTxfil6gxfeR7dpvdiL4-EQ21yjHL7e3LNtZnkBQ")
	if err != nil {
		t.Error("test case fail")
	}
	err = checkWorkspaceUser("rahul@xenonstack.com", "rahul-172", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQxMjExMTYwLCJpZCI6MTcyLCJuYW1lIjoiUmFodWwgS3VtYXIiLCJvcmlnX2lhdCI6MTY0MTIwOTM2MCwic3lzX3JvbGUiOiJ1c2VyIn0.fikFTxfil6gxfeR7dpvdiL4-EQ21yjHL7e3LNtZnkBQ")
	if err == nil {
		t.Error("test case fail")
	}
}

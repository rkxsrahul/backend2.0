package web

import (
	"log"
	"os"
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/config"
	"git.xenonstack.com/util/continuous-security-backend/src/database"
	"github.com/jinzhu/gorm"
)

func init() {

	os.Remove(os.Getenv("HOME") + "/account-testing.db")
	db, err := gorm.Open("sqlite3", os.Getenv("HOME")+"/account-testing.db")
	if err != nil {
		log.Println(err)
		log.Println("Exit")
		os.Exit(1)
	}
	config.DB = db

	//create table
	database.CreateDBTablesIfNotExists()
	var workspace, workspace2 database.RequestInfo
	workspace.ID = 1
	workspace.Email = "testing@xenonstack.com"
	workspace.UUID = "c6o4qf77dscabskn9v6g"
	workspace.Workspace = "rahul-172"
	workspace.RepoLang = "Node JS"

	db.Create(&workspace)

	workspace2.ID = 2
	workspace2.Email = "testing2@xenonstack.com"
	workspace2.Workspace = "rahul-172"
	workspace2.RepoLang = "Node JS"

	db.Create(&workspace2)

	scan := database.ScanResult{}
	scan.UUID = "c6o4qf77dscabskn9v6g"
	scan.Result = `{"description":"HTTP Expect-CT header is not implemented properly","header":"Expect-CT","heading":"Expect-CT","id":"expect-ct","impact":"PASS","message":"implemented\ntesting","secure":true}`
	scan.Status = true
	db.Create(&scan)
	scan2 := database.ScanResult{}
	scan2.UUID = "c762uhv7dscd50dgf860"
	scan2.Result = `{"description":"HTTP Expect-CT header is not implemented properly","header":"Expect-CT","heading":"Expect-CT","id":"expect-ct","impact":"PASS","message":"implemented\ntesting","secure":true}`
	scan2.Status = false
	db.Create(&scan2)
	scan3 := database.ScanResult{}
	scan3.UUID = "c761jqf7dscd50dgf820"
	scan3.Result = "*****"
	scan3.Status = false
	db.Create(&scan3)
}

func TestIntegration(t *testing.T) {
	code, _ := Integration("rahul@xenonstack.com", "rahul-172", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQxMjEyMDczLCJpZCI6MTcyLCJuYW1lIjoiUmFodWwgS3VtYXIiLCJvcmlnX2lhdCI6MTY0MTIxMDI3Mywic3lzX3JvbGUiOiJ1c2VyIn0.vMyXfvk15ffA0bKvhfEP9S55SuoRY-f4Bp-k79RN3nA")
	log.Println(code)
	code, _ = Integration("rahul@xenonstack.com", "rahul", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQxMjEyMDczLCJpZCI6MTcyLCJuYW1lIjoiUmFodWwgS3VtYXIiLCJvcmlnX2lhdCI6MTY0MTIxMDI3Mywic3lzX3JvbGUiOiJ1c2VyIn0.vMyXfvk15ffA0bKvhfEP9S55SuoRY-f4Bp-k79RN3nA")

}

func TestScanInformation(t *testing.T) {
	ScanInformation("c6o4qf77dscabskn9v6g", "rahul-172", "rahul@xenonstack.com", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQxMjE1MzQyLCJpZCI6MTcyLCJuYW1lIjoiUmFodWwgS3VtYXIiLCJvcmlnX2lhdCI6MTY0MTIxMzU0Miwic3lzX3JvbGUiOiJ1c2VyIn0.nyWlkAVmRQ2t_IGwgewCMjeyonVoVcD78AG0WvYk9zU")
	ScanInformation("c6o4qf77dscabskn9v6g", "rahul-172", "rahul@xenonstack.com", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQxMjEzOTI3LCJpZCI6MTcyLCJuYW1lIjoiUmFodWwgS3VtYXIiLCJvcmlnX2lhdCI6MTY0MTIxMjEyNywic3lzX3JvbGUiOiJ1c2VyIn0.jUti8Lbf4jyrYRbjmc11_9c6PRSuUstg01qJaGiPIwQ")

}

func TestGitResult(t *testing.T) {
	info := database.RequestInfo{}
	info.UUID = "c762uhv7dscd50dgf860"
	gitResult(info)
	info = database.RequestInfo{}
	info.UUID = "c762uhv7dscd50dgf86"
	gitResult(info)
	info = database.RequestInfo{}
	info.UUID = "c761jqf7dscd50dgf820"
	gitResult(info)
}

func TestwebsiteResult() {

}

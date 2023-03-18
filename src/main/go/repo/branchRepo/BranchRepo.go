package branchRepo

import (
	config "github.com/SbFinanceFbd/BranchRead/src/main/go/config"
	l "github.com/SbFinanceFbd/golib/logger"
	"github.com/SbFinanceFbd/golib/sql"
)

func GetBranchs(start, size int64) (rowData [][]string, err error) {
	action := "::GetBranchs::"
	rows, err := config.Db.Query("SELECT id,name,address,place,taluka,district,state,pin_code,branch_txn_tbl_name,is_approved,status,created_at,updated_at,created_by,updated_by,remarks FROM "+config.AppSchema.MustGetString("table.branch")+" order by id desc limit $1 offset $2", size, start)
	defer sql.Close(rows)
	if err != nil {
		l.Error.Printf("%v Error-> %v\n", action, err.Error())
		return
	}

	_, rowData, err = sql.Scan(rows)
	if err != nil {
		l.Error.Printf("Query Error-> : %v \n", err.Error())
		return
	}
	return
}

func GetBranchByBranchId(branchId int64) (rowData [][]string, err error) {
	action := "::GetBranchByBranchId::"
	rows, err := config.Db.Query("SELECT id,name,address,place,taluka,district,state,pin_code,branch_txn_tbl_name,is_approved,status,created_at,updated_at,created_by,updated_by,remarks FROM "+config.AppSchema.MustGetString("table.branch")+" where id=$1", branchId)
	defer sql.Close(rows)
	if err != nil {
		l.Error.Printf("%v Error-> %v\n", action, err.Error())
		return
	}

	_, rowData, err = sql.Scan(rows)
	if err != nil {
		l.Error.Printf("Query Error-> : %v \n", err.Error())
		return
	}
	return
}

func GetBranchByBranchName(branchName string) (rowData [][]string, err error) {
	action := "::GetBranchByBranchName::"
	rows, err := config.Db.Query("SELECT id,name,address,place,taluka,district,state,pin_code,branch_txn_tbl_name,is_approved,status,created_at,updated_at,created_by,updated_by,remarks FROM "+config.AppSchema.MustGetString("table.branch")+" where name=$1", branchName)
	defer sql.Close(rows)
	if err != nil {
		l.Error.Printf("%v Error-> %v\n", action, err.Error())
		return
	}

	_, rowData, err = sql.Scan(rows)
	if err != nil {
		l.Error.Printf("Query Error-> : %v \n", err.Error())
		return
	}
	return
}

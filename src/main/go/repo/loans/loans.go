package loans

import (
	config "github.com/SbFinanceFbd/BranchRead/src/main/go/config"
	"github.com/SbFinanceFbd/BranchRead/src/main/go/model"
	l "github.com/SbFinanceFbd/golib/logger"
)

func Insert(loanSanctionRequest model.LoanSanctionRequest) (err error) {
	action := "::Insert Loans::"
	_, err = config.Db.Query("INSERT INTO "+config.AppSchema.MustGetString("table.loans")+" (loan_number,relative_name1,relative_name2,relative_num1,relative_num2,amount,date,referred_by,loan_collect_tbl_name,penalty_tbl_name,penalty_date,kyc_id,loa_pack_id,pigmy_num,created_by,remarks) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)",
		loanSanctionRequest.LoanNum, loanSanctionRequest.RelativeName1, loanSanctionRequest.RelativeName2, loanSanctionRequest.RelativeNum1, loanSanctionRequest.RelativeNum2, loanSanctionRequest.Amount, loanSanctionRequest.Date, loanSanctionRequest.ReferredBy,
		loanSanctionRequest.LoanCollectTblName, loanSanctionRequest.PenaltyTblName, loanSanctionRequest.PenaltyDate, loanSanctionRequest.KycId,
		loanSanctionRequest.LoaPackId, loanSanctionRequest.PigmyNum, loanSanctionRequest.CreatedBy, loanSanctionRequest.Remarks)
	if err != nil {
		l.Error.Printf("%v Error-> %v\n", action, err.Error())
		return
	}
	// l.Debug.Printf("%v DBInsert=> [ Data-> sysUserId = %v , uniqueToken = %v , expiredAt = %v , isValid = %v ]\n", action, sysUserId, uniqueToken, expiredAt, 1)
	return
}

// func Insert() (err error) {
// 	action := "::Insert Loans::"
// 	_, err = config.Db.Query("INSERT INTO "+config.AppSchema.MustGetString("table.loans")+" (sysuser_id,unique_token,expired_at,is_valid) values($1,$2,$3,'Y')", sysUserId, uniqueToken, expiredAt)
// 	if err != nil {
// 		l.Error.Printf("%v Error-> %v\n", action, err.Error())
// 		return
// 	}
// 	l.Debug.Printf("%v DBInsert=> [ Data-> sysUserId = %v , uniqueToken = %v , expiredAt = %v , isValid = %v ]\n", action, sysUserId, uniqueToken, expiredAt, 1)
// 	return
// }

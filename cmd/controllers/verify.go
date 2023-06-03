// oreste v0.1


package controllers

import (
	"strings"

	"github.com/alpkeskin/oreste/cmd/utils"
)

func VerifyEmail(email string) {
	defer utils.ProgressBar.Add(10)
	ret, err := utils.Verifier.Verify(email)
	if err != nil {
		utils.Verify_result.Err = err
	}
	if !ret.Syntax.Valid {
		utils.Verify_result.IsVerified = false
	} else {
		domain := strings.Split(email, "@")[1]
		utils.Verify_result.IsVerified = true
		if !utils.Verifier.IsDisposable(domain) {
			utils.Verify_result.IsDisposable = false
		} else {
			utils.Verify_result.IsDisposable = true
		}
	}
}

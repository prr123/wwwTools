package etcpwd


import (
	"testing"
)

func TestLoadEtcPwd(t *testing.T) {

	pwdList, err := LoadEtcPwd()
	if err != nil {t.Errorf("cannot load etcpwd: %v", err)}
	if len(pwdList) < 1 {t.Errorf("no entries in etcpwd!")}
	PrintPwdList(pwdList)
}

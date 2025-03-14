// etcpwd.go
//
// original from https://github.com/astromechza/etcpwdparse/blob/master/pwd.go
//
// modifications
// - allow direct address to 
//


package etcpwd

import (
	"fmt"
	"os"
)


// EtcPasswdEntry is a parsed line from the etc passwd file. It contains all 7 parts of the structure.
// Remember that the password field is encrypted or refers to an item in an alternative authentication scheme.
type PwdEntry struct {
    Username string
    Password string
    Uid      int
    Gid      int
    Info     string
    Home     string
    Shell    string
}

func LoadEtcPwd() (pwdList []PwdEntry, err error) {

	pwdB, err := os.ReadFile("/etc/passwd")
	if err!= nil {return nil, fmt.Errorf("pwd file not found: %v!", err)}

	//parse pwdB
	istate := 0
	entry:= -1
	ist := -1
	user := PwdEntry{}
	for i:=0; i< len(pwdB); i++ {
		ch := pwdB[i]
		switch istate {
			// new entry
			case 0:
				entry++
				istate = 1
				ist = i
			// username
			case 1:
				if ch == ':' {
					user.Username = string(pwdB[ist:i])
					istate = 2;
					break;
				}
			//passwd
			case 2:
				if ch == ':' {istate = 3; break;}
			// Uid
			case 3:
				if ch == ':' {istate = 4; break;}
				num := int(ch)-48
//fmt.Printf("dbg: entry: %d num: %d %q\n", entry, num, ch)
				if num < 0 || num > 9 {return nil, fmt.Errorf("entry %d Uid conversion!", entry)}
				user.Uid = user.Uid*10 + num
			//Gid
			case 4:
				if ch == ':' {ist = i+1; istate = 5; break;}
				num := int(ch)-48
//fmt.Printf("dbg: entry: %d num: %d %q\n", entry, num, ch)
				if num < 0 || num > 9 {return nil, fmt.Errorf("entry %d Gid conversion!", entry)}
				user.Gid = user.Gid*10 + num

			// Info
			case 5:
				if ch == ':' {ist=i+1;istate = 6; break;}

			// Home
			case 6:
				if ch == ':' {
					user.Home = string(pwdB[ist:i])
					ist=i+1;
					istate=8; 
					break;
				}
			// Shell
			case 8:
				if ch == '\n' {
					user.Shell = string(pwdB[ist:i])
					istate = 0;
					pwdList = append(pwdList, user)
					user.Uid = 0
					user.Gid = 0
					break;
				}

			default:
				return nil, fmt.Errorf("invalid state: %d\n", istate)
		}
	}

	return pwdList, nil
}

func GetUser(username string, pwdList []PwdEntry) (id int, err error) {

	if pwdList == nil {return -1, fmt.Errorf("no pwdList provided!")}

	for i:=0; i< len(pwdList); i++ {
		if user == pwdList[i].Username {
			return i, nil
		}
	}
	return -1, nil
}


func Sync(pwdList []PwdEntry) (npwdList []PwdEntry, err error) {

	// no changes
	return nil, nil

	// found changes

}

func PrintPwdList(pwdList []PwdEntry) {
	fmt.Println("************** pwd file *************")
	for i:=0; i<len(pwdList); i++ {
		user:= pwdList[i]
		fmt.Printf("--%d: Name: %s Uid: %d Gid %d\n",i+1, user.Username, user.Uid, user.Gid)
		fmt.Printf("      Home: %s Shell: %s\n", user.Home, user.Shell)
	}
	fmt.Println("************ end pwd file ***********")
}

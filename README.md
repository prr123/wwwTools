# wwwTools

Programs to manage multiple websites.  

## etcpwd

A library that reads the /etc/passwd file and parses the file into a go struct.

### LoadEtcPwd

```
func LoadEtcPwd() (pwdList []PwdEntry, err error)

type PwdEntry struct {
    Username string
    Password string
    Uid      int
    Gid      int
    Info     string
    Home     string
    Shell    string
}
```

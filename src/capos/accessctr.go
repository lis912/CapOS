package capos

import (
	// "bufio"
	"fmt"
	// "io"
	// "log"
	// "os"
	"os/exec"
	"strings"
)

// 访问控制：ssystem32权限和文件共享
type Accesctr struct {
	System32 Term
	Netshare Term
	Totdeg   int
	Mindeg   int
	Maxdeg   int
}

// 默认账户的访问权限
type DefaultUseAccess struct {
	RenameAdmin  Term
	Nodifypwd    Term
	DisableGuset Term
	Totdeg       int
	Mindeg       int
	Maxdeg       int
}

// 无重复多余的用户
type NouSers struct {
	LocakUses      Term
	AdminGroudUses Term
	Totdeg         int
	Mindeg         int
	Maxdeg         int
}

func NousersCheck() (nou NouSers) {
	allocalname := strings.Split(LocalUser(), ",")
	allAdminUses := strings.Split(AdminGroup(), ",")
	if IsrepeatStr(allocalname) {
		nou.LocakUses = Term{LocalUser(), N}
	} else {
		nou.LocakUses = Term{LocalUser(), Y}
	}

	if IsrepeatStr(allAdminUses) {
		nou.AdminGroudUses = Term{AdminGroup(), N}
	} else {
		nou.AdminGroudUses = Term{AdminGroup(), Y}
	}
	nou.Totdeg = (nou.LocakUses.Deg + nou.AdminGroudUses.Deg)
	nou.Maxdeg = 2
	nou.Mindeg = 1

	return nou
}

func DefaultUseAccessCheck() (dua DefaultUseAccess) {
	if isrenameadmin() {
		dua.RenameAdmin = Term{"已", Y}
	} else {
		dua.RenameAdmin = Term{"未", N}
	}
	dua.Nodifypwd = Term{"已", Y}
	if guestAccount() {
		dua.DisableGuset = Term{"已", Y}
	} else {
		dua.DisableGuset = Term{"未", N}
	}
	dua.Totdeg = (dua.RenameAdmin.Deg + dua.Nodifypwd.Deg + dua.DisableGuset.Deg)
	dua.Maxdeg = 3
	dua.Mindeg = 2

	return dua
}

// 获取net share 共享名称 返回C$,D$,E$,F$,G$,IPC$,ADMIN$,如果禁用了共享返回 null
func NetshareCheck() string {
	out, err := exec.Command("cmd", "/C", "net", "share").Output()
	if err != nil {
		return "null"
	}
	restr := ConvertToString(string(out), "gbk", "utf-8")
	sepstr := strings.Split(restr, "\n")
	var netshare string
	for _, line := range sepstr {
		if strings.Contains(line, "$") {
			ret := strings.Split(line, "$")
			netshare = fmt.Sprint(netshare, ret[0], "$", ",")
		}
	}

	return netshare
}

func AccesctrCheck() (access Accesctr) {
	tmp := UserpermisCheck("C:\\Windows\\System32", "Users")
	access.System32 = Term{Format2China(tmp), UsergroupDegCheck(tmp)}
	if NetshareCheck() == "null" {
		access.Netshare = Term{"已禁用", Y}
	} else {
		access.Netshare = Term{NetshareCheck(), N}
	}
	access.Totdeg = (access.System32.Deg + access.Netshare.Deg)
	access.Maxdeg = 2
	access.Mindeg = 1
	return access
}

func Format2China(rwx string) (china string) {
	switch rwx {
	case "N":
		return "无访问权限"
	case "F":
		return "完全访问权限"
	case "M":
		return "修改权限"
	case "RX":
		return "读取和执行权限"
	case "R":
		return "只读权限"
	case "W":
		return "只写权限"
	case "D":
		return "删除权限"
	default:
		return "null"
	}
}

func UsergroupDegCheck(rwx string) int {
	switch rwx {
	case "N":
		return Y
	case "F":
		return Y
	case "M":
		return N
	case "RX":
		return Y
	case "R":
		return Y
	case "W":
		return N
	case "D":
		return N
	default:
		return N
	}
}

func isrenameadmin() bool {
	allname := strings.Split(LocalUser(), ",")
	for _, name := range allname {
		if name == "Administrator" {
			return false
		}
	}
	return true
}

func guestAccount() bool {
	for _, line := range CfgData {
		if strings.Contains(line, "EnableGuestAccount") {
			s := strings.Split(line, " ")
			if string(s[2][0]) == "0" {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

package capos

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sys/windows/registry"
)

// 审计策略，系统下的审计测评统计
type AuditPolicy struct {
	PolicyChange    Term
	LogonEvents     Term
	ObjectAccess    Term
	ProcessTracking Term
	DSAccess        Term
	PrivilegeUse    Term
	SystemEvents    Term
	AccountLogon    Term
	AccountManage   Term
	Totdeg          int
	Mindeg          int
	Maxdeg          int
}

// 剩余信息保护1
type InfoProtion struct {
	NoShowLastname Term
	Totdeg         int
	Mindeg         int
	Maxdeg         int
}

// 剩余信息保护2
type VitmenRemove struct {
	Virmenremove Term // 清楚虚拟内存页面
	Pwdremove    Term // 可还原的密码
	Totdeg       int
	Mindeg       int
	Maxdeg       int
}

// 剩余信息保护3
type PatchPortUpgrade struct {
	Patches  Term // 补丁情况
	Updata   Term // 自动升级
	RiskPort Term // 高危端口
	Totdeg   int
	Mindeg   int
	Maxdeg   int
}

// 屏幕锁定
type Screenlock struct {
	TimeoutLock Term // 屏幕锁定
	Totdeg      int
	Mindeg      int
	Maxdeg      int
}

// 登录超时锁定
func LoginTimeoutLockCheck() (scrlock Screenlock) {
	scrlock.TimeoutLock = ScreenSaveTime()
	scrlock.Totdeg = scrlock.TimeoutLock.Deg
	scrlock.Maxdeg = 1
	scrlock.Mindeg = 0

	return scrlock
}

// 屏保锁定
func ScreenSaveTime() Term {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.QUERY_VALUE)
	if err != nil {
		return Term{"null", N}
	}
	defer k.Close()
	s, _, err := k.GetStringValue("ScreenSaveTimeOut")
	if err != nil {
		return Term{"null", N}
	}
	timemin, err := strconv.Atoi(s)
	timemin = timemin / 60
	d := strconv.Itoa(timemin)
	return Term{d, Y}
}

// 补丁检查
func PatchCheck() (ppu PatchPortUpgrade) {
	patchname, instime := ListNewKb()
	if patchname == "null" {
		ppu.Patches = Term{patchname, N}
	} else {
		slit := strings.Split(instime, "/")
		form := fmt.Sprintf("%s:%s年%s月%s日", patchname, slit[0], slit[1], slit[2])
		if CompareKBTime(time.Now().Format("2006-01-02"), instime, 4) {
			ppu.Patches = Term{form, Y}
		} else {
			ppu.Patches = Term{form, N}
		}
	}

	if ServerState("Windows Update") {
		ppu.Updata = Term{"已开启", Y}
	} else {
		ppu.Updata = Term{"已关闭", N}
	}

	if PortState("445") || PortState("135") {
		str := ""
		str2 := ""
		if PortState("445") {
			str = "445"
		}
		if PortState("135") {
			str2 = "135"
		}
		ppu.RiskPort = Term{(fmt.Sprintf("%s,%s", str, str2)), N}
	} else {
		ppu.RiskPort = Term{"未开启", Y}
	}

	ppu.Totdeg = (ppu.Patches.Deg + ppu.Updata.Deg + ppu.RiskPort.Deg)
	ppu.Maxdeg = 3
	ppu.Mindeg = 1

	return ppu
}

// 查询windows系统的服务状态，返回bool值 开启返回true 未开启返回 false
func ServerState(serername string) bool {
	out, err := exec.Command("cmd", "/C", "net", "start").Output()
	if err != nil {
		return false
	}
	var cmdata []string
	var line []byte
	for _, by := range out {
		if by == 10 {
			cmdata = append(cmdata, string(line))
			line = nil
			continue
		}
		line = append(line, by)
	}
	for _, lines := range cmdata {
		if strings.Contains(lines, serername) {
			return true
		}
	}
	return false
}

// systime format:2018-08-12 and kbtime format:2017/12/14 补丁安装时间和系统时间的差值小于standmonths个月返回真，否则返回假
func CompareKBTime(systime string, kbtime string, standmonths int) bool {
	// 将两个时间字符串格式化成切片
	sys := strings.Split(systime, "-")
	kb := strings.Split(kbtime, "/")
	var difmonths int
	// 计算出年的差值
	// int, err := strconv.Atoi(string)
	s, _ := strconv.Atoi(sys[0])
	k, _ := strconv.Atoi(kb[0])
	difyears := s - k
	// difyears := strconv.Atoi(sys[0]) - strconv.Atoi(kb[0])
	// 如果是同一年的话，进而比较月，而月的系统时间必定大于等于补丁安装月时间
	if difyears == 0 {
		// 计算出月的差值：months必定大于等0
		s, _ = strconv.Atoi(sys[1])
		k, _ = strconv.Atoi(kb[1])
		difmonths = s - k
		// difmonths = strconv.Atoi(sys[1]) - strconv.Atoi(kb[1])
	} else {
		s, _ = strconv.Atoi(sys[1])
		k, _ = strconv.Atoi(kb[1])
		difmonths = s - k
		// 如果年数差值大于0，则这样计算计算出总共的月差数
		// difmonths = (difyears * 12) + (strconv.Atoi(sys[1]) - strconv.Atoi(kb[1]))
		difmonths = (difyears * 12) + difmonths
	}
	// 再次判断时间是否小于标准月数
	if difmonths < standmonths {
		return true
	} else {
		return false
	}
	return true
}

// 获取最新补丁安装信息 返回格式:（KB3150513 2017/12/14）
func ListNewKb() (kbname string, instatime string) {
	command := "powershell"
	a := []string{"get-hotfix", "|", "findstr", "KB"}
	out, err := exec.Command(command, a...).Output()
	if err != nil {
		return "null", "null"
	}
	retcov := ConvertToString(string(out), "gbk", "utf-8")
	lines := strings.Split(retcov, "\n")
	firstline := pasestring(lines[0], ',')
	s := strings.Split(firstline, ",")
	for _, kb := range s {
		if strings.Contains(kb, "KB") {
			kbname = kb
		}
		// 如果某一个字符串包含20，那么必定是时间字符
		if strings.Contains(kb, "20") {
			// instatime = kb
			instatime = Removebyte(kb, '.')
		}
	}
	// instatime = s[len(s)-2]
	return kbname, instatime
}

func VirMemPwdCheck() (info VitmenRemove) {
	info.Virmenremove = Vitmen()
	info.Pwdremove = Passredus()
	info.Totdeg = (info.Virmenremove.Deg + info.Pwdremove.Deg)
	info.Maxdeg = 2
	info.Mindeg = 1

	return info
}

// 是否删除虚拟页面
func Vitmen() Term {
	ret := Term{"null", N}
	for _, line := range CfgData {
		if strings.Contains(line, "ClearPageFileAtShutdown") {
			s := strings.Split(line, ",")
			if s[1][0] == '0' {
				// ret = "2)关机清除虚拟内存页面文件: 已禁用 不符合"
				ret = Term{"已禁用", N}
			} else {
				ret = Term{"已启用", Y}
			}
			return ret
		}
	}
	return ret
}

// 用可还原的加密来储存密码
func Passredus() Term {
	ret := Term{"null", N}
	for _, line := range CfgData {
		if strings.Contains(line, "ClearTextPassword") {
			s := strings.Split(line, " ")
			if s[2][0] == '0' {
				ret = Term{"已禁用", Y}
			} else {
				ret = Term{"已启用", N}
			}
			return ret
		}
	}
	return ret
}

func AuditCheck() (audit AuditPolicy) {
	audit.PolicyChange = AuditpolicyCheck("AuditPolicyChange")
	audit.LogonEvents = AuditpolicyCheck("AuditLogonEvents")
	audit.ObjectAccess = AuditpolicyCheck("AuditObjectAccess")
	audit.ProcessTracking = AuditpolicyCheck("AuditProcessTracking")
	audit.DSAccess = AuditpolicyCheck("AuditDSAccess")
	audit.PrivilegeUse = AuditpolicyCheck("AuditPrivilegeUse")
	audit.SystemEvents = AuditpolicyCheck("AuditSystemEvents")
	audit.AccountLogon = AuditpolicyCheck("AuditAccountLogon")
	audit.AccountManage = AuditpolicyCheck("AuditAccountManage")

	audit.Totdeg = (audit.PolicyChange.Deg + audit.LogonEvents.Deg + audit.ObjectAccess.Deg + audit.ProcessTracking.Deg + audit.DSAccess.Deg + audit.PrivilegeUse.Deg + audit.SystemEvents.Deg + audit.AccountLogon.Deg + audit.AccountManage.Deg)
	audit.Maxdeg = 9
	audit.Mindeg = 2

	return audit
}

//审计策略
func AuditpolicyCheck(src string) Term {
	// 解析配置文件
	for _, line := range CfgData {
		if strings.Contains(line, src) {
			s := strings.Split(line, " ")
			switch string(s[2]) {
			case "0":
				return Term{"无审核", N}
			case "1":
				return Term{"成功", Y}
			case "2":
				return Term{"失败", Y}
			case "3":
				return Term{"成功，失败", Y}
			default:
				return Term{"null", N}
			}
		}
	}
	return Term{"null", N}
}

// 不显示最后登录的用户名
func NoLastnameCheck() (lastname InfoProtion) {
	lastname.NoShowLastname = Loginuser()
	lastname.Totdeg = lastname.NoShowLastname.Deg
	lastname.Maxdeg = 1
	lastname.Mindeg = 0

	return lastname
}

// 不显示上次登录用户名
func Loginuser() Term {
	// var ret Term
	ret := Term{"null", N}
	for _, line := range CfgData {
		if strings.Contains(line, "DontDisplayLastUserName") {
			s := strings.Split(line, ",")
			if s[1][0] == '0' {
				ret = Term{"已禁用", N}
			} else {
				ret = Term{"已启用", Y}
			}
			return ret
		}
	}
	return ret
}

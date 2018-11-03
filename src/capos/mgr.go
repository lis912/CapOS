package capos

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

var CfgData []string

func init() {
	PrintLogo()
	CreatCfgFile()
	ParsingFile()
	DelCfgFile()

}

func CreatCfgFile() bool {
	// 生成配置文件：
	info := exec.Command("cmd", "/C", "secedit", "/export", "/cfg", "C:\\cfg1.inf")
	if err := info.Run(); err != nil {
		fmt.Println("Error: ", err)
		return false

	}
	return true
}

func ParsingFile() {
	// 打开配置文件解析文本
	// fi, err := os.Open("C:/cfg.inf")
	fi, err := os.Open("C:/cfg1.inf")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	// 解析文本
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		// 将UCS-2编码数据流 转化为 []byte 切片
		a = []byte(a)
		var sa []byte
		for i := 1; i < len(a)-1; i = i + 2 {
			// 经过研究，如果是纯英文之母的UCS-2编码，只需要把奇数index拿出来即可完成
			// utf-8的转换，所以这个代码比较投机，而不具备普遍性
			// 去除尾部回车键
			if a[i] == 13 {
				continue
			}
			sa = append(sa, a[i])

		}

		CfgData = append(CfgData, string(sa))

	}

}
func DelCfgFile() bool {
	// 删除配置文件：
	info := exec.Command("cmd", "/C", "del", "C:\\cfg1.inf")
	if err := info.Run(); err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	return true
}

// 密码复杂度
func passwordComplexity() Term {

	p := Term{"null", 0}
	// 解析配置文件
	for _, line := range CfgData {
		if strings.Contains(line, "PasswordComplexity") {
			s := strings.Split(line, " ")
			if s[2][0] == '0' {
				p = Term{"已禁用", 0}

			} else {
				p = Term{"已启用", 1}

			}

			return p
		}
	}
	return p
}

// 检查其他的值
func Checkcfg(cfg string, stand int32) Term {
	p := Term{"null", N}
	// 解析配置文件
	for _, line := range CfgData {
		if strings.Contains(line, cfg) {
			s := strings.Split(line, " ")
			p.Val = s[2]
			if StrToint(s[2]) < stand {
				p.Deg = N
			} else {
				p.Deg = Y
			}

			return p
		}
	}

	return p
}

func LogFailCheck() (login Logfaile) {
	login.LockTime = Checkcfg("LockoutDuration", 30)
	if login.LockTime.Val == "null" {
		login.LockTime.Val = "不适用"
	}
	login.LockCout = Checkcfg("LockoutBadCount", 3)
	login.ResTime = Checkcfg("ResetLockoutCount", 30)
	if login.ResTime.Val == "null" {
		login.ResTime.Val = "不适用"
	}
	login.Totdeg = (login.LockTime.Deg + login.LockCout.Deg + login.ResTime.Deg)
	login.Mindeg = 1
	login.Maxdeg = 3

	return login
}

// 检查远程管理
func RemoManCheck() (sysman RemoManag) {
	if PortState("23") {
		sysman.Telnet = Term{"开启", N}
	} else {
		sysman.Telnet = Term{"禁用", Y}
	}
	if PortState("3389") {
		sysman.WinDesk = Term{"开启", Y}
	} else {
		sysman.WinDesk = Term{"禁用", N}
	}
	sysman.Totdeg = (sysman.Telnet.Deg + sysman.WinDesk.Deg)
	sysman.Maxdeg = 2
	sysman.Mindeg = 1

	return sysman
}

// 检查重复账户
func UsersCheck() (users Userinfo) {

	if LocalUser() == "null" {
		users.Username = Term{LocalUser(), N}
	} else {
		users.Username = Term{LocalUser(), Y}
	}

	users.Totdeg = (users.Username.Deg + users.Adminame.Deg)
	users.Maxdeg = 1
	users.Mindeg = 0
	return users
}

// 获取本地用户状态
func LocalUser() string {

	out, err := exec.Command("cmd", "/C", "net", "user").Output()
	if err != nil {
		return "null"
	}

	cmdata := ByteToString(out)
	var strtmp string
	for _, lines := range cmdata[4:(len(cmdata) - 2)] {
		strtmp += lines
	}

	return pasestring(strtmp, ',')
}

// 获取管理员用户状态
func AdminGroup() string {
	out, err := exec.Command("cmd", "/C", "net", "Localgroup", "administrators").Output()
	if err != nil {
		return "null"
	}
	cmdata := ByteToString(out)
	var strtmp string
	for _, lines := range cmdata[6:(len(cmdata) - 2)] {
		strtmp += (lines + " ")
	}
	return pasestring(strtmp, ',')
}

// 是否勾选用户名密码登录
func IsUseIdPwd() UseIdPwd {
	var s UseIdPwd
	str := capLocalReginfo(`SOFTWARE\Microsoft\Windows NT\CurrentVersion\Winlogon`, "AutoAdminLogon")
	if str == "null" {
		s.IsUse = Term{"null", N}
	} else if str == "1" {
		s.IsUse = Term{"未勾选", N}
	} else {
		s.IsUse = Term{"已勾选", Y}
	}
	s.Totdeg = s.IsUse.Deg
	s.Maxdeg = 1
	s.Mindeg = 0
	return s
}

// 密码策略
func PaswdPolicy() Passwd {
	var pwd Passwd
	// 密码复杂度
	pwd.Comple = passwordComplexity()
	// 密码长度最小值标准8
	pwd.Miniwd = Checkcfg("MinimumPasswordLength", 8)
	// 最短使用期限标准 1
	pwd.Agewd = Checkcfg("MinimumPasswordAge", 1)
	// 最长使用期限标准 30
	pwd.Longwd = Checkcfg("MaximumPasswordAge", 60)
	// 强制密码历史标准 24
	pwd.Hisywd = Checkcfg("PasswordHistorySize", 24)
	// 计算本项总分
	pwd.Totdeg = (pwd.Comple.Deg + pwd.Miniwd.Deg + pwd.Agewd.Deg + pwd.Longwd.Deg + pwd.Hisywd.Deg)
	// 判断标准
	pwd.Mindeg = 2
	pwd.Maxdeg = 5

	return pwd
}

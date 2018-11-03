package capos

import (
	"fmt"
	// "golang.org/x/sys/windows/registry"
	"os/exec"
	// "strconv"
	"strings"
	// "time"
)

// Get-Service | Where-Object Status -eq Running
// 打印所有的服务状态
func PnowServer() string {
	cmdstring := []string{"/C", "net", "start"}
	out := Prinfcmd("cmd", cmdstring)
	return fmt.Sprintf("%s", out)
}

// 打印所有的补丁情况
func Pkbinfo() string {
	cmdstring := []string{"Get-hotfix"}
	out := Prinfcmd("powershell", cmdstring)

	return fmt.Sprintf("\n\n--补丁安装明细-------------%s", out)
}

// 打印所有运行进程
func Pnowpross() string {
	cmdstring := []string{"ps"}
	out := Prinfcmd("powershell", cmdstring)

	return fmt.Sprintf("--系统进程列表-------------%s", out)
}

// 打印所有的安装的程序
// Get-WmiObject -class Win32_Product |Select-Object -Property name,version
func PinstallExe() string {
	cmdstring := []string{"Get-WmiObject", "-class", "Win32_Product", "|", "Select-Object", "-Property", "name", ",", "version"}
	out := Prinfcmd("powershell", cmdstring)

	return fmt.Sprintf("--程序安装列表-------------%s", out)
}

// 本地组合管理员组情况
// 获取本地用户状态
func PsysUser() string {
	return fmt.Sprintf("--系统账户信息-------------\n本地所有账户：%s\n管理员组账户：%s\n", LocalUser(), AdminGroup())
}

// 打印systeminfo输出
func Psysteminfo() string {
	var retstr string
	cmdstring := []string{"/C", "systeminfo"}
	out := Prinfcmd("cmd", cmdstring)
	outslice := strings.Split(out, "\n")
	for _, line := range outslice {
		if strings.Contains(line, "修补程序") {
			break
		}
		retstr = fmt.Sprint(retstr, line)
	}

	return fmt.Sprintf("--系统信息-------------%s", retstr)
}

// 打印输出系统命令
func Prinfcmd(cmdtype string, cmdslice []string) string {
	out, err := exec.Command(cmdtype, cmdslice...).Output()
	if err != nil {
		return "未获取到数据"
	}
	restr := ConvertToString(string(out), "gbk", "utf-8")
	return restr
}

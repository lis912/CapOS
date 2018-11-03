package capos

import (
	"fmt"
	"net"
	// "os"
	// "os/exec"
	// "strconv"
	"strings"
	"time"
)

// 系统和硬件信息
type Sysinfos struct {
	SysVersion     string
	Board          string
	Model          string
	Bios           string
	Motherboard    string
	Memory         string
	Storage        string
	Cpu            string
	ExportFileName string
}

func SysinfosCheck() (sysinfo Sysinfos) {
	sysinfo.ExportFileName = ExportFileName()

	return sysinfo
}

func ExportFileName() string {
	systime := time.Now().Format("20060102-150405")
	ipaddr := GetPulicIP()
	filename := fmt.Sprint(ipaddr, "-", systime)
	return filename
}

func GetPulicIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "localhost"
		// log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

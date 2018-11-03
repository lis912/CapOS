package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"golang.org/x/sys/windows/registry"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Passwd struct {
	Comple term
	Miniwd term
	Agewd  term
	Longwd term
	Hisywd term
	Totdeg int
}

type term struct {
	name string
	age  int
}

const bannerLogo = `
别名     administrators
注释     管理员对计算机/域有不受限制的完全访问权

成员

-------------------------------------------------------------------------------
Administrator
命令成功完成。`

func main() {
	// info := "my name      is lishichangfasf as fdas afas  fds f"
	// hel := haha{"lishichang", 30}
	// hel = haha{"lizhong", 60}
	// fmt.Println(hel)
	// fmt.Println(StrToint("543453245"))
	// fmt.Printf("%T", StrToint("54"))

	// fmt.Println("*************")
	// t := "12"
	// fmt.Println(len(t))
	// for _, c := range t {

	// 	fmt.Printf("%T   ", c)
	// 	fmt.Println(c)
	// }
	// test("3")

	// test2()
	// fmt.Println(pasestring(bannerLogo, ':'))

	// fmt.Println(IsUserState())
	// fmt.Println(AdminGroud2())
	// IsGroudState()
	// testgolang(bannerLogo)
	// getscreen()
	// fmt.Println(System32ACL())  C:\file
	// k := UserpermisCheck("C:\\Windows\\System32", "Users")
	// fmt.Println(k)
	// k = UserpermisCheck("C:\\file", "Administrators")
	// fmt.Println(k)
	// fmt.Println(NetshareCheck())
	// str := "fsadfa"
	// t := Capstr(str, '(', ')')
	// AllProcess()
	// fmt.Println(t)

	// lele := []string{"lld", "sfas", "sga", "dsfas", "kk"}
	// if !IsrepeatStr(lele) {
	// 	fmt.Println("无重复")
	// }
	systime := time.Now().Format("20060102-150405")
	fmt.Println(systime)
	// ipaddr := GetPulicIP()
	ipaddr := getLocalIp()
	filename := fmt.Sprint(ipaddr, "-", systime)
	fmt.Println(filename)
	// fmt.Printf("%T \n", time)
	// fmt.Println(ListNewKb())
	// _, kbinstime := ListNewKb()
	// if CompareKBTime(systime, kbinstime, 7) {
	// 	fmt.Println("真的")
	// } else {
	// 	fmt.Println("假的")
	// }

	fmt.Println()
	// fmt.Println("\n\n")
	// fmt.Println("任意键退出………")
	var str string
	for {
		fmt.Scan(&str)
		break
	}

}

func GetPulicIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("========================")
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

func getLocalIp() (IpAddr string) {
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		// log.Error("Get local IP addr failed!!!")
		IpAddr = "localhost"
		return IpAddr
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				IpAddr = ipnet.IP.String()
				return IpAddr
			}
		}
	}
	return "localhost"
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

// systime format:2018-08-12 and kbtime format:2017/12/14 补丁安装时间和系统时间的差值小于standmonths个月返回真，否则返回假
// func CompareKBTime(systime string, kbtime string, standmonths int32) bool {
// 	// 将两个时间字符串格式化成切片
// 	sys := strings.Split(systime, "-")
// 	kb := strings.Split(kbtime, "/")
// 	var difmonths int32
// 	// 计算出年的差值
// 	difyears := StrToint(sys[0]) - StrToint(kb[0])
// 	// 如果是同一年的话，进而比较月，而月的系统时间必定大于等于补丁安装月时间
// 	if difyears == 0 {
// 		// 计算出月的差值：months必定大于等0
// 		difmonths = StrToint(sys[1]) - StrToint(kb[1])
// 	} else {
// 		// 如果年数差值大于0，则这样计算计算出总共的月差数
// 		difmonths = (difyears * 12) + (StrToint(sys[1]) - StrToint(kb[1]))
// 	}
// 	// 再次判断时间是否大于标准月数
// 	if difmonths < standmonths {
// 		return true
// 	} else {
// 		return false
// 	}
// 	return true
// }

// 获取最新补丁安装信息
func ListNewKb() (kbname string, instatime string) {
	command := "powershell"
	// get-hotfix  | Select-Object HotFixID, InstalledOn | findstr "KB"
	// b := []string{"get-hotfix", "|", "Select-Object", "HotFixID", "InstalledOn"}
	a := []string{"get-hotfix", "|", "findstr", "KB"}
	out, err := exec.Command(command, a...).Output()
	if err != nil {
		fmt.Println("gsdf")
		log.Fatal(err)
		return "未安装任何系统补丁", "null"
	}
	retcov := ConvertToString(string(out), "gbk", "utf-8")
	// fmt.Println(retcov)
	lines := strings.Split(retcov, "\n")
	firstline := pasestring(lines[0], ',')
	// fmt.Println(lines)
	s := strings.Split(firstline, ",")
	for _, kb := range s {
		if strings.Contains(kb, "KB") {
			kbname = kb
		}
	}
	instatime = s[len(s)-2]
	// 返回如下格式:（KB3150513 2017/12/14）
	// fmt.Println(kbname, instatime)
	return kbname, instatime
	// return "KB3150513", "2017/12/14"
}

// 判断一个字符串切片中是否有重复内容
func IsrepeatStr(src []string) bool {
	tmp := src
	cnt := 1
	for i, t := range tmp {
		for _, s := range src[(i + 1):] {
			fmt.Println(s)
			if t == s {
				cnt += 1
				fmt.Println(s)
			}
		}
	}
	// fmt.Printf("cnt = %d", cnt)
	if cnt > 1 {
		fmt.Println("reap", cnt)
		return true
	} else {
		fmt.Println("noreap", cnt)
		return false
	}
}

// 获取系统所有安装的程序和组件
func AllProcess() {
	key, exists, err := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\AppCompatFlags\Compatibility Assistant\Store`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	if !exists {
		fmt.Println("ERROR: 未找到注册表信息！")
	}

	// 枚举所有值名称
	// qq管家 QQPCTray.exe
	keynames, _ := key.ReadValueNames(0)
	for _, key := range keynames {
		nameexe := strings.Split(key, `\`)
		// prossname := strings.Split(s[len(s)-1], `.`)
		// if strings.Contains(s[len(s)-1], "QQPCTray") {
		// 	prossname := strings.Split(s[len(s)-1], `.`)
		// 	fmt.Println(prossname[0])
		// 	break
		// }

		// prossname := strings.Split(nameexe[len(nameexe)-1], `.`)
		// fmt.Println(prossname[0])

		fmt.Println(nameexe[len(nameexe)-1])
	}

}

// 字符串数字转换为纯整形数字,不能转负数
func StrToint(str string) (num int32) {
	for _, char := range str {
		if (char >= '0') && (char <= '9') {
			// 利用ASCII码的错位实现
			num = num*10 + (char - '0')
		} else {
			fmt.Println("parame error!")
			return 0
		}
	}

	return num
}

func test2() {
	var tmp Passwd
	tmp.Comple = term{"li", 30}
	tmp.Miniwd = term{"shi", 39}
	tmp.Agewd = term{"liujin", 58}
	tmp.Longwd = term{"mam", 23}
	tmp.Hisywd = term{"he", 90}

	p := &tmp.Comple
	fmt.Printf("%T ########### ", ':')
	fmt.Println(*p)

	fmt.Println(*p)
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// func pasestring(str string) string{
// 	var words []int32
// 	flag := false
// 	// for _, lines := range str {
// 	for _, word := range str {
// 		if word != 32 {
// 				words = append(words, word)
// 				if flag {
// 					words = append(words, ',')
// 					flag = false
// 				}
// 				continue
// 			}
// 			if word == 32 {
// 				flag = true
// 				continue
// 			}

// 		}
// 		fmt.Println(string(words))

// }

// 将一个字符语句，按照sep分隔开来,可以把一段语句的空格 回车换行去掉 例如：pasestring(string, ',')
func pasestring(str string, sep int32) string {
	flag := false
	var words []int32
	for _, word := range str {
		if word != 32 && word != 13 && word != 10 {
			if flag {
				words = append(words, sep)
				flag = false
			}
			// fmt.Println(word)
			words = append(words, word)
			continue
		}
		if word == 32 {
			flag = true
			continue
		}

	}
	return string(words)
}

// 获取本地用户状态
func IsUserState() string {

	out, err := exec.Command("cmd", "/C", "net", "user").Output()
	if err != nil {
		// return ret
		log.Fatal(err)
	}

	cmdata := ByteToString(out)
	var strtmp string
	for _, lines := range cmdata[4:(len(cmdata) - 2)] {
		strtmp += lines
	}

	return pasestring(strtmp, ',')
}

// []uint8 转 []string
func ByteToString(bytes []uint8) (strings []string) {
	var line []byte
	for _, by := range bytes {
		if by == '-' {
			continue
		}
		// 10是换行的ASCII码，我这里写死了
		if by == 10 {
			strings = append(strings, string(line))
			line = nil
			continue
		}
		line = append(line, by)
	}
	return strings
}

func AdminGroud2() string {

	out, err := exec.Command("cmd", "/C", "net", "Localgroup", "administrators").Output()
	if err != nil {
		// return ret
		log.Fatal(err)
	}

	cmdata := ByteToString(out)
	var strtmp string
	for _, lines := range cmdata[6:(len(cmdata) - 2)] {
		strtmp += lines
	}
	fmt.Println(strtmp)
	return pasestring(strtmp, ',')
}

// 获取用户组状态
func IsGroudState() {

	out, err := exec.Command("cmd", "/C", "systeminfo").Output()
	if err != nil {
		// return ret
		log.Fatal(err)

	}

	// cmdout := string(out)
	// fmt.Println(string(out))
	// fmt.Printf("%T", cmdout)

	haha := ConvertToString(string(out), "gbk", "utf-8")
	fmt.Println(haha)

	// 由于输出的out是 []uint8 ,所以我得大费周章按照命令输出的每一行添加到[]string
	// var cmdata []string
	// var line []byte

	// for _, by := range out {
	// 	fmt.Printf("%T", by)
	// 	fmt.Println(by)
	// 	// if by == '-' {
	// 	// 	continue
	// }
	// // 10是换行的ASCII码，我这里写死了
	// if by == 10 {
	// 	// line = append(line, ' ')
	// 	cmdata = append(cmdata, string(line))
	// 	line = nil
	// 	continue
	// }
	// line = append(line, by)

	// }
	// var strtmp string
	// for _, lines := range cmdata[6:(len(cmdata) - 2)] {
	// 	// fmt.Println(lines)
	// 	// for _, w := range lines {
	// 	// 	fmt.Println(w)
	// 	// }

	// 	strtmp += (lines + " ")
	// }
	// fmt.Println(pasestring(strtmp, ','))
	// for _, sss := range pasestring(strtmp, ',') {
	// 	fmt.Println(sss)
	// }
	// return pasestring(strtmp, ',')

}

func getscreen() {
	command := "cmd"
	a := []string{"/C", "systeminfo"}
	cmd := exec.Command(command, a...)
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	fmt.Printf("out: %#v\n", string(out))
	fmt.Printf("err: %#v\n", err)
	if err != nil {
		log.Fatal(err)
	}
}

func testgolang(str string) {
	// fmt.Println(str)
	for _, s := range str {

		fmt.Println(s)
	}
}

// 查询组成员命令，目前有适配有问题
func UserGroupCheck(groupname string) string {

	out, err := exec.Command("cmd", "/C", "net", "Localgroup", groupname).Output()
	if err != nil {
		// return ret
		log.Fatal(err)
	}

	cmdata := ByteToString(out)
	// return cmdata
	var strtmp string
	for _, lines := range cmdata[6:(len(cmdata) - 2)] {
		fmt.Sprint(lines, " ")
		strtmp += lines
	}
	// fmt.Println(strtmp)
	return strtmp
	// return string(words)
	// return pasestring(strtmp, ' ')
}

// Cacls C:\Windows\System32
func System32ACL() string {

	out, err := exec.Command("cmd", "/C", "Cacls", "C:\\Windows\\System32").Output()
	if err != nil {
		log.Fatal(err)
	}
	restr := ConvertToString(string(out), "gbk", "utf-8")
	fmt.Println(restr)
	sepstr := strings.Split(restr, "\n")
	for _, line := range sepstr {
		if strings.Contains(line, "Users") {
			ret := strings.Split(line, ":")
			return RemoveEnt(ret[1])
		}
	}

	return "null"
}

// 获取groupname组对系统文件夹pathname文件夹的权限 返回 'R', 'F', 'C'  (F)完全控制 (C)写入
func UserpermisCheck(pathname string, groupname string) string {

	out, err := exec.Command("cmd", "/C", "Cacls", pathname).Output()
	if err != nil {
		log.Fatal(err)
	}
	restr := ConvertToString(string(out), "gbk", "utf-8")
	sepstr := strings.Split(restr, "\n")
	for _, line := range sepstr {
		if strings.Contains(line, groupname) {
			ret := strings.Split(line, ":")
			// Capstr(line, '', end)
			return RemoveEnt(ret[1])
		}
	}

	return "null"
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

// 过滤掉一个字符串中的空格和回车
func RemoveEnt(str string) string {
	// flag := false
	var words []int32
	for _, word := range str {
		if word != 32 && word != 13 && word != 10 {
			words = append(words, word)
			continue
		}
	}
	return string(words)
}

// import (
// 	"bufio"
// 	"io"
// 	"os"
// 	"strings"
// )

// func ReadLine(fileName string, handler func(string)) error {
// 	f, err := os.Open(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	buf := bufio.NewReader(f)
// 	for {
// 		line, err := buf.ReadString('\n')
// 		line = strings.TrimSpace(line)
// 		handler(line)
// 		if err != nil {
// 			if err == io.EOF {
// 				return nil
// 			}
// 			return err
// 		}
// 	}
// 	return nil
// }

// 截取字符串 begin 到 end 之间的内容  例子：Capstr(str, '(', ')')
func Capstr(src string, begin int32, end int32) string {
	var ret []int32
	flag := false
	for _, word := range src {
		if word == begin {
			flag = true
			continue
		} else if word == end {
			flag = false
			continue
		}
		if flag {
			ret = append(ret, word)
		}
	}
	return string(ret)
}

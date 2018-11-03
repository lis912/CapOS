package capos

import (
	// "bufio"
	"fmt"
	"golang.org/x/sys/windows/registry"
	// "io"
	"log"
	// "os"
	"github.com/axgle/mahonia"
	"os/exec"
	"strings"
)

/*
-------------------------------------------------------------------------------
本文件函数一览：

// 获取注册表 LOCAL_MACHINE 下的值
func capLocalReginfo(path string, key string) string
// 获取端口状态,返回true为打开状态，false是关闭状态
func PortState(port string) bool
// 打印组成员命令,目前可以使用
func PrintGroupCheck(groupname string)
// 获取groupname组对系统文件夹pathname文件夹的权限 返回 'R', 'F', 'C'  (F)完全控制 (C)写入  有bug未解决
func UserpermisCheck(pathname string, groupname string) string
// 打印文件夹权限
func Prinfcmd(pathname string) string

// cmd输出转码 GBK 转 UTF-8
func ConvertToString(src string, srcCode string, tagCode string) string
// 字符串数字转换为纯整形数字,不能转负数
func StrToint(str string) (num int32)
// 将一个字符语句，按照sep分隔开来,可以把一段语句的空格 回车换行去掉 例如：pasestring(string, ',')
func pasestring(str string, sep int32) string
// 截取字符串 begin 到 end 之间的内容  例子：Capstr(str, '(', ')')
func Capstr(src string, begin int32, end int32) string
// 过滤掉一个字符串中的空格和回车
func RemoveEnt(str string) string
// 过滤掉一个字符串中的任意某个字符
func Removebyte(str string, by int32) string
// 判断一个字符串切片中是否有重复内容
func IsrepeatStr(src []string) bool
// []uint8 转 []string，专用来作为 cmd命令输出后的转码
func ByteToString(bytes []uint8) (strings []string)
--------------------------------------------------------------------------------------
*/

// 获取注册表 LOCAL_MACHINE 下的值
func capLocalReginfo(path string, key string) string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE)
	if err != nil {
		return "null"
	}
	defer k.Close()
	value, _, err := k.GetStringValue(key)
	if err != nil {
		return "null"
	}

	return value
}

// 字符串数字转换为纯整形数字,不能转负数
func StrToint(str string) (num int32) {
	for _, c := range str {
		if (c >= '0') && (c <= '9') {
			// 利用ASCII码的错位实现
			num = num*10 + (c - '0')
		} else {
			fmt.Println("parame error!")
			return 0
		}
	}

	return num
}

// 计算最终的符合度，鉴于符合度有一定主观性，所以我先只是判断符合，不符合，部分符合，不适用 四项
func CalConDeg(Totdeg int, Mindeg int, Maxdeg int) (Deg string) {
	if Totdeg == NoApply {
		Deg = "不适用"
	}
	if Maxdeg == 1 {
		if Totdeg == 0 {
			Deg = "不符合"
		} else if Totdeg == 1 {
			Deg = "符合"
		} else {
			Deg = "本工具计算异常"
		}
		return Deg
	}

	if Totdeg < Mindeg {
		Deg = "不符合"
	} else if (Totdeg >= Mindeg) && (Totdeg < Maxdeg) {
		Deg = "部分符合"
	} else if Totdeg == Maxdeg {
		Deg = "符合"
	} else {
		Deg = "本工具计算异常"
	}

	return Deg
}

// 获取端口状态,返回true为打开状态，false是关闭状态
func PortState(port string) bool {
	ret := false
	// 得到命令的输出out
	out, err := exec.Command("cmd", "/C", "netstat", "-ano", "|", "findstr", port).Output()
	// 如果命令结果没有输出，这里会发生报错，进入if语句，当然这也说明没有开启
	if err != nil {
		return ret
	}
	// 由于输出的out是 []uint8 ,所以我得大费周章按照命令输出的每一行添加到[]string
	var cmdata []string
	var line []byte

	for _, by := range out {
		// 10是换行的ASCII码，我这里写死了
		if by == 10 {
			cmdata = append(cmdata, string(line))
			line = nil
			continue
		}
		line = append(line, by)
	}
	// 遍历[]string，提取我们最终要的数据
	// 经过分析后得知，netstat -ano输出的每一行的格式是以12个空格间隔和4个空格为间隔的列
	// 总之一步得到了最后的结果，测试的很辛苦啊
	kong12 := "            "
	kong4 := "    "
	kongdian := ":"
	for _, lines := range cmdata {
		// 这一步检查LISTENING其实不太重要，保险起见吧
		if strings.Contains(lines, "LISTENING") {
			// "  TCP    0.0.0.0:445"
			s := strings.Split(lines, kong12)
			// "0.0.0.0:445"
			s = strings.Split(s[0], kong4)
			// "445"
			s = strings.Split(s[1], kongdian)
			// 取出最后一个判断
			if s[len(s)-1] == port {
				ret = true
				break
			}
		}
	}
	return ret
}

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

// 过滤掉一个字符串中的空格和回车
func RemoveEnt(str string) string {
	var words []int32
	for _, word := range str {
		if word != 32 && word != 13 && word != 10 {
			words = append(words, word)
			continue
		}
	}
	return string(words)
}

// 过滤掉一个字符串中的任意某个字符
func Removebyte(str string, by int32) string {
	var words []int32
	for _, word := range str {
		if word != by {
			words = append(words, word)
			continue
		}
	}
	return string(words)
}

// 判断一个字符串切片中是否有重复内容
func IsrepeatStr(src []string) bool {
	tmp := src
	cnt := 1
	for i, t := range tmp {
		for _, s := range src[(i + 1):] {
			if t == s {
				cnt += 1
			}
		}
	}
	if cnt > 1 {
		return true
	} else {
		return false
	}
}

// []uint8 转 []string，专用来作为 cmd命令输出后的转码
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

// 打印组成员命令,目前可以使用
func PrintGroupCheck(groupname string) {

	out, err := exec.Command("cmd", "/C", "net", "Localgroup", groupname).Output()
	if err != nil {
		fmt.Println("用户列表为空")
		return
	}
	cmdata := ByteToString(out)
	for _, lines := range cmdata[6:(len(cmdata) - 2)] {
		fmt.Println(lines)
	}
}

// cmd输出转码 GBK 转 UTF-8
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// 获取groupname组对系统文件夹pathname文件夹的权限 返回 'R', 'F', 'C'  (F)完全控制 (C)写入  有bug未解决
func UserpermisCheck(pathname string, groupname string) string {

	out, err := exec.Command("cmd", "/C", "Icacls", pathname).Output()
	if err != nil {
		log.Fatal(err)
	}
	restr := ConvertToString(string(out), "gbk", "utf-8")
	sepstr := strings.Split(restr, "\n")
	for _, line := range sepstr {
		if strings.Contains(line, groupname) {
			return Capstr(line, '(', ')')
		}
	}

	return "null"
}

// 打印文件夹权限
func PrinFolderPess(pathname string) string {

	out, err := exec.Command("cmd", "/C", "Cacls", pathname).Output()
	if err != nil {
		log.Fatal(err)
	}
	restr := ConvertToString(string(out), "gbk", "utf-8")
	return restr
}

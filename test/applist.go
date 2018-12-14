package main

import (
	"bufio"
	"fmt"
	"io"
	// "log"
	"os"
	"os/exec"
)

var CfgData []string

func init() {
	// PrintLogo()
	CreatCfgFile()
	ParsingFile()
	DelCfgFile()

}

// var CfgData []string

func main() {
	fmt.Println(CfgData)
}

func CreatCfgFile() bool {
	// 生成配置文件：
	info := exec.Command("cmd", "/C", "secedit", "/export", "/cfg", "C:\\cfg.inf")
	if err := info.Run(); err != nil {
		fmt.Println("Error: ", err)
		return false

	}
	return true
}

func DelCfgFile() bool {
	// 删除配置文件：
	info := exec.Command("cmd", "/C", "del", "C:\\cfg.inf")
	if err := info.Run(); err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	return true
}

func ParsingFile() {
	// 打开配置文件解析文本
	// fi, err := os.Open("C:/cfg.inf")
	fi, err := os.Open("C:/cfg.inf")
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

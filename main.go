package main

import (
	"bufio"
	"capos"
	"fmt"
	"os"
	"outs"
)

const exportinfo = `
	1) 生成TXT数据
	2) 生成docx测评记录文档(Office2010版本以上打开)
	3) 同时生成以上

`

func main() {

	// 获取系统数据
	osdata := capos.CapOS()
	// PrintTxtTB格式化osdata输出到命令行界面，并返回字符串供 Outdocx输出到文本文件中
	s := outs.PrintTxtTB(osdata)

	fmt.Printf("%s\n\n", exportinfo)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Capos> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if input[0] == 10 {
			continue
		}
		instr := capos.ConvertToString(input, "gbk", "utf-8")
		instr = capos.RemoveEnt(instr)
		switch instr {
		case "1":
			fmt.Println("Capos> 文档已生成在当前目录！")
			outs.OutdateTxt(osdata.Systeminfo.ExportFileName, s)
		case "2":
			fmt.Println("Capos> 文档已生成在当前目录！")
			outs.Outdocx(osdata.Systeminfo.ExportFileName, outs.OutStd(osdata))
		case "3":
			fmt.Println("Capos> 文档已生成在当前目录！")
			outs.OutdateTxt(osdata.Systeminfo.ExportFileName, s)
			outs.Outdocx(osdata.Systeminfo.ExportFileName, outs.OutStd(osdata))
		default:
			fmt.Println("Capos> ")
			continue

		}

	}
}

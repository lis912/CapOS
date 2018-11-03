package outs

import (
	"capos"
	"fmt"
	"github.com/bndr/gotabulate"
	"os"
)

// 打印输出表格到命令行界面，并且返回界面内容到字符串
func PrintTxtTB(os capos.OsDate) string {
	// 打印口令策略表
	pwdtab := PrintPwdTb(os.PwdPolicy, os.VirMemPwd)
	pwdtab = fmt.Sprintf("%s\n\n%s", capos.BannerLogo, pwdtab)
	// totstr := fmt.Sprintf("%s\n\n", pwdtab)
	// 账户锁定策略
	totstr := fmt.Sprintf("%s\n\n%s", pwdtab, PrintLoginFalse())
	// 端口状态
	totstr = fmt.Sprintf("%s\n\n%s", totstr, Portstatetb())
	// 审计策略
	totstr = fmt.Sprintf("%s\n\n%s", totstr, Auditstatetb())
	// 其他单项
	totstr = fmt.Sprintf("%s\n\n%s", totstr, Othersingeltb())

	// 开始打印所有附录信息：
	// 获取本地用户状态
	totstr = fmt.Sprintf("%s\n\n%s", totstr, capos.PsysUser())
	// 打印systeminfo输出
	totstr = fmt.Sprintf("%s\n\n%s", totstr, capos.Psysteminfo())
	// 打印所有的补丁情况
	totstr = fmt.Sprintf("%s\n\n%s", totstr, capos.Pkbinfo())
	// 打印所有运行的服务
	totstr = fmt.Sprintf("%s\n\n%s", totstr, capos.PnowServer())
	// 打印所有运行进程
	totstr = fmt.Sprintf("%s\n\n%s", totstr, capos.Pnowpross())
	// 打印所有已经安装程序输出
	totstr = fmt.Sprintf("%s\n\n%s", totstr, capos.PinstallExe())

	return totstr
}

func PrintPwdTb(p capos.Passwd, s capos.VitmenRemove) string {
	pwd := FormaPwdTxt(p, s)
	row_1 := []interface{}{pwd[0].Nam, pwd[0].Ret, pwd[0].Deg, pwd[0].Com}
	row_2 := []interface{}{pwd[1].Nam, pwd[1].Ret, pwd[1].Deg, pwd[1].Com}
	row_3 := []interface{}{pwd[2].Nam, pwd[2].Ret, pwd[2].Deg, pwd[2].Com}
	row_4 := []interface{}{pwd[3].Nam, pwd[3].Ret, pwd[3].Deg, pwd[3].Com}
	row_5 := []interface{}{pwd[4].Nam, pwd[4].Ret, pwd[4].Deg, pwd[4].Com}
	row_6 := []interface{}{pwd[5].Nam, pwd[5].Ret, pwd[5].Deg, pwd[5].Com}
	// t := gotabulate.Create(pwd)
	t := gotabulate.Create([][]interface{}{row_1, row_2, row_3, row_4, row_5, row_6})
	t.Title = "系统口令策略"
	t.SetHeaders([]string{"测评项", "系统值", "符合度", "建议值"})
	t.SetAlign("left")
	fmt.Println(t.Render("grid"))
	return t.Render("grid")

}

func PrintLoginFalse() string {
	s := LoginFalseP()
	row_0 := []interface{}{s[0].Nam, s[0].Ret, s[0].Deg, s[0].Com}
	row_1 := []interface{}{s[1].Nam, s[1].Ret, s[1].Deg, s[1].Com}
	row_2 := []interface{}{s[2].Nam, s[2].Ret, s[2].Deg, s[2].Com}
	t := gotabulate.Create([][]interface{}{row_0, row_1, row_2})
	t.Title = "账户登录失败锁定策略"
	t.SetHeaders([]string{"测评项", "系统值", "符合度", "建议值"})
	t.SetAlign("left")
	fmt.Println(t.Render("grid"))
	return t.Render("grid")

}

func Portstatetb() string {
	s := Portsta()
	row_0 := []interface{}{s[0].Nam, s[0].Ret, s[0].Deg, s[0].Com}
	row_1 := []interface{}{s[1].Nam, s[1].Ret, s[1].Deg, s[1].Com}
	row_2 := []interface{}{s[2].Nam, s[2].Ret, s[2].Deg, s[2].Com}
	row_3 := []interface{}{s[3].Nam, s[3].Ret, s[3].Deg, s[3].Com}
	t := gotabulate.Create([][]interface{}{row_0, row_1, row_2, row_3})
	t.Title = "端口协议状态"
	t.SetHeaders([]string{"端口服务", "监听状态", "符合度", "建议值"})
	t.SetAlign("left")
	fmt.Println(t.Render("grid"))
	return t.Render("grid")
}

func Auditstatetb() string {
	s := Auditstate()
	row_0 := []interface{}{s[0].Nam, s[0].Ret, s[0].Deg, s[0].Com}
	row_1 := []interface{}{s[1].Nam, s[1].Ret, s[1].Deg, s[1].Com}
	row_2 := []interface{}{s[2].Nam, s[2].Ret, s[2].Deg, s[2].Com}
	row_3 := []interface{}{s[3].Nam, s[3].Ret, s[3].Deg, s[3].Com}
	row_4 := []interface{}{s[4].Nam, s[4].Ret, s[4].Deg, s[4].Com}
	row_5 := []interface{}{s[5].Nam, s[5].Ret, s[5].Deg, s[5].Com}
	row_6 := []interface{}{s[6].Nam, s[6].Ret, s[6].Deg, s[6].Com}
	row_7 := []interface{}{s[7].Nam, s[7].Ret, s[7].Deg, s[7].Com}
	row_8 := []interface{}{s[8].Nam, s[8].Ret, s[8].Deg, s[8].Com}
	t := gotabulate.Create([][]interface{}{row_0, row_1, row_2, row_3, row_4, row_5, row_6, row_7, row_8})
	t.Title = "系统审核策略"
	t.SetHeaders([]string{"审核项", "系统值", "符合度", "建议值"})
	t.SetAlign("left")
	fmt.Println(t.Render("grid"))
	return t.Render("grid")

}

func Othersingeltb() string {
	s := Othersingel()
	row_0 := []interface{}{s[0].Nam, s[0].Ret, s[0].Deg, s[0].Com}
	row_1 := []interface{}{s[1].Nam, s[1].Ret, s[1].Deg, s[1].Com}
	row_2 := []interface{}{s[2].Nam, s[2].Ret, s[2].Deg, s[2].Com}
	row_3 := []interface{}{s[3].Nam, s[3].Ret, s[3].Deg, s[3].Com}
	row_4 := []interface{}{s[4].Nam, s[4].Ret, s[4].Deg, s[4].Com}
	row_5 := []interface{}{s[5].Nam, s[5].Ret, s[5].Deg, s[5].Com}
	row_6 := []interface{}{s[6].Nam, s[6].Ret, s[6].Deg, s[6].Com}
	row_7 := []interface{}{s[7].Nam, s[7].Ret, s[7].Deg, s[7].Com}
	// row_8 := []interface{}{s[8].Nam, s[8].Ret, s[8].Deg, s[8].Com}
	t := gotabulate.Create([][]interface{}{row_0, row_1, row_2, row_3, row_4, row_5, row_6, row_7})
	t.Title = "其他选项"
	t.SetHeaders([]string{"测评项", "系统值", "符合度", "建议值"})
	t.SetAlign("left")
	fmt.Println(t.Render("grid"))
	return t.Render("grid")

}

func RetDeg(d int) string {
	if d == 1 {
		return "符合"
	} else {
		return "不符合"
	}
}

// 导出文本文件
func OutdateTxt(pathname string, data string) {
	pathname = fmt.Sprint(pathname, ".txt")
	fs, error := os.OpenFile(pathname, os.O_RDWR|os.O_CREATE, 0766)
	if error != nil {
		fmt.Println(error)
	}
	fs.WriteString(data)
	// fs.WriteString("\n\n\n\n\ntest！！！")
	fs.Close()
}

// func ReturnTable(OutDate []Singctr) string {
// 	// 空接口切片 一维切片 []interface{}{初始化}
// 	row_1 := []interface{}{OutDate[1].Nam, OutDate[1].Ctr, OutDate[1].Ret, OutDate[1].Deg}
// 	row_2 := []interface{}{OutDate[2].Nam, OutDate[2].Ctr, OutDate[2].Ret, OutDate[2].Deg}
// 	row_3 := []interface{}{OutDate[3].Nam, OutDate[3].Ctr, OutDate[3].Ret, OutDate[3].Deg}
// 	row_4 := []interface{}{OutDate[4].Nam, OutDate[4].Ctr, OutDate[4].Ret, OutDate[4].Deg}
// 	row_5 := []interface{}{OutDate[5].Nam, OutDate[5].Ctr, OutDate[5].Ret, OutDate[5].Deg}
// 	row_6 := []interface{}{"访问控制"}
// 	// row_3 := []interface{}{4435345, "dfas"}

// 	// Create an object from 2D interface array
// 	t := gotabulate.Create([][]interface{}{row_1, row_2, row_3, row_4, row_5, row_6})

// 	t.Title = "等级保护测评"

// 	// 是否要隐藏表格顶行 "top"
// 	// t.HideLines = append(t.HideLines, "belowheader")

// 	// Set the Headers (optional)
// 	t.SetHeaders([]string{"控制项", "测评结果", "符合度(仅做参考)"})

// 	// Set the Empty String (optional)
// 	// t.SetEmptyString("lishichnag")

// 	// Set Align (Optional)  center  left
// 	t.SetAlign("left")

// 	// Print the result: grid, or simple
// 	// fmt.Println(t.Render("grid"))
// 	// fmt.Println("\n\n")

// 	return t.Render("grid")
// }

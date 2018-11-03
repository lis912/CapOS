package capos

import "fmt"

const BannerLogo = `
    ___             ___  ____                                
  / ___|__ _ _ __  / _ \/ ___| 
 | |   / _* | '_ \| | | \___ \ 
 | |__| (_| | |_) | |_| |___) |
  \____\__,_| .__/ \___/|____/ 
            |_|                 Version v1.3 for Windows System x64
            		        2018.8         Author： Li Shichang

            		        提示：请 ‘右键->以管理员身份运行’ 本工具！
--------------------------------------------------------------------------------+
`

func PrintLogo() {
	fmt.Println(BannerLogo)
	fmt.Println(CaposUsage)

}

// Author： Li shichang

const CaposUsage = `
	正在获取系统数据......

`

// usage:
// 	os 		 [>] [filename.docx | filenam.txt] 	 检测windows系统或导出结果
// 	oracle
// 	mysql
// 	mssql
// 	exit										 退出工具

// These are common CapOS commands used in various situations:

// 	CapOS -v				View version information
// 	CapOS -h				Ask for help
// --------------------------------------------------------------------------------+
// `

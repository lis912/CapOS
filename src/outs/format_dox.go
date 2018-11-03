package outs

import (
	"capos"
	"fmt"
	"strings"
)

type Singctr struct {
	Nam string // 控制项
	Ctr string // 测评项
	Ret string // 结果记录
	Deg string // 整体符合度
	Com string // 整改建议
	Rem string // 备注
}

// var OutDate []Singctr

func OutStd(os capos.OsDate) (OutDate []Singctr) {

	var s Singctr
	{
		s.Nam = "测评指标"
		s.Ctr = "控制项"
		s.Deg = "符合情况"
		s.Ret = "结果记录"
		OutDate = append(OutDate, s)
	}
	{
		tmp := os.LogPolicy
		s.Nam = "身份鉴别"
		s.Ctr = "a)应对登录操作系统和数据库系统的用户进行身份标识和鉴别；"
		s.Deg = capos.CalConDeg(tmp.Totdeg, tmp.Mindeg, tmp.Maxdeg)
		s.Ret = fmt.Sprintf("1)%s“要使用本机，用户必须输入用户名和密码”, 本项%s；", tmp.IsUse.Val, s.Deg)
		OutDate = append(OutDate, s)
	}

	{
		tmp2 := os.PwdPolicy
		s.Nam = ""
		s.Ctr = "b)操作系统和数据库系统管理用户身份标识应具有不易被冒用的特点，口令应有复杂度要求并定期更换；"
		s.Deg = capos.CalConDeg(tmp2.Totdeg, tmp2.Mindeg, tmp2.Maxdeg)
		s.Ret = fmt.Sprintf("经查看，1)密码必须符合复杂性要求：%s；2)密码最长最小值%s；3)密码最长使用期限%s天；4)密码最短使用期限%s天；5)强制密码历史%s个,密码策略%s；", tmp2.Comple.Val, tmp2.Miniwd.Val, tmp2.Agewd.Val, tmp2.Longwd.Val, tmp2.Hisywd.Val, s.Deg)
		// s.Ret = `经查看，1)密码必须符合复杂性要求：启用；2)密码最长最小值8；3)密码最长使用期限2天；4)密码最短使用期限2天；5)强制密码历史24个,密码策略符合`

		OutDate = append(OutDate, s)
	}

	{
		tmp3 := os.LogFaPicy
		s.Nam = ""
		s.Ctr = "c)应启用登录失败处理功能，可采取结束会话、限制非法登录次数和自动退出等措施；"
		s.Deg = capos.CalConDeg(tmp3.Totdeg, tmp3.Mindeg, tmp3.Maxdeg)
		s.Ret = fmt.Sprintf("经查看：1)“复位账户锁定计数器”：%s次; 2)“账户锁定时间”:%s分钟; 3)“账户锁定阈值”:%s次无效登录，本项%s", tmp3.LockTime.Val, tmp3.ResTime.Val, tmp3.LockCout.Val, s.Deg)
		OutDate = append(OutDate, s)
	}

	{
		tmp4 := os.RemoMan
		s.Nam = ""
		s.Ctr = "d)当对服务器进行远程管理时，应采取必要措施，防止鉴别信息在网络传输过程中被窃听；"
		s.Deg = capos.CalConDeg(tmp4.Totdeg, tmp4.Mindeg, tmp4.Maxdeg)
		s.Ret = fmt.Sprintf("经查看：1)系统%s了Telnet服务; 2)%s了3389远程桌面加密服务，本项%s", tmp4.Telnet.Val, tmp4.WinDesk.Val, s.Deg)
		OutDate = append(OutDate, s)
	}

	{
		tmp5 := os.LocUser
		s.Nam = ""
		s.Ctr = "e)应为操作系统和数据库系统的不同用户分配不同的用户名，确保用户名具有唯一性；"
		s.Deg = capos.CalConDeg(tmp5.Totdeg, tmp5.Mindeg, tmp5.Maxdeg)
		s.Ret = fmt.Sprintf("1)查看系统本地用户列表：%s, 不存在重复账户，本项%s；", tmp5.Username.Val, s.Deg)
		OutDate = append(OutDate, s)
	}

	{
		// tmp6 := os.LocUser
		s.Nam = ""
		s.Ctr = "f)应使用两种或两种以上组合的鉴别技术对管理用户进行身份鉴别;"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		tmp7 := os.Access
		s.Nam = "访问控制"
		s.Ctr = "a)应启用访问控制功能，依据安全策略控制用户对资源的访问；"
		s.Deg = capos.CalConDeg(tmp7.Totdeg, tmp7.Mindeg, tmp7.Maxdeg)
		s.Ret = fmt.Sprintf("经查看，1)普通用户组对system32系统重要目录具有%s; 2)通过命令‘net share’查看系统默认共享状态为：%s 本项%s", tmp7.System32.Val, tmp7.Netshare.Val, s.Deg)
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "b)应根据管理用户的角色分配权限，实现管理用户的权限分离，仅授予管理用户所需的最小权限；"
		s.Deg = ""
		s.Ret = "须查看本地策略->用户权限分配"
		OutDate = append(OutDate, s)
	}

	{

		s.Nam = ""
		s.Ctr = "c)应实现操作系统和数据库系统特权用户的权限分离；"
		s.Deg = ""
		s.Ret = "请查看附录"
		OutDate = append(OutDate, s)
	}

	{
		tmp10 := os.DefaUse
		s.Nam = ""
		s.Ctr = "d)应限制默认账户的访问权限，重命名系统默认账户，修改这些账户的默认口令；"
		s.Deg = capos.CalConDeg(tmp10.Totdeg, tmp10.Mindeg, tmp10.Maxdeg)
		s.Ret = fmt.Sprintf("经访谈查看，1)系统%s重命名系统账户；2)系统%s修改账户默认口令；3)系统%s禁用Guest来宾账户；本项%s。", tmp10.RenameAdmin.Val, tmp10.Nodifypwd.Val, tmp10.DisableGuset.Val, s.Deg)
		OutDate = append(OutDate, s)
	}

	{
		tmp11 := os.Nousers
		s.Nam = ""
		s.Ctr = "e)应及时删除多余的、过期的账户，避免共享账户的存在;"
		s.Deg = capos.CalConDeg(tmp11.Totdeg, tmp11.Mindeg, tmp11.Maxdeg)
		s.Ret = fmt.Sprintf("1)经查看本地用户列表：%s;系统管理员组类表: %s,无多余的、过期的账户。本项%s。", tmp11.LocakUses.Val, tmp11.AdminGroudUses.Val, s.Deg)
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "f)应对重要信息资源设置敏感标记；"
		s.Deg = ""
		s.Ret = "访谈管理人员"
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "g)应依据安全策略严格控制用户对有敏感标记重要信息资源的操作。"
		s.Deg = ""
		s.Ret = "访谈管理人员"
		OutDate = append(OutDate, s)
	}

	{
		tmp14 := os.Audit
		s.Nam = "安全审计"
		s.Ctr = "a)审计范围应覆盖到服务器和重要客户端上的每个操作系统用户和数据库用户；"
		s.Deg = capos.CalConDeg(tmp14.Totdeg, tmp14.Mindeg, tmp14.Maxdeg)
		if tmp14.Totdeg == 0 {
			s.Ret = fmt.Sprintf("1)经查看系统审核策略，没有开启安全审核功能，本项%s。", s.Deg)
		} else {
			s.Ret = fmt.Sprintf("经查看，系统审核策略：1)审核策略更改:%s; 2)审核登录事件:%s; 3)审核对象访问:%s; 4)审核进程跟踪:%s; 5)审核目录服务访问:%s; 6)审核特权使用:%s; 7)审核系统事件:%s; 8)审核账户登录事件:%s; 9)审核账户管理%s,本项%s。", tmp14.PolicyChange.Val, tmp14.LogonEvents.Val, tmp14.ObjectAccess.Val, tmp14.ProcessTracking.Val, tmp14.DSAccess.Val, tmp14.PrivilegeUse.Val, tmp14.SystemEvents.Val, tmp14.AccountLogon.Val, tmp14.AccountManage.Val, s.Deg)
		}
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "b)审计内容应包括重要用户行为、系统资源的异常使用和重要系统命令的使用等系统内重要的安全相关事件；"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "c)审计记录应包括事件的日期、时间、类型、主体标识、客体标识和结果等；"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "d)应能够根据记录数据进行分析，并生成审计报表；"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "e)应保护审计进程，避免受到未预期的中断。"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "f)应保护审计记录，避免受到未预期的删除、修改或覆盖"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		tmp20 := os.NoLastname
		s.Nam = "剩余信息保护"
		s.Ctr = "a)应保证操作系统和数据库管理系统用户的鉴别信息所在的存储空间，被释放或再分配给其他用户前得到完全清除，无论这些信息是存放在硬盘上还是在内存中；"
		s.Deg = capos.CalConDeg(tmp20.Totdeg, tmp20.Mindeg, tmp20.Maxdeg)
		s.Ret = fmt.Sprintf("1)经查看本地策略-安全选项,“交互式登录：不显示最后的用户名”：%s，本项%s。", tmp20.NoShowLastname.Val, s.Deg)
		OutDate = append(OutDate, s)
	}

	{
		tmp21 := os.VirMemPwd
		s.Nam = ""
		s.Ctr = "b)应确保系统内的文件、目录和数据库记录等资源所在的存储空间，被释放或重新分配给其他用户前得到完全清除"
		s.Deg = capos.CalConDeg(tmp21.Totdeg, tmp21.Mindeg, tmp21.Maxdeg)
		s.Ret = fmt.Sprintf("1)经查看本地策略-安全选项,“关机前清除虚拟内存页面”：%s; 2)查看系统密码策略，“用可还原的加密来储存密码”:%s, 本项%s。", tmp21.Virmenremove.Val, tmp21.Pwdremove.Val, s.Deg)
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = "入侵防范"
		s.Ctr = "a)应能够检测到对重要服务器进行入侵的行为，能够记录入侵的源IP、攻击类型，攻击目的，攻击时间，并在发生严重入侵事件时提供报警；"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "b)应能够对重要程序完整性进行检测，并在检测到完整性受到破坏后具有恢复措施；"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		tmp24 := os.Patch
		s.Nam = ""
		s.Ctr = "c)操作系统遵循最小安装的原则，仅安装需要的组件和应用程序，并通过设置升级服务器等方式保持系统补丁及时得到更新"
		s.Deg = capos.CalConDeg(tmp24.Totdeg, tmp24.Mindeg, tmp24.Maxdeg)
		if tmp24.Patches.Val == "null" {
			s.Ret = fmt.Sprintf("1)系统安装的组件和应用程序遵循了最小安装的原则； 2)系统已安装补丁列表为空；3)系统%s“Windows Update” 4)端口监听状态：%s;本项%s", tmp24.Updata.Val, tmp24.RiskPort.Val, s.Deg)
		} else {
			kbinfo := strings.Split(tmp24.Patches.Val, ":")
			s.Ret = fmt.Sprintf("1)系统安装的组件和应用程序遵循了最小安装的原则； 2)查看系统已安装补丁列表，最近更新补丁%s安装时间为%s；3)系统%s“Windows Update”自动更新服务； 4)开启了不必要的端口%s;本项%s", kbinfo[0], kbinfo[1], tmp24.Updata.Val, tmp24.RiskPort.Val, s.Deg)
		}
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = "恶意代码防范"
		s.Ctr = "a)应安装防恶意代码软件，并及时更新防恶意代码软件版本和恶意代码库；"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "b)主机防恶意代码产品应具有与网络恶意代码产品不同的恶意代码库；"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "c)应支持防恶意代码软件的统一管理。"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = "资源控制"
		s.Ctr = "a)应通过设定终端接入方式、网络地址范围等条件限制终端登录；"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		tmp29 := os.LoginTimeoutLock
		s.Nam = ""
		s.Ctr = "b)应根据安全策略设置登录终端的操作超时锁定；"
		s.Deg = capos.CalConDeg(tmp29.Totdeg, tmp29.Mindeg, tmp29.Maxdeg)
		if tmp29.TimeoutLock.Val == "null" {
			s.Ret = fmt.Sprintf("1)经查看，系统未设置屏幕保护程序，本项%s。", s.Deg)
		} else {
			s.Ret = fmt.Sprintf("1)经查看，系统已设置屏幕超时锁定：%s分钟，本项%s。", tmp29.TimeoutLock.Val, s.Deg)
		}

		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "c)应对重要服务器进行监视，包括监视服务器的CPU、硬盘、内存、网络等资源的使用情况；"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "d)应限制单个用户对系统资源的最大或最小使用限度;"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}

	{
		s.Nam = ""
		s.Ctr = "e)应能够对系统的服务水平降低到预先规定的最小值进行检测和报警。"
		s.Deg = ""
		s.Ret = ""
		OutDate = append(OutDate, s)
	}
	return OutDate

}

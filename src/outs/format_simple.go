package outs

import (
	"capos"
	"fmt"
)

var showos capos.OsDate

func init() {
	firstexe()
}

func firstexe() {
	showos = capos.CapOS()
}

// 本模块输出的总数据结构
// type OsDate struct {
// 	LogPolicy UseIdPwd  // 系统登录方式
// 	PwdPolicy Passwd    // 口令强度策略
// 	LogFaPicy Logfaile  // 登录失败策略
// 	RemoMan   RemoManag // 远程管理方式
// 	LocUser   Userinfo  // 检查重复账户

// 	Access  Accesctr         // 访问权限和共享
// 	DefaUse DefaultUseAccess // 关于默认账户的访问权限
// 	Nousers NouSers          // 无多余重复的用户

// 	Audit AuditPolicy // 审计策略

// 	NoLastname InfoProtion      // 剩余信息保护
// 	VirMemPwd  VitmenRemove     // 剩余信息保护2
// 	Patch      PatchPortUpgrade // 补丁情况和端口情况自动升级

// 	LoginTimeoutLock Screenlock // 屏幕超时锁定

// 	Systeminfo Sysinfos // 系统软硬信息
// }

// 系统审计
func Auditstate() (PwdTable []Singctr) {
	var s Singctr
	tmp := showos.Audit

	{
		s.Nam = "审核策略更改"
		s.Ret = tmp.PolicyChange.Val
		s.Deg = RetDeg(tmp.PolicyChange.Deg)
		s.Com = "成功，失败"
		PwdTable = append(PwdTable, s)

		s.Nam = "审核登录事件"
		s.Ret = tmp.LogonEvents.Val
		s.Deg = RetDeg(tmp.LogonEvents.Deg)
		s.Com = "成功，失败"
		PwdTable = append(PwdTable, s)

		s.Nam = "审核对象访问"
		s.Ret = tmp.ObjectAccess.Val
		s.Deg = RetDeg(tmp.ObjectAccess.Deg)
		s.Com = "成功，失败"
		PwdTable = append(PwdTable, s)

		s.Nam = "审核进程跟踪"
		s.Ret = tmp.ProcessTracking.Val
		s.Deg = RetDeg(tmp.ProcessTracking.Deg)
		s.Com = "成功，失败"
		PwdTable = append(PwdTable, s)

		s.Nam = "审核目录服务访问"
		s.Ret = tmp.DSAccess.Val
		s.Deg = RetDeg(tmp.DSAccess.Deg)
		s.Com = "成功，失败"
		PwdTable = append(PwdTable, s)

		s.Nam = "审核特权使用"
		s.Ret = tmp.PrivilegeUse.Val
		s.Deg = RetDeg(tmp.PrivilegeUse.Deg)
		s.Com = "成功，失败"
		PwdTable = append(PwdTable, s)

		s.Nam = "审核系统事件"
		s.Ret = tmp.SystemEvents.Val
		s.Deg = RetDeg(tmp.SystemEvents.Deg)
		s.Com = "成功，失败"
		PwdTable = append(PwdTable, s)

		s.Nam = "审核账户登录事件"
		s.Ret = tmp.AccountLogon.Val
		s.Deg = RetDeg(tmp.AccountLogon.Deg)
		s.Com = "成功，失败"
		PwdTable = append(PwdTable, s)

		s.Nam = "审核账户管理"
		s.Ret = tmp.AccountManage.Val
		s.Deg = RetDeg(tmp.AccountManage.Deg)
		s.Com = "成功，失败"
		PwdTable = append(PwdTable, s)
	}

	return PwdTable
}

// 端口和服务状态
func Portsta() (PwdTable []Singctr) {
	var s Singctr
	tmp := showos.RemoMan
	// tmp2 := showos.Patch
	{
		s.Nam = "23 Telnet协议"
		s.Ret = tmp.Telnet.Val
		s.Deg = RetDeg(tmp.Telnet.Deg)
		s.Com = "禁用"
		PwdTable = append(PwdTable, s)

		s.Nam = "3389 远程桌面"
		s.Ret = tmp.WinDesk.Val
		s.Deg = RetDeg(tmp.WinDesk.Deg)
		s.Com = "非服务器设备需禁用"
		PwdTable = append(PwdTable, s)

		s.Nam = "445 局域网共享"
		if capos.PortState("445") {
			s.Ret = "开启"
			s.Deg = "不符合"
		} else {
			s.Ret = "关闭"
			s.Deg = "符合"
		}
		s.Com = "禁用"
		PwdTable = append(PwdTable, s)

		s.Nam = "135 RPC远程"
		if capos.PortState("135") {
			s.Ret = "开启"
			s.Deg = "不符合"
		} else {
			s.Ret = "关闭"
			s.Deg = "符合"
		}
		s.Com = "禁用"
		PwdTable = append(PwdTable, s)
	}

	return PwdTable
}

// 登录失败策略表
func LoginFalseP() (PwdTable []Singctr) {
	var s Singctr
	tmp := showos.LogFaPicy
	{
		s.Nam = "账户锁定时间"
		s.Ret = tmp.LockTime.Val
		s.Deg = RetDeg(tmp.LockTime.Deg)
		s.Com = ">30分钟"
		PwdTable = append(PwdTable, s)

		s.Nam = "账户锁定阈值"
		s.Ret = tmp.LockCout.Val
		s.Deg = RetDeg(tmp.LockCout.Deg)
		s.Com = "3~5次"
		PwdTable = append(PwdTable, s)

		s.Nam = "重置账户锁定计数器"
		s.Ret = tmp.ResTime.Val
		s.Deg = RetDeg(tmp.ResTime.Deg)
		s.Com = "30分钟之后"
		PwdTable = append(PwdTable, s)
	}

	return PwdTable
}

// 登录失败策略表
func Othersingel() (PwdTable []Singctr) {
	var s Singctr
	// tmp := showos.LogFaPicy
	{
		s.Nam = "是否勾选登录必须输入(用户名+密码)"
		s.Ret = showos.LogPolicy.IsUse.Val
		s.Deg = RetDeg(showos.LogPolicy.IsUse.Deg)
		s.Com = "建议勾选"
		PwdTable = append(PwdTable, s)

		s.Nam = "系统默认共享"
		s.Ret = showos.Access.Netshare.Val
		s.Deg = RetDeg(showos.Access.Netshare.Deg)
		s.Com = "禁用共享"
		PwdTable = append(PwdTable, s)

		s.Nam = "system32目录User组访问权限"
		s.Ret = showos.Access.System32.Val
		s.Deg = RetDeg(showos.Access.System32.Deg)
		s.Com = "限制权限"
		PwdTable = append(PwdTable, s)

		s.Nam = "是否已重命名Administrator账户"
		s.Ret = fmt.Sprintf("%s重命名", showos.DefaUse.RenameAdmin.Val)
		s.Deg = RetDeg(showos.DefaUse.RenameAdmin.Deg)
		s.Com = "重命名"
		PwdTable = append(PwdTable, s)

		s.Nam = "是否已禁用Guest来宾账户"
		s.Ret = fmt.Sprintf("%s禁用", showos.DefaUse.DisableGuset.Val)
		s.Deg = RetDeg(showos.DefaUse.DisableGuset.Deg)
		s.Com = "禁用"
		PwdTable = append(PwdTable, s)

		s.Nam = "关机前清除虚拟内存页面"
		s.Ret = showos.VirMemPwd.Virmenremove.Val
		s.Deg = RetDeg(showos.VirMemPwd.Virmenremove.Deg)
		s.Com = "启用"
		PwdTable = append(PwdTable, s)

		s.Nam = "不显示最后的用户名"
		s.Ret = showos.NoLastname.NoShowLastname.Val
		s.Deg = RetDeg(showos.NoLastname.NoShowLastname.Deg)
		s.Com = "启用"
		PwdTable = append(PwdTable, s)

		s.Nam = "屏幕保护锁定时间"
		if showos.LoginTimeoutLock.TimeoutLock.Val == "null" {
			s.Ret = "未开启"
		} else {
			s.Ret = fmt.Sprintf("%s 分钟", showos.LoginTimeoutLock.TimeoutLock.Val)
		}
		s.Deg = RetDeg(showos.LoginTimeoutLock.TimeoutLock.Deg)
		s.Com = ">15分钟"
		PwdTable = append(PwdTable, s)

	}

	return PwdTable
}

func FormaPwdTxt(tmp capos.Passwd, tmp2 capos.VitmenRemove) (PwdTable []Singctr) {
	var s Singctr
	{
		// tmp := os.PwdPolicy
		s.Nam = "密码必须符合复杂性要求"
		s.Ret = tmp.Comple.Val
		s.Deg = RetDeg(tmp.Comple.Deg)
		s.Com = "启用"
		PwdTable = append(PwdTable, s)

		s.Nam = "密码最长最小值"
		s.Ret = tmp.Miniwd.Val
		s.Deg = RetDeg(tmp.Miniwd.Deg)
		s.Com = "8"
		PwdTable = append(PwdTable, s)

		s.Nam = "密码最长使用期限"
		s.Ret = tmp.Agewd.Val
		s.Deg = RetDeg(tmp.Agewd.Deg)
		s.Com = "70"
		PwdTable = append(PwdTable, s)

		s.Nam = "密码最短使用期限"
		s.Ret = tmp.Longwd.Val
		s.Deg = RetDeg(tmp.Longwd.Deg)
		s.Com = "2"
		PwdTable = append(PwdTable, s)

		s.Nam = "强制密码历史"
		s.Ret = tmp.Hisywd.Val
		s.Deg = RetDeg(tmp.Hisywd.Deg)
		s.Com = "24"
		PwdTable = append(PwdTable, s)

		s.Nam = "用可还原的加密来储存密码"
		s.Ret = tmp2.Pwdremove.Val
		s.Deg = RetDeg(tmp2.Pwdremove.Deg)
		s.Com = "禁用"
		PwdTable = append(PwdTable, s)

	}
	return PwdTable
}

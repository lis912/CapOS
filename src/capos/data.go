package capos

const (
	Y       = 1  // 符合
	N       = 0  // 不符合
	NoApply = 10 // 不适用
)

// 单条数据结构
type Term struct {
	Val string
	Deg int
}

// 密码口令策略
type Passwd struct {
	Comple Term
	Miniwd Term
	Agewd  Term
	Longwd Term
	Hisywd Term
	Totdeg int
	Mindeg int
	Maxdeg int
}

// 登录失败锁定策略
type Logfaile struct {
	LockTime Term
	LockCout Term
	ResTime  Term
	Totdeg   int
	Mindeg   int
	Maxdeg   int
}

// 勾选登录方式
type UseIdPwd struct {
	IsUse  Term
	Totdeg int
	Mindeg int
	Maxdeg int
}

// 远程管理方式： 检查 23 3389 端口
type RemoManag struct {
	Telnet  Term
	WinDesk Term
	Totdeg  int
	Mindeg  int
	Maxdeg  int
}

// 系统用户，是否有重复用户
type Userinfo struct {
	Username Term
	Adminame Term
	Totdeg   int
	Mindeg   int
	Maxdeg   int
}
type DoubleFactor struct {
}

// 本模块输出的总数据结构
type OsDate struct {
	LogPolicy UseIdPwd  // 系统登录方式
	PwdPolicy Passwd    // 口令强度策略
	LogFaPicy Logfaile  // 登录失败策略
	RemoMan   RemoManag // 远程管理方式
	LocUser   Userinfo  // 检查重复账户

	Access  Accesctr         // 访问权限和共享
	DefaUse DefaultUseAccess // 关于默认账户的访问权限
	Nousers NouSers          // 无多余重复的用户

	Audit AuditPolicy // 审计策略

	NoLastname InfoProtion      // 剩余信息保护
	VirMemPwd  VitmenRemove     // 剩余信息保护2
	Patch      PatchPortUpgrade // 补丁情况和端口情况自动升级

	LoginTimeoutLock Screenlock // 屏幕超时锁定

	Systeminfo Sysinfos // 系统软硬信息
}

// 本模块总输出函数
func CapOS() (os OsDate) {
	os.LogPolicy = IsUseIdPwd()
	os.PwdPolicy = PaswdPolicy()
	os.LogFaPicy = LogFailCheck()
	os.RemoMan = RemoManCheck()
	os.LocUser = UsersCheck()

	os.Access = AccesctrCheck()
	os.DefaUse = DefaultUseAccessCheck()
	os.Nousers = NousersCheck()
	os.Audit = AuditCheck()

	os.NoLastname = NoLastnameCheck()
	os.VirMemPwd = VirMemPwdCheck()
	os.Patch = PatchCheck()

	os.LoginTimeoutLock = LoginTimeoutLockCheck()

	os.Systeminfo = SysinfosCheck()
	return os
}

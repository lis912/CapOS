package main

import (
	"capos"
)

type Singctr struct {
	Nam string
	Ctr string
	Ret string
	Fra string
}

// var Docdate []Singctr

func OutRet(os capos.OsDate) (s []Singctr) {
	tmp := os.LogPolicy.
	s.Nam = "身份鉴别"
	s.Ctr = "a)应对登录操作系统和数据库系统的用户进行身份标识和鉴别；"
	s.Ret = "1)已勾选“要使用本机，用户必须输入用户名和密码”, 本项符合；"
	s.Fra = "符合(5分)"
	Docdate = append(Docdate, s)

	s.Nam = ""
	s.Ctr = "b)操作系统和数据库系统管理用户身份标识应具有不易被冒用的特点，口令应有复杂度要求并定期更换；"
	s.Ret = "经查看，1)密码必须符合复杂性要求：启用；2)密码最长最小值8；3)密码最长使用期限90天；4)密码最短使用期限1天；5)强制密码历史5个；6)密码策略符合要求；"
	s.Fra = "符合(5分)"
	Docdate = append(Docdate, s)

	s.Ctr = "c)应启用登录失败处理功能，可采取结束会话、限制非法登录次数和自动退出等措施；"
	s.Ret = "查看账户锁定策略：1)复位账户锁定计数器：5分钟后；2)账户锁定时间：5分钟；3)账户锁定阑值：6次无效登录；本项符合；"
	s.Fra = "符合(5分)"
	Docdate = append(Docdate, s)

	s.Ctr = "d)当对服务器进行远程管理时，应采取必要措施，防止鉴别信息在网络传输过程中被窃听；"
	s.Ret = "1)系统禁用了telent服务；2)采取了3389远程桌面服务，并通过堡垒机登录维护；3)本项符合；"
	s.Fra = "符合(5分)"
	Docdate = append(Docdate, s)

	s.Ctr = "e)应为操作系统和数据库系统的不同用户分配不同的用户名，确保用户名具有唯一性；"
	s.Ret = "1)不存在多用户使用同一系统账户Administrator的情况，本项符合；"
	s.Fra = "符合(5分)"
	Docdate = append(Docdate, s)

	s.Ctr = "f)应采用两种或两种以上组合的鉴别技术对管理用户进行身份鉴别;"
	s.Ret = "1)只有‘用户名+密码’一种鉴别技术对管理用户进行身份鉴别；本项不符合;"
	s.Fra = "不符合"
	Docdate = append(Docdate, s)

	// 访问控制
	// var j Singctr
	s.Nam = "访问控制"
	s.Ctr = "a)应启用访问控制功能，依据安全策略控制用户对资源的访问；"
	s.Ret = "经查看，1)系统开启了IPC默认共享；2.system重要的目录只有管理员组有操作权限；"
	s.Fra = "部分符合(3分)"
	Docdate = append(Docdate, s)

	return s
}

package MySQL

import "errors"

var (
	ErrorUserExit 	   = errors.New("用户已存在")
	ErrorInvalidID     = errors.New("无效的ID")
	ErrorGenIDFailed   = errors.New("创建用户ID失败")
	ErrorQueryFailed   = errors.New("查询数据失败")
	ErrorUserNotExit   = errors.New("用户不存在")
	ErrorInsertFailed  = errors.New("插入数据失败")
	ErrorDeleteFailed  = errors.New("删除数据失败")
	ErrorPasswordWrong = errors.New("密码错误")
)

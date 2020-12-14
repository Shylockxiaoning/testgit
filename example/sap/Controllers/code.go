package Controllers

type MyCode int64
const (
	CodeSuccess       	   		 MyCode = 1000
	CodeUserExist     	   		 MyCode = 1001
	CodeUserNotExist         	 MyCode = 1002
	CodeRegisterSuccess    		 MyCode = 1003
	CodeServerBusy		  		 MyCode = 1004
	CodeInvalidCookies     	 	 MyCode = 1005
	CodeInvalidNoCookies  	 	 MyCode = 1006
	CodeConfirmPasswordfailed 	 MyCode = 1007
)
var msgFlags = map[MyCode]string{
	CodeSuccess:         		 "登陆成功",
	CodeUserExist:       	  	 "用户名重复",
	CodeUserNotExist:  	   		 "用户不存在",
	CodeRegisterSuccess:  		 "注册成功",
	CodeServerBusy:				 "服务繁忙",
	CodeInvalidCookies:    		 "无效的Cookie",
	CodeInvalidNoCookies:  		 "缺少Cookie值",
	CodeConfirmPasswordfailed:	 "两次密码不一致",
}
func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
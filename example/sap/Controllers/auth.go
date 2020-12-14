package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sap/Model"
)


//CheckAuth 检查cookie值
func CheckAuth(c *gin.Context){
	ms,err:=c.Request.Cookie("cookie")
	//获取cookie值错误
	if err != nil{
		ResponseError(c,CodeInvalidCookies)
		zap.L().Error("request cookie failed",zap.Error(err))
		c.Abort()
		return
	}
	//cookie值为空时
	if ms == nil{
		ResponseErrorWithMsg(c,CodeInvalidNoCookies,"缺少Cookie值")
		c.Abort()
		return
	}
	//cookie值不匹配时
	if ms.Value != "thiscookies"{
		ResponseErrorWithMsg(c,CodeInvalidCookies,"无效的Cookie值")
		zap.L().Error("ms cookie failed",zap.Error(err))
		c.Abort()
		return
	}
	//当cookie时间小于十秒时，重新设置
	if ms.MaxAge<=10{
		c.SetCookie("cookie","thiscookies",300,"","",false,true)
	}
	c.Next()
}

//Admin() 检查权限
func Admin(c *gin.Context){
	var user Model.User
	s:=user.Position
	fmt.Println("sfiafhhf",s)
	c.Next()
}

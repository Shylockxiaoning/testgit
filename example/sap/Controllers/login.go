package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sap/Dao/MySQL"
	"sap/Model"
)

// LoinHandler 用户登陆 Json
func LoinHandler(c *gin.Context){
	var user Model.User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		zap.L().Error("ShouldBindJSON failed",zap.Error(err))
		return
	}
	if err:=MySQL.Login(&user);err!=nil{
		zap.L().Error("ConfirmUser failed",zap.Error(err))
		if err==MySQL.ErrorUserNotExit{
			c.Redirect(http.StatusMovedPermanently,"register")
			return
		}
		ResponseErrorWithMsg(c,CodeUserNotExist,fmt.Sprint(err))
		return
	}
	//设置cookie值
	c.SetCookie("cookie","thiscookies",300,"","",false,true)
	c.Redirect(http.StatusMovedPermanently,"/admin/index")
}

//LoginHandler 用户登陆 Html
func LoginHandler(c *gin.Context){
	var user Model.User
	user.Username=c.PostForm("username")
	user.Password=c.PostForm("password")
	if err:=MySQL.Login(&user);err!=nil {
		zap.L().Error("ConfirmUser failed",zap.Error(err))
		if err==MySQL.ErrorUserNotExit{
			c.Redirect(http.StatusMovedPermanently,"register")
			return
		}
		ResponseErrorWithMsg(c,CodeUserNotExist,fmt.Sprint(err))
		return
	}
	c.Header("position",user.Position)
	c.SetCookie("cookie","thiscookies",300,"","",false,true)
	c.Redirect(http.StatusMovedPermanently,"/admin/index")
}

//Register 用户注册
func Register(c *gin.Context){
	var admin Model.Register
	if err:=c.ShouldBindJSON(&admin);err!=nil{
		zap.L().Error("ShouldBindJSON failed",zap.Error(err))
		return
	}
	if admin.Password != admin.ConfirmPassword{
		ResponseErrorWithMsg(c,CodeConfirmPasswordfailed,"两次密码不一致")
		return
	}
	if err:=MySQL.Register(&admin);err!=nil{
		zap.L().Error("Register user failed",zap.Error(err))
		return
	}
}


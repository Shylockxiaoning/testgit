package MySQL

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sap/Model"
)

//encryptPassword 数据加盐 并进行MD5转换
func encryptPassword(data[]byte)(result string){
	s:=md5.New()
	s.Write([]byte(viper.GetString("mysql.secret")))
	result =hex.EncodeToString(s.Sum(data))
	return result
}
//Login 数据库用户密码校验
func Login(user *Model.User) (err error) {
	userpassword:=user.Password
	s:= encryptPassword([]byte(userpassword))
	sqlstr:="select username,password,position from admin where username = ?"
	err = db.Get(user,sqlstr,user.Username)
	if err!=nil && err!=sql.ErrNoRows{
		zap.L().Error("err",zap.Error(err))
		return ErrorQueryFailed
	}
	if err == sql.ErrNoRows{
		return ErrorUserNotExit
	}
	if s != user.Password{
		return ErrorPasswordWrong
	}
	return
}

//Register 用户注册
func Register(user *Model.Register)(err error){
	if CheckUsername(user.Username) == true{
		zap.L().Error("账号已存在",zap.Error(err))
		return
	}
	switch user.Position {
	case "secreat":user.Position="组长"
	case "root":user.Position="店长"
	default:
		user.Position="操作员"
	}
	password:=user.Password
	password=encryptPassword([]byte(password))
	sqlstr:="insert into admin (username, password,position) values (?,?,?)"
	_, err = db.Exec(sqlstr,user.Username,password,user.Position)
	if err!=nil{
		zap.L().Error("insert into mysql failed",zap.Error(err))
	}
	return
}

//CheckUsername 检查用户名是否重复
func CheckUsername(username string)bool {
	sqlstr:="select count(username) from admin where username = ?"
	var count int64
	err:=db.Get(&count,sqlstr,username)
	if err != nil && err != sql.ErrNoRows{
		zap.L().Error("query user exist failed", zap.Error(err))
		return true
	}
	return count > 0
}
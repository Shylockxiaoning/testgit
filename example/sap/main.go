package main

import (
	"go.uber.org/zap"
	"sap/Controllers"
	"sap/Dao/MySQL"

	// "sap/Dao/MySQL"
	logger "sap/Logger"
	"sap/Router"
)

//程序入口
func main(){
	//Init vipper zap mysql
	logger.Initlog()	//初始化日志库
	zap.L().Info("init logerr success")
	defer zap.L().Sync() //日志来不及实时写入
	// 加载配置信息
	if err:=Controllers.Initconfig();err != nil{
		zap.L().Error("Initconfig failed",zap.Error(err))
	}
	zap.L().Info("Initconfig Success")
	// 初始化MySQL
	if err:=MySQL.InitMySQL();err !=nil{
		zap.L().Error("InitMySQL failed",zap.Error(err))
		return
	}
	zap.L().Info("InitMySQL Success")
	defer MySQL.Close()
	r:=Router.SetRouter()
	r.Run(":8080")
}

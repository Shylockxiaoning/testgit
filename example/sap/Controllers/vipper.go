package Controllers

import (
"github.com/spf13/viper"
)

//Initconfig() 从配置文件中加载信息
func Initconfig()error{
	viper.SetConfigName("config") 	//指定配置文件名称（不需要后缀）
	viper.SetConfigType("yaml")		//指定配置文件类型
	viper.AddConfigPath(".")			//指定配置文件路径（使用相对路径）
	err := viper.ReadInConfig()
	viper.WatchConfig()  				//监控配置文件变更
	return err
}

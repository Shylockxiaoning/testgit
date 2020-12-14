package MySQL

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	//导入sqlx需要导入驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//InitMySQL() 初始化数据库
func InitMySQL()(err error){
	//链接mysql
	dsn :=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db"),
		)
	db,err = sqlx.Connect("mysql",dsn)
	if err != nil{
		zap.L().Error("Connect MySQL failed",zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.MaxOpenCons"))
	db.SetMaxIdleConns(viper.GetInt("mysql.MaxIdleCons"))
	return
}

//对外暴露关闭方法
func Close(){
	db.Close()
}

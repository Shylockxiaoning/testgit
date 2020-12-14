package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

//Initlogger 初始化zap日志
func Initlog(){
	writesyncer := getLogWrite()  //日志输出位置相关
	encoder := getEncoder()			//日志的格式相关
	core := zapcore.NewCore(encoder,writesyncer,zapcore.DebugLevel)
	logger := zap.New(core,zap.AddCaller())		//根据上面的配置创建logger   zap.AddCaller() // 增加行
	zap.ReplaceGlobals(logger)  //替换zap库里全局的
}
//getEncoder() 日志的格式相关
func getEncoder() zapcore.Encoder{
	encoderConfig :=zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// zapcore.NewTee()  多个日志输出
	return zapcore.NewConsoleEncoder(encoderConfig)  //可读日志
	// zapcore.NewJSONEncoder(encoderConfig) //json格式日志
}

//可以利用第三方库来进行日志库的分割，替换该位置即可
//getLogWrite() 日志输出位置相关
func getLogWrite() zapcore.WriteSyncer{
	file,_ := os.Create("./logger/logger.log")
	return zapcore.AddSync(file)
}

package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code MyCode `json:"code"`		  //业务状态码
	Message string `json:"message"`   //响应信息
	Data interface{} `json:"data"`    //自定义数据
}

//ResponseSuccess 正常响应信息
func ResponseSuccess(ctm *gin.Context,data interface{}){
	re := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctm.JSON(http.StatusOK, re)
}

//ResponseError 错误响应状态
func ResponseError(ctm *gin.Context,c MyCode){
	re :=&ResponseData{
		Code: c,
		Message: c.Msg(),
		Data: nil,
	}
	ctm.JSON(http.StatusOK,re)
}

//ResponseErrorWithMsg 返回错误信息
func ResponseErrorWithMsg(ctm *gin.Context,code MyCode,errMsg string){
	re :=&ResponseData{
		Code: code,
		Message: errMsg,
		Data: nil,
	}
	ctm.JSON(http.StatusOK,re)
}
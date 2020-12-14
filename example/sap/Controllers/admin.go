package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sap/Dao/MySQL"
	"sap/Model"
	"strconv"
)

//登陆成功的首页
//func IndexHandler(c *gin.Context){
//	var godown Model.Godown
//	c.HTML(http.StatusOK,"index.html", gin.H{
//		"res":godown,
//	})
//}

//AddHandler 增加数据
func AddHandler(c *gin.Context){
	var godown Model.Godown
	price,_:=strconv.Atoi(c.PostForm("price"))
	inventory,_:=strconv.Atoi(c.PostForm("inventory"))
	{
		godown.Title=c.PostForm("title")
		godown.IssueDate=c.PostForm("issue_date")
		godown.Synopsis=c.PostForm("synopsis")
		godown.Category=c.PostForm("category")
		godown.Author=c.PostForm("author")
		godown.Price=price
		godown.Inventory=inventory
	}
	if err:=MySQL.InsertIntoMySQL(&godown);err!=nil{
		zap.L().Error("增加失败",zap.Error(err))
		ResponseErrorWithMsg(c,CodeServerBusy,"增加失败")
		return
	}
	c.Redirect(http.StatusMovedPermanently,"/admin/post")
}
// DeleteHandler 删除数据
func DeleteHandler(c *gin.Context)  {
	c.HTML(http.StatusOK,"delete.html","")
	title:=c.Query("title")
	if err:=MySQL.DeleteMySQL(title);err!=nil{
		ResponseErrorWithMsg(c,CodeServerBusy,fmt.Sprintf("删除%v失败,请确认库存是否为0",title))
		return
	}
	c.String(http.StatusOK,"删除成功")
}


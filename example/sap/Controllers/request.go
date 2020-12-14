package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sap/Dao/MySQL"
	"sap/Logic"
)

func Get(c *gin.Context){
	//page:=c.Request.URL.Query().Get("page")
	////number:=Logic.Getpage(page)
	//fmt.Println(page)
	msgs,count:=MySQL.SelectMySQL()
	list,_:=Logic.Page(count)
	c.HTML(http.StatusOK,"index.html",gin.H{
		"res":msgs,
		"list":list,
	})
}

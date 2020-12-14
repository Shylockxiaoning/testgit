package Router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sap/Controllers"
)

// SetRouter() 配置路由信息
func SetRouter() * gin.Engine{
	router:=gin.Default()
	router.LoadHTMLGlob("./template/*")
	sap:=router.Group("/sap")
	{
		sap.POST("/login",Controllers.LoginHandler)
		sap.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK,"login.html","")
		})
		//登录时，没有账号引导进入注册页面
		sap.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK,"register.html","")
		})
		sap.POST("/register",Controllers.Register)
		//登录成功后才能访问的页面
		admin := router.Group("/admin")
		{
			//验证中间件
			admin.Use(Controllers.CheckAuth)
			admin.GET("/index",Controllers.Get) 		//查

			admin.GET("/post", func(c *gin.Context) {
				c.HTML(http.StatusOK,"add.html","")
			})
			admin.POST("/post",Controllers.AddHandler)		//增
			admin.Use(Controllers.Admin)
			admin.GET("/delete",Controllers.DeleteHandler)	//删
			admin.PUT("/put") 		//改
		}
	}
	return router
}

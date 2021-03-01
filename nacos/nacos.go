package nacos

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nacos-web/auth"
	"nacos-web/config"
	"os"
)

const (
	pathSeparator = string(os.PathSeparator)
	staticPath    = "static"
)

var indexHandler = func(c *gin.Context) {
	c.File("static" + pathSeparator + "index.html")
}

func Start() {
	CreateProxy(config.AppConfig.Nacos.Endpoint)
	c := gin.Default()
	c.Any("/", func(c *gin.Context) {
		c.Redirect(302, "/nacos/")
	})
	c.Any("/nacos/", indexHandler)
	c.Any("/nacos/js/*fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		c.File("static" + pathSeparator + "js" + pathSeparator + fileName)
	})
	c.Any("/nacos/css/*fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		c.File("static" + pathSeparator + "css" + pathSeparator + fileName)
	})
	c.Any("/nacos/img/*fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		c.File("static" + pathSeparator + "img" + pathSeparator + fileName)
	})
	c.Any("/nacos/console-ui/*fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")
		c.File("static" + pathSeparator + "console-ui" + pathSeparator + fileName)
	})
	c.Any("/nacos/v1/*path", func(c *gin.Context) {
		accessToken := c.Query("accessToken")
		path := c.Param("path")
		fmt.Println(path)
		if path != "/auth/users/login" {
			ok, subject, err := auth.Verify(accessToken)
			if err != nil {
				panic(err)
			}
			if !ok || subject != "nacos" {
				c.JSON(403, gin.H{
					"message": "用户登录状态已失效",
				})
			}
			c.Abort()
			return
		}
		proxy := balanceProxy()
		proxy.ServeHTTP(c.Writer, c.Request)
	})
	c.Run(config.AppConfig.Addr)
}

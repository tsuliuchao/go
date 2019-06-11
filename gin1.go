package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"fmt"
	"os"
	"io"
)

func main() {
	r := gin.Default()
	//创建日志文件，记录文件同时，输出到控制台
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
//http://10.115.1.18:8080/user/john/send
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
//默认查询参数 http://10.115.1.18:8080/welcome?firstname=licuhao
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
//http://10.115.1.18:8080/form_post
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
//query + post form
// curl -v -d"name=chao" "http://10.115.1.18:8080/post?id=1234&page=1" 
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		//fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.JSON(200,gin.H{
			"id":  id,
			"page": page,
			"name":    name,
			"message": message,
		})
	})
//分组情况
//	v1 := r.Group("/v1")
//	{
//		v1.POST("/login", loginEndpoint)
//		v1.POST("/submit", submitEndpoint)
//		v1.POST("/read", readEndpoint)
//	}

	r.Run() // listen and serve on 0.0.0.0:8080
}

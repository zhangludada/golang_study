package main

import (
	"path"

	"github.com/gin-gonic/gin"
)

// go build then ./main
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件，
		f, _ := c.FormFile("f1")

		// 将读取到的文件保存在本地(服务端)
		//dst := fmt.Sprintf("./%s", f.Filename)
		dst := path.Join("./", f.Filename)
		c.SaveUploadedFile(f, dst)

		c.JSON(200, gin.H{"status": "ok"})
	})
	r.Run(":8080")
}

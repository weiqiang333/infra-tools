package fileserver

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"infra-tools-fileserver/model"
)

// FileDownload 允许下载目录管理
// http.ServeFile 实现断点续传
func FileDownload(c *gin.Context)  {
	filePath := "/" + c.Param("filepath") + "/"
	fileName := c.Param("filename")
	for _, dir := range model.Dir {
		if filePath == dir {
			c.File(filePath + fileName)
			return
		}
	}
	c.Redirect(301, "/download/")
	return
}

// Index 入口文件
func Index(c *gin.Context)  {
	c.HTML(http.StatusOK, "download.tmpl", gin.H{
		"dir": model.Dir,
	})
}
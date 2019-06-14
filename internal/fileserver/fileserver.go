package fileserver

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"infra-tools/model"
)

// FileDownload 允许下载目录管理
// http.ServeFile 实现断点续传
func FileDownload(c *gin.Context)  {
	urlpath := c.Request.URL.Path
	filepath := strings.Split(urlpath, "/download")[1]
	c.File(filepath)
	return
}

// Index 入口文件
func Index(c *gin.Context)  {
	c.HTML(http.StatusOK, "download.tmpl", gin.H{
		"dir": model.Dir,
	})
}
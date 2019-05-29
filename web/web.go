package web

import (
	"github.com/gin-gonic/gin"

	"infra-tools/internal/fileserver"
)

// Web 监控路由入口
func Web()  {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/**/*")
	router.GET("/stat", func(c *gin.Context) {
		c.String(200, "ok")
	})
	down := router.Group("/download")
	{
		down.GET(":filepath/*filename", fileserver.FileDownload)
	}
	router.NoRoute(fileserver.Index)
	router.Run()
}
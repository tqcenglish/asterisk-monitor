package main

import (
	"asterisk-monitor/web"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
)

func main() {
	root, _ := fs.Sub(web.WWWFiles, "www")
	ginS.StaticFS("/www", http.FS(root))
	ginS.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/www")
	})

	ginS.GET("/hello", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}

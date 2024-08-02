package api

import (
	"asterisk-monitor/web"
	"io/fs"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
)

func HttpServer() {
	root, _ := fs.Sub(web.WWWFiles, "www")
	ginS.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	ginS.StaticFS("/www", http.FS(root))
	ginS.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/www")
	})

	ginS.GET("/hello", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.GET("/api/systeminfo", SystemInfo)
	ginS.GET("/api/storageinfo", StorageInfo)
	ginS.GET("/api/queuestatus", QueueStatus)
	ginS.GET("/api/calllog", CallLog)
	ginS.GET("/api/extensionstatus", Extensionstatus)
	ginS.GET("/api/trunkstatus", Trunkstatus)
	ginS.Run()
}

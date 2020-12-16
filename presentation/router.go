package presentation

import (
	"math/rand"
	"net/http"

	"github.com/KendoCross/kendoDDD/presentation/api"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//表现层主要的职责在于，表现形式的多样化。

func InitRouter() *gin.Engine {

	if viper.GetString("APP_MODE") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	r := gin.Default()
	// 自定义error handler
	//r.Use(middleware.Errors())
	//r.Use(md.GinCors())

	// 请求记录写入日志文件
	//r.Use(logger.RequestHandler(logger.Logger), gin.Recovery())
	r.GET("/", healthHand)

	files := r.Group("files")
	r.GET("/:id", api.GetFile)
	files.POST("", api.AddFile)

	return r
}

var pongStr = []string{"Hello World", "你好，世界", "こんにちは世界", "Hallo Welt", "Привет, мир", "Bonjour le monde", "Hei maailma", "Saluton mondo", "salve Orbis Terrarum", "Сайн уу", "Nyob zoo lub ntiaj teb"}

//健康检查API
func healthHand(c *gin.Context) {
	index := rand.Intn(10)
	c.JSON(http.StatusOK, pongStr[index])
}

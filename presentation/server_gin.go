package presentation

import (
    "github.com/gin-gonic/gin"
    "github.com/hollson/goddd/proxy/handler"
    "github.com/spf13/viper"
)

// 表现层主要的职责在于，表现形式的多样化。
func NewGinSerer() *gin.Engine {
    if viper.GetString("APP_MODE") == "prod" {
        gin.SetMode(gin.ReleaseMode)
        gin.DisableConsoleColor()
    }

    router := gin.Default()
    // 自定义error handler
    // r.Use(middleware.Errors())
    // r.Use(md.GinCors())
    f := router.Group("/file")
    f.GET("/:id", handler.GetFileHandler)
    f.POST("/upload", handler.AddFileHandler)

    return router
}

// var pongStr = []string{"Hello World", "你好，世界", "こんにちは世界", "Hallo Welt", "Привет, мир", "Bonjour le monde", "Hei maailma", "Saluton mondo", "salve Orbis Terrarum", "Сайн уу", "Nyob zoo lub ntiaj teb"}
//
// // 健康检查API
// func healthHand(c *gin.Context) {
//     index := rand.Intn(10)
//     c.JSON(http.StatusOK, pongStr[index])
// }

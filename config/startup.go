package config

import (
    "fmt"

    mysql "github.com/hollson/goddd/infrastructure/repos_mysql"
    "github.com/hollson/goddd/domain"
    "github.com/hollson/goddd/infrastructure/logs"

    "github.com/spf13/viper"
)

func StartUp() {

    viper.SetConfigFile("config/config.yaml")
    viper.SetConfigType("yaml")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Sprintf("Error while loading config file [conf/config.yaml]: %s", err.Error()))
    }

    // 日志组件
    logs.InitLogs(viper.GetString("LOG_FILE"))

    // 仓储层DB引擎
    mysql.InitDB()
    // redis.InitRedis()

    // 最后启动领域层
    domain.StartUp()
}

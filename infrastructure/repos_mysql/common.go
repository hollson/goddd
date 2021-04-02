package repos_mysql

import (
    "fmt"

    "github.com/hollson/goddd/interfaces"
    "github.com/spf13/viper"

    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

var readEngine *xorm.Engine
var writeEngine *xorm.Engine

// InitDB 初始化DB引擎
func InitDB() {

    source := viper.GetString("MYSQL_DSN")
    var err error
    readEngine, err = xorm.NewEngine("mysql", source)
    if err != nil {
        panic(fmt.Errorf("mysql error: %v", err))
    }

    readEngine.SetMaxIdleConns(viper.GetInt("MYSQL_MAXIDLE"))
    readEngine.SetMaxOpenConns(viper.GetInt("MYSQL_MAXOPEN"))

    // 读写需要分开不同的库
    writeEngine, err = readEngine.Clone()
    if err != nil {
        panic(fmt.Errorf("mysql error: %v", err))
    }
    writeEngine.SetMaxIdleConns(viper.GetInt("MYSQL_MAXIDLE"))
    writeEngine.SetMaxOpenConns(viper.GetInt("MYSQL_MAXOPEN"))

    if viper.GetBool("MYSQL_SHOWSQL") {
        readEngine.ShowSQL(true)
        writeEngine.ShowSQL(true)
    }

    writeEngine.Sync2(new(interfaces.FileInfo))
}

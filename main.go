package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/ckeyer/beewechat/models"
	"github.com/hiprice/beewechat/conf"
	_ "github.com/hiprice/beewechat/routers"
	"github.com/hiprice/beewechat/wechat"
)

func init() {
	config := conf.NewConfig()
	force := false   // 强制重新建表
	verbose := false // 打印SQL语句
	orm.Debug = true

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", config.WX_MYSQL_CONN)
	wechat.RegDB()
	orm.RunSyncdb("default", force, verbose)
}

func main() {
	beego.Run()
}

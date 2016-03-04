package wechat

import (
	"github.com/astaxie/beego/orm"
	"github.com/hiprice/beewechat/conf"
	"github.com/hiprice/beewechat/wechat/event"
	"github.com/hiprice/beewechat/wechat/msg"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/redis"
)

var (
	redcli redis.Client
	config *conf.CkConfig
)

func init() {
	config = conf.NewConfig()
	redcli.Addr = config.REDIS_ADDR
}

func RegDB() {
	msg.RegDB()
	event.RegDB()
	orm.RegisterModel(new(WebUserInfo))
}

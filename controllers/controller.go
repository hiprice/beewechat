package controllers

import (
	"github.com/hiprice/beewechat/conf"
)

var (
	config *conf.CkConfig
)

func init() {
	config = conf.NewConfig()
}

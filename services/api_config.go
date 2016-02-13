package services

import conf "github.com/andboson/configlog"

var ApiName, _ = conf.AppConfig.String("api_name")
var Debug, _ = conf.AppConfig.Bool("debug")

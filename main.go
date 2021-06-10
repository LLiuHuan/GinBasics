package main

import (
	"fairman-server/global"
	"fairman-server/initialize"
)

func init() {
	global.FMS_VP = initialize.Viper()
	global.FMS_LOG = initialize.Zap()
}

func main() {
	initialize.RunServer()
}
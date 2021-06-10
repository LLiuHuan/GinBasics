package main

import (
	"gin-basics/global"
	"gin-basics/initialize"
)

func init() {
	global.GB_VP = initialize.Viper()
	global.GB_LOG = initialize.Zap()
}

func main() {
	initialize.RunServer()
}

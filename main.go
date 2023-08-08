package main

import (
	"github.com/nico612/newbee-goo/core"
	"github.com/nico612/newbee-goo/global"
	"github.com/nico612/newbee-goo/initialize"
)

func main() {
	global.Viper = core.Viper()
	global.Logger = core.Zap()
	global.DB = initialize.Gorm()

	core.RunWindowsServer()
}

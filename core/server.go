package core

import (
	"fmt"
	"github.com/nico612/newbee-goo/global"
	"github.com/nico612/newbee-goo/initialize"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.Config.System.Addr)

	s := initServer(address, router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.Logger.Info("server run success on ", zap.String("address", address))

	err := s.ListenAndServe()
	if err != nil {
		global.Logger.Error(err.Error())
	}
}

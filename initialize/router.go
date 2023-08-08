package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/nico612/newbee-goo/global"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.StaticFS(global.Config.Local.Path, http.Dir(global.Config.Local.Path)) // 为用户头像和文件提供静态地址
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.Logger.Info("use middleware logger")

	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	//global.Logger.Info("use middleware cors")

	// 方便统一添加路由组前缀 多服务器上线使用
	//商城后管路由

	return r

}

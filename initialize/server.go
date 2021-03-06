package initialize

import (
	"fmt"
	"gin-basics/core"
	"gin-basics/global"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

func init() {
	// 初始化Validator翻译器
	if err := InitTrans("zh"); err != nil {
		fmt.Printf("init validator failed, err: %v\n", err)
		global.GB_LOG.Error("init validator failed, err", zap.Error(err))
		os.Exit(0)
		return
	}

	// 初始化Docker API
	InitContent("local")
}

func RunServer() *http.Server {
	// 注册路由
	router := Routers("")

	address := fmt.Sprintf(":%d", global.GB_CONFIG.System.Port)
	// 运行
	s := core.InitServer(address, router)
	time.Sleep(10 * time.Microsecond)

	global.GB_LOG.Info("服务启动成功，监听端口为：  ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 %s
	当前版本: %s
	默认自动化文档地址: http://127.0.0.1%s/swagger/index.html
	`, global.GB_CONFIG.System.Name, global.GB_CONFIG.System.Version, address)
	fmt.Println()

	//global.GB_LOG.Error(s.ListenAndServe().Error())
	//fmt.Println("33333333333333333")
	Close(s)
	return s
}

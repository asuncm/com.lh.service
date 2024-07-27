package main

import (
	"com.lh.basic/config"
	"com.lh.service/locales"
	"com.lh.service/tools"
	"com.lh.service/yugabyte"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	config.InitConfig("com.lh.service")
	configs := config.GetConfig("service")
	yugabyte.InitConfig()
	app.Use(tools.Cors())
	//app.Use(tools.MiddleWare(configs))
	locales.Init()
	//router.Router(app)
	app.Run(fmt.Sprintf("%s:%s", configs.Host, configs.Port))
}

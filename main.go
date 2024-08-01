package main

import (
	auth "com.lh.auth/locales"
	"com.lh.basic/config"
	basic "com.lh.basic/locales"
	serve "com.lh.service/locales"
	"com.lh.service/tools"
	"com.lh.service/yugabyte"
	user "com.lh.user/locales"
	web "com.lh.web.service/locales"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	config.InitConfig("com.lh.service")
	configs := config.GetServe("service")
	yugabyte.InitConfig()
	app.Use(tools.Cors())
	root := config.GetKey("Root").(string)
	//app.Use(tools.MiddleWare(configs))
	auth.Init(root)
	basic.Init(root)
	user.Init(root)
	web.Init(root)
	serve.Init(root)
	//router.Router(app)
	app.Run(fmt.Sprintf("%s:%s", configs.Host, configs.Port))
}

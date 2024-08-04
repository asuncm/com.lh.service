package main

import (
	"com.lh.basic/config"
	"com.lh.service/tools"
	"com.lh.service/yugabyte"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	opts := config.InitConfig("com.lh.service")
	yugabyte.InitConfig()
	app.Use(tools.Cors())
	//app.Use(tools.MiddleWare(configs))
	//router.Router(app)
	app.Run(fmt.Sprintf("%s:%s", opts["host"], opts["ports"]))
}

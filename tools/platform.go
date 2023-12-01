package tools

import (
	"github.com/gin-gonic/gin"
	"os"
	"runtime"
	"strings"
)

type PlatformConfig struct {
	Platform string `json:"platform"` // 平台类型
	Env      string `json:"env"`      // 当前环境
	Dir      string `json:"dir"`      // 获取目录
}

func Platform(dir string) PlatformConfig {
	platform := runtime.GOOS
	env := os.Getenv("mode")

	return PlatformConfig{
		Platform: platform,
		Env:      env,
		Dir:      dir,
	}
}

func Pathname(c *gin.Context, ext string) string {
	dataCache, _ := c.Get("DataCache")
	serve, _ := c.Get("Serve")
	paths := []string{any(dataCache).(string), any(serve).(string)}
	return strings.Join(paths, ext)
}

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
	if IsStringEmtpy(env) {
		env = "local"
	}
	return PlatformConfig{
		Platform: platform,
		Env:      env,
		Dir:      dir,
	}
}

type PathConfig struct {
	Root string `json:"root"`
	Path string `json:"path"`
	Item string `json:"item"`
}

func Pathname(c *gin.Context, ext string) PathConfig {
	dataCache, _ := c.Get("DataCache")
	serve, _ := c.Get("Serve")
	path := os.Getenv(any(dataCache).(string))
	paths := []string{path, any(serve).(string)}
	ps := strings.Join(paths, ext)
	ps = strings.Replace(ps, "\\", "/", -1)
	return PathConfig{
		Root: path,
		Path: ps,
		Item: any(serve).(string),
	}
}

func GetPath(key string, suffix string) string {
	path := os.Getenv(any(key).(string))
	path = strings.Replace(path, "\\", "/", -1)
	paths := []string{path, suffix}
	ps := strings.Join(paths, "/")
	return ps
}

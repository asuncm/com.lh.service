package tools

import (
	"os"
	"runtime"
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

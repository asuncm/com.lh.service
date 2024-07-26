package tools

import (
	"os"
	"regexp"
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

func GetPath(key string, suffix string) string {
	path := os.Getenv(any(key).(string))
	path = strings.Replace(path, "\\", "/", -1)
	paths := []string{path, suffix}
	ps := strings.Join(paths, "/")
	isB, err := regexp.MatchString("/$", ps)
	if isB && err == nil {
		return ps[:len(ps)-1]
	}
	return ps
}

package proxy

import (
	"com.lh.service/src/config/yaml"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
)

func Proxy(c *gin.Context, config yaml.ServeConf) {
	re := regexp.MustCompile("^/api/")
	URL := c.Request.URL
	uri := re.ReplaceAllString(URL.Path, "")
	hosts := []string{"http://", config.Host}
	domain := []string{strings.Join(hosts, ""), config.Port}
	target := strings.Join(domain, ":")
	proxyUrl, _ := url.Parse(target)
	c.Request.URL.Path = uri
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	proxy.ServeHTTP(c.Writer, c.Request)
}

package tools

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
)

func Proxy(c *gin.Context, config MapConf) {
	re := regexp.MustCompile("^/api/")
	URL := c.Request.URL
	uri := re.ReplaceAllString(URL.Path, "")
	hosts := []string{"http://", config["host"].(string)}
	domain := []string{strings.Join(hosts, ""), config["port"].(string)}
	target := strings.Join(domain, ":")
	proxyUrl, _ := url.Parse(target)
	c.Request.URL.Path = uri
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	proxy.ServeHTTP(c.Writer, c.Request)
}

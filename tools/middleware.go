package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, SSID, Verify, UUID, Language")
		c.Header("Access-Control-Allow-Methods", method)
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func MiddleWare(options ServeConf) gin.HandlerFunc {
	return func(c *gin.Context) {
		list := reflect.ValueOf(&options)
		elem := list.Elem()
		refType := elem.Type()
		for i := 0; i < refType.NumField(); i++ {
			name := refType.Field(i).Name
			c.Set(name, elem.Field(i).Interface())
		}
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

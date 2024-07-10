package tools

import (
	"com.lh.auth/locales"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Code500(msg string, c *gin.Context) {
	defer func() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"msg":     msg,
			"success": false,
			"data":    nil,
		})
		return
	}()
	c.Next()
	return
}

func Code200(data interface{}, c *gin.Context) {
	arrs := []string{"success", "msg"}
	msg := locales.GetKey(c, arrs)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"msg":     msg,
		"success": true,
		"data":    data,
	})
}

func Code404(msg error, c *gin.Context) {
	defer func() {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"msg":     msg,
			"success": false,
			"data":    nil,
		})
		return
	}()
	c.Next()
	return
}

func Code400(msg error, c *gin.Context) {
	defer func() {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"msg":     msg,
			"success": false,
			"data":    nil,
		})
		return
	}()
	c.Next()
	return
}

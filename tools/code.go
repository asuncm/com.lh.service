package tools

import (
	service "com.lh.service/locales"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Code500(msg string, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"msg":     msg,
		"success": false,
	})
	return
}

func Code200(data interface{}, c *gin.Context) {
	arrs := []string{"success", "msg"}
	msg := service.GetKey(c, arrs)
	fmt.Println(c, "poooooo")
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

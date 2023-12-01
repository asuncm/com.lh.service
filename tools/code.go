package tools

import "github.com/gin-gonic/gin"

type codeConf = map[string]any

func Code500(msg error) gin.H {
	return codeConf{
		"code":    500,
		"msg":     msg,
		"success": false,
		"data":    nil,
	}
}

func Code200(msg error, data interface{}) gin.H {
	return gin.H{
		"code":    200,
		"msg":     msg,
		"success": true,
		"data":    data,
	}
}

func Code404(msg error) gin.H {
	return gin.H{
		"code":    400,
		"msg":     msg,
		"success": false,
		"data":    nil,
	}
}

func Code400(msg error) gin.H {
	return gin.H{
		"code":    400,
		"msg":     msg,
		"success": false,
		"data":    nil,
	}
}

func Code(code int, msg error) gin.H {
	return gin.H{
		"code":    code,
		"msg":     msg,
		"success": false,
		"data":    nil,
	}
}

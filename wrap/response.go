package wrap

import (
	"github.com/gin-gonic/gin"

	"github.com/zhenzhongfu/gocommon/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode int, errCode e.StandardError, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode.ErrCode,
		"msg":  errCode.ErrMsg,
		"data": data,
	})

	return
}

package starUtils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 0 // 操作成功
	FAILED  int = 1
)

func GinSuccess(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "成功",
		"data": v,
	})
}

func GinFailed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  v,
	})
}

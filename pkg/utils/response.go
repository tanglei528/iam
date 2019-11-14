package utils

import (
    "github.com/gin-gonic/gin"
    e "iam/pkg/exception"
)

type Gin struct {
    C *gin.Context
}

func (g *Gin) Response(httpCode, errorCode int, data interface{}, errors interface{}) {
    g.C.JSON(httpCode, gin.H{
        "code": errorCode,
        "msg": e.GetMsg(errorCode),
        "data": data,
        "errors": errors,
    })

    return
}

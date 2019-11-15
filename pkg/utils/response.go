package utils

import (
    "github.com/gin-gonic/gin"
    e "iam/pkg/exception"
    "reflect"
)

type Gin struct {
    C *gin.Context
}

func (g *Gin) Response(httpCode, errorCode int, data interface{}, errors interface{}) {
    vi := reflect.ValueOf(errors)
    if vi.Kind() == reflect.Slice && vi.IsNil() {
        errors = e.GetMsg(errorCode)
    }
    g.C.JSON(httpCode, gin.H{
        "code": httpCode,
        "msg": e.GetMsg(errorCode),
        "data": data,
        "errors": errors,
    })

    return
}

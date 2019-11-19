package v1

import (
    "github.com/gin-gonic/gin"
    "iam/app/services/auth_service"
    "iam/pkg/utils"
)

func Login(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := auth_service.Login(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

// 外部平台token校验
func ValidateToken(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := auth_service.ValidateToken(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

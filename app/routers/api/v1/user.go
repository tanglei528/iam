package v1

import (
    "github.com/gin-gonic/gin"
    "iam/app/services/user_service"
    "iam/pkg/utils"
)

// 创建用户
func CreateUser(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := user_service.CreateUser(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func DeleteUser(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := user_service.DeleteUser(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func UpdateUser(c *gin.Context) {
   appG := utils.Gin{C: c}
   httpCode, errCode, data, errorsMsg := user_service.UpdateUser(&appG)
   appG.Response(httpCode, errCode, data, errorsMsg)
}

func ListUsers(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := user_service.ListUsers(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func GetUserByID(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := user_service.GetUserByID(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

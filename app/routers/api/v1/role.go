package v1

import (
    "github.com/gin-gonic/gin"
    "iam/app/services/role_service"
    "iam/pkg/utils"
)

func CreateAppRole(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := role_service.CreateAppRole(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func DeleteAppRole(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := role_service.DeleteAppRole(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func UpdateAppRole(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := role_service.UpdateAppRole(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func ListAppRoles(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := role_service.ListAppRoles(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func GetAppRoleByID(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := role_service.GetAppRoleByID(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

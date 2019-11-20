package v1

import (
    "github.com/gin-gonic/gin"
    "iam/app/services/action_service"
    "iam/pkg/utils"
)

func CreateAppAction(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := action_service.CreateAppAction(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func DeleteAppAction(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := action_service.DeleteAppAction(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func UpdateAppAction(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := action_service.UpdateAppAction(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func ListAppActions(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := action_service.ListAppActions(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

func GetAppActionByID(c *gin.Context) {
    appG := utils.Gin{C: c}
    httpCode, errCode, data, errorsMsg := action_service.GetAppActionByID(&appG)
    appG.Response(httpCode, errCode, data, errorsMsg)
}

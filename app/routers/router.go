package routers

import (
    "github.com/gin-gonic/gin"
    "iam/app/routers/api/v1"
    "iam/pkg/utils"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    apiv1 := r.Group("/api/v1")
    apiv1.POST("/login", v1.Login)
    apiv1.POST("/validate", v1.ValidateToken)

    apiv1.Use(utils.JWT())

    {
        apiv1.GET("/apps", v1.GetApps)
        apiv1.GET("/apps/:id", v1.GetAppByID)
        apiv1.POST("/apps", v1.CreateApp)
        apiv1.PUT("/apps/:id", v1.UpdateApp)
        apiv1.DELETE("/apps/:id", v1.DeleteApp)

        apiv1.POST("/apps/:id/roles", v1.CreateAppRole)
        apiv1.GET("/apps/:id/roles", v1.ListAppRoles)
        apiv1.GET("/apps/:id/roles/:role_id", v1.GetAppRoleByID)
        apiv1.PUT("/apps/:id/roles/:role_id", v1.UpdateAppRole)
        apiv1.DELETE("/apps/:id/roles/:role_id", v1.DeleteAppRole)

        apiv1.POST("/users", v1.CreateUser)
        apiv1.DELETE("/users/:id", v1.DeleteUser)
        apiv1.PUT("/users/:id", v1.UpdateUser)
        apiv1.GET("/users", v1.ListUsers)
        apiv1.GET("/users/:id", v1.GetUserByID)

    }

    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "test",
        })
    })

    return r
}

package v1

import (
    "github.com/Unknwon/com"
    "github.com/gin-gonic/gin"
    "iam/app/services"
    e "iam/pkg/exception"
    "iam/pkg/utils"
    "net/http"
)

// 可以把逻辑抽取到service层，api层只保留基础的调用，参照role

type AddAppForm struct {
    Name          string    `form:"name" json:"name" valid:"Required;MaxSize(100)"`
    IsActive      int       `form:"is_active" json:"is_active" valid:"Required;Range(0,1)"`
    Description   string    `form:"description" json:"description" valid:"Required;MaxSize(255)"`
    IndexUrl      string    `form:"index_url" json:"index_url" valid:"Required;MaxSize(255)"`
    LoginUrl      string    `form:"login_url" json:"login_url" valid:"Required;MaxSize(255)"`
}

func CreateApp(c *gin.Context) {
    var (
        appG = utils.Gin{C: c}
        form AddAppForm
    )

    httpCode, errCode, errorsMsg := utils.BindAndValid(c, &form)
    if errCode != e.Success {
        appG.Response(httpCode, errCode, nil, errorsMsg)
        return
    }

    isActive, err := utils.IntToBool(form.IsActive)
    if err != nil {
        appG.Response(http.StatusInternalServerError, e.ErrorConvIntToBool, nil, nil)
        return
    }
    appInfo := services.AppInfo{
        Name:           form.Name,
        IsActive:       isActive,
        Description:    form.Description,
        IndexUrl:       form.IndexUrl,
        LoginUrl:       form.LoginUrl,
    }

    data, err := appInfo.CreateApp();
    if  err != nil {
        appG.Response(http.StatusInternalServerError, e.ErrorCreateAppFail, nil, nil)
        return
    }
    appG.Response(http.StatusOK, e.SuccessCreate, data, nil)
}

func DeleteApp(c *gin.Context) {
    appG := utils.Gin{C: c}
    id := com.StrTo(c.Param("id")).MustInt()
    appInfo := services.AppInfo{
        ID: id,
    }
    if err := appInfo.DeleteApp(); err != nil {
        appG.Response(http.StatusInternalServerError, e.ErrorGetAppFail, nil, nil)
        return
    }

    data := make(map[string]interface{})
    data["items"] = nil
    appG.Response(http.StatusOK, e.SuccessResponse, data, nil)
}

type UpdateAppForm struct {
    Name          string    `form:"name" json:"name" valid:"MaxSize(100)"`
    IsActive      int       `form:"is_active" json:"is_active" valid:"Range(0,1)"`
    Description   string    `form:"description" json:"description" valid:"MaxSize(255)"`
    IndexUrl      string    `form:"index_url" json:"index_url" valid:"MaxSize(255)"`
    LoginUrl      string    `form:"login_url" json:"login_url" valid:"MaxSize(255)"`
}

func UpdateApp(c *gin.Context) {
    var (
        appG = utils.Gin{C: c}
        form UpdateAppForm
    )

    id := com.StrTo(c.Param("id")).MustInt()

    httpCode, errCode, errorsMsg := utils.BindAndValid(c, &form)
    if errCode != e.Success {
        appG.Response(httpCode, errCode, nil, errorsMsg)
        return
    }

    isActive, err := utils.IntToBool(form.IsActive)
    if err != nil {
        appG.Response(http.StatusInternalServerError, e.ErrorConvIntToBool, nil, nil)
        return
    }
    appInfo := services.AppInfo{
        ID:             id,
        Name:           form.Name,
        IsActive:       isActive,
        Description:    form.Description,
        IndexUrl:       form.IndexUrl,
        LoginUrl:       form.LoginUrl,
    }

    data, err := appInfo.UpdateApp();
    if  err != nil {
        appG.Response(http.StatusInternalServerError, e.ErrorCreateAppFail, nil, nil)
        return
    }
    appG.Response(http.StatusOK, e.Success, data, nil)
}

func GetApps(c *gin.Context) {
    appG := utils.Gin{C: c}

    //appView := services.AppView{
    //    PageNum: pageNum,
    //    PageSize: pageSize,
    //}
    appInfo := services.AppInfo{}
    total, err := appInfo.Count()
    if err != nil  {
        appG.Response(http.StatusInternalServerError, e.ErrorCountAppFail, nil, nil)
        return
    }

    pageNum := utils.GetPage(c, total)
    pageSize := utils.GetPageSize(c)

    apps, err := appInfo.GetAll(pageNum, pageSize)
    if err != nil {
        appG.Response(http.StatusInternalServerError, e.ErrorGetAppsFail, nil, nil)
        return
    }

    data := make(map[string]interface{})
    data["items"] = apps
    data["total"] = total
    data["page_number"] = utils.ShowPageNum(pageNum, pageSize, total)
    data["page_size"] = pageSize
    appG.Response(http.StatusOK, e.Success, data, nil)
}

func GetAppByID(c *gin.Context) {
    appG := utils.Gin{C: c}
    id := com.StrTo(c.Param("id")).MustInt()
    appInfo := services.AppInfo{
        ID: id,
    }
    apps, err := appInfo.GetByID()
    if err != nil {
        appG.Response(http.StatusInternalServerError, e.ErrorGetAppFail, nil, nil)
        return
    }
    data := make(map[string]interface{})
    data["items"] = apps
    appG.Response(http.StatusOK, e.Success, data, nil)
}

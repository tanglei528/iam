package services

import (
    "iam/app/models"
    "reflect"
    "time"
)

type AppInfo struct {
    ID             int          `json:"id"`
    Name           string       `json:"name"`
    IsActive       bool         `json:"is_active"`
    Description    string       `json:"description"`
    IndexUrl       string       `json:"index_url"`
    LoginUrl       string       `json:"login_url"`
    CreatedAt      time.Time    `json:"created_at"`
    UpdatedAt      time.Time    `json:"updated_at"`
}

//type AppView struct {
//    //ID             int
//    //Name           string
//    //IsActive       bool
//    //Description    string
//    //IndexUrl       string
//    //LoginUrl       string
//    //CreatedAt      time.Time
//    //UpdatedAt      time.Time
//
//    PageNum        int
//    PageSize       int
//}

func (appInfo *AppInfo) CreateApp() (*AppInfo, error) {
    app := map[string]interface{}{
        "name":         appInfo.Name,
        "is_active":    appInfo.IsActive,
        "description":  appInfo.Description,
        "index_url":    appInfo.IndexUrl,
        "login_url":    appInfo.LoginUrl,
    }
    dbApp, err := models.AddApp(app);
    if err != nil {

        return nil, err
    }
    appInfo.ID = dbApp.ID
    appInfo.CreatedAt = dbApp.CreatedAt
    appInfo.UpdatedAt = dbApp.UpdatedAt
    return appInfo, nil
}

func (appInfo *AppInfo) DeleteApp() error {
    err := models.DeleteApp(appInfo.ID)
    return err
}

func (appInfo *AppInfo) UpdateApp() ([]AppInfo, error) {
    app := map[string]interface{}{
        "name":         appInfo.Name,
        "is_active":    appInfo.IsActive,
        "description":  appInfo.Description,
        "index_url":    appInfo.IndexUrl,
        "login_url":    appInfo.LoginUrl,
    }
    dbApp, err := models.EditApp(appInfo.ID, app);
    if err != nil {
        return nil, err
    }
    return formatData(dbApp), nil
}

func (appInfo *AppInfo) GetByID() ([]AppInfo, error) {
    data, err := models.GetAppByID(appInfo.ID)

    if err != nil {
        return nil, err
    }
    return formatData(data), err
}

func (appInfo *AppInfo) GetAll(pageNum, pageSize int) ([]AppInfo, error) {
    apps, err := models.GetApps(pageNum, pageSize, appInfo.getParamMaps())
    if err != nil {
        return nil, err
    }
    appsInfo := formatData(apps)
    //jsonAppsInfo, err := json.Marshal(appsInfo)
    return appsInfo, nil
}

func (appInfo *AppInfo) Count() (int, error) {
    return models.GetTotalApp(appInfo.getParamMaps())
}

func (appInfo *AppInfo) getParamMaps() map[string]interface{} {
    maps := make(map[string]interface{})

    if appInfo.Name != "" {
        maps["Name"] = appInfo.Name
    }

    return maps
}

func formatData(obj interface{}) []AppInfo {
    appsInfo := []AppInfo{}
    typ := reflect.ValueOf(obj)
    //val := reflect.ValueOf(obj)
    kd := typ.Kind() //获取kd对应的类别
    if kd == reflect.Slice {
        for _, v := range obj.([]models.App) {
            appsInfo = append(appsInfo, *generateData(&v))
        }
    } else if kd == reflect.Ptr {
        appsInfo = append(appsInfo, *generateData(obj.(*models.App)))
    }
    return appsInfo
}

func generateData(data *models.App) *AppInfo{
    appInfo := AppInfo{
        ID:           data.ID,
        Name:         data.Name,
        IsActive:     data.IsActive,
        Description:  data.Description,
        IndexUrl:     data.IndexUrl,
        LoginUrl:     data.LoginUrl,
        CreatedAt:    data.CreatedAt,
        UpdatedAt:    data.UpdatedAt,
    }
    return &appInfo
}
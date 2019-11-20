package action_service

import (
    "github.com/Unknwon/com"
    "iam/app/models"
    e "iam/pkg/exception"
    "iam/pkg/logging"
    "iam/pkg/utils"
    "net/http"
    "reflect"
    "time"
)

type ActionInfo struct {
    ID              int         `json:"id"`
    Name            string      `json:"name"`
    DisplayName     string      `json:"display_name"`
    IsActive        bool        `json:"is_active"`
    AppID           int         `json:"app_id"`
    CreatedAt       time.Time   `json:"created_at"`
    UpdatedAt       time.Time   `json:"updated_at"`
}

type AddUpdateActionForm struct {
    Name            string      `form:"name" json:"name" valid:"Required;MaxSize(100)"`
    DisplayName     string      `form:"display_name" json:"display_name" valid:"Required;MaxSize(255)"`
    IsActive        int         `form:"is_active" json:"is_active" valid:"Required;Range(0,1)"`
    AppID           int         `form:"app_id" json:"app_id"`
}

func CreateAppAction(appG *utils.Gin) (int, int, *ActionInfo, []string) {
    actionForm := AddUpdateActionForm{}
    httpCode, errCode, errorsMsg := utils.BindAndValid(appG.C, &actionForm)
    appID := com.StrTo(appG.C.Param("id")).MustInt()

    if errCode != e.Success {
        logging.Error(errorsMsg)
        return httpCode, errCode, nil, errorsMsg
    }

    action := models.Action{
        Name:           actionForm.Name,
        IsActive:       actionForm.IsActive == 1,
        DisplayName:    actionForm.DisplayName,
        AppID:          appID,
    }

    dbAction, err := action.AddAppAction()
    if err != nil {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorCreateAppFail), errors)
        return http.StatusInternalServerError, e.ErrorCreateAppFail, nil, errors
    }

    actionInfo := ActionInfo{
        ID:             dbAction.ID,
        Name:           dbAction.Name,
        DisplayName:    dbAction.DisplayName,
        IsActive:       dbAction.IsActive,
        CreatedAt:      dbAction.CreatedAt,
        UpdatedAt:      dbAction.UpdatedAt,
        AppID:          appID,
    }
    return http.StatusOK, e.SuccessCreate, &actionInfo, nil
}

func DeleteAppAction(appG *utils.Gin) (int, int, map[string]interface{}, []string) {
    actionID := com.StrTo(appG.C.Param("action_id")).MustInt()
    action := models.Action{}
    err := action.DeleteAppAction(actionID)
    if err != nil {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.Error), errors)
        return http.StatusInternalServerError, e.Error, nil, errors
    }
    return http.StatusOK, e.SuccessResponse, nil, nil
}

func UpdateAppAction(appG *utils.Gin) (int, int, *ActionInfo, []string) {
    actionForm := AddUpdateActionForm{}
    httpCode, errCode, errorsMsg := utils.BindAndValid(appG.C, &actionForm)
    if errCode != e.Success {
        return httpCode, errCode, nil, errorsMsg
    }

    appID := com.StrTo(appG.C.Param("id")).MustInt()
    actionID := com.StrTo(appG.C.Param("action_id")).MustInt()

    action := models.Action{
        AppID:          appID,
    }

    maps := make(map[string]interface{})
    maps["name"] = actionForm.Name
    maps["display_name"] = actionForm.DisplayName
    maps["is_active"] = actionForm.IsActive

    dbAction, err := action.EditAppAction(actionID, maps)
    if err != nil {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorUpdateAppAction), errors)
        return http.StatusInternalServerError, e.ErrorUpdateAppAction, nil, errors
    }

    actionInfo := ActionInfo{
        ID:             dbAction.ID,
        Name:           dbAction.Name,
        DisplayName:    dbAction.DisplayName,
        IsActive:       dbAction.IsActive,
        CreatedAt:      dbAction.CreatedAt,
        UpdatedAt:      dbAction.UpdatedAt,
        AppID:          appID,
    }
    return http.StatusOK, e.Success, &actionInfo, nil
}

func ListAppActions(appG *utils.Gin) (int, int, map[string]interface{}, []string) {
    appID := com.StrTo(appG.C.Param("id")).MustInt()
    actionInfo := ActionInfo{
        AppID: appID,
    }
    action := models.Action{}
    totalActionActions, err := action.CountAppActions(actionInfo.getParamMaps())
    if err != nil  {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorCountAppAction), errors)
        return http.StatusInternalServerError, e.ErrorCountAppAction, nil, utils.ConvErrorToSlice(err, []string{})
    }

    pageNum := utils.GetPage(appG.C, totalActionActions)
    pageSize := utils.GetPageSize(appG.C)

    actions, err := action.ListAppActions(pageNum, pageSize, actionInfo.getParamMaps())
    if err != nil  {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorListAppActions), errors)
        return http.StatusInternalServerError, e.ErrorListAppActions, nil, utils.ConvErrorToSlice(err, []string{})
    }
    actionsInfo := formatData(actions)

    data := make(map[string]interface{})
    data["items"] = actionsInfo
    data["total"] = totalActionActions
    data["page_number"] = utils.ShowPageNum(pageNum, pageSize, totalActionActions)
    data["page_size"] = pageSize

    return http.StatusOK, e.Success, data, nil
}

func GetAppActionByID(appG *utils.Gin) (int, int, []ActionInfo, []string) {
    appID := com.StrTo(appG.C.Param("id")).MustInt()
    actionID := com.StrTo(appG.C.Param("action_id")).MustInt()
    actionInfo := ActionInfo{
        AppID: appID,
        ID:    actionID,
    }
    action := models.Action{}

    actions, err := action.GetAppActionByID(actionInfo.getParamMaps())
    if err != nil  {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(actionID, e.GetMsg(e.ErrorGetAppAction), errors)
        return http.StatusInternalServerError, e.ErrorGetAppAction, nil, utils.ConvErrorToSlice(err, []string{})
    }
    actionsInfo := formatData(actions)
    return http.StatusOK, e.Success, actionsInfo, nil
}

func (actionInfo *ActionInfo) getParamMaps() map[string]interface{} {
    maps := make(map[string]interface{})
    maps["app_id"] = actionInfo.AppID

    if actionInfo.Name != "" {
        maps["name"] = actionInfo.Name
    }
    if actionInfo.ID != 0 {
        maps["id"] = actionInfo.ID
    }

    return maps
}

// 方法需要重写，每个service里的formatData需要合并为一个方法
func formatData(obj interface{}) []ActionInfo {
    actionsInfo := []ActionInfo{}
    typ := reflect.ValueOf(obj)
    //val := reflect.ValueOf(obj)
    kd := typ.Kind() //获取kd对应的类别
    if kd == reflect.Slice {
        for _, v := range obj.([]models.Action) {
            actionsInfo = append(actionsInfo, *generateData(&v))
        }
    } else if kd == reflect.Ptr {
        actionsInfo = append(actionsInfo, *generateData(obj.(*models.Action)))
    }
    return actionsInfo
}

func generateData(data *models.Action) *ActionInfo{
    appInfo := ActionInfo{
        ID:           data.ID,
        Name:         data.Name,
        DisplayName:  data.DisplayName,
        IsActive:     data.IsActive,
        CreatedAt:    data.CreatedAt,
        UpdatedAt:    data.UpdatedAt,
        AppID:        data.AppID,
    }
    return &appInfo
}

package role_service

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

type RoleInfo struct {
    ID              int         `json:"id"`
    Name            string      `json:"name"`
    Description     string      `json:"description"`
    Permission      string      `json:"permission"`
    AppID           int         `json:"app_id"`
    CreatedAt       time.Time   `json:"created_at"`
    UpdatedAt       time.Time   `json:"updated_at"`
}

type AddUpdateRoleForm struct {
    Name            string      `form:"name" json:"name" valid:"Required;MaxSize(100)"`
    Description     string      `form:"description" json:"description" valid:"Required;MaxSize(255)"`
    Permission      string      `form:"permission" json:"permission" valid:"Required;MaxSize(50)"`
    AppID           int         `form:"app_id" json:"app_id"`
}

func CreateAppRole(appG *utils.Gin) (int, int, *RoleInfo, []string) {
    roleForm := AddUpdateRoleForm{}
    httpCode, errCode, errorsMsg := utils.BindAndValid(appG.C, &roleForm)
    appID := com.StrTo(appG.C.Param("id")).MustInt()

    if errCode != e.Success {
        return httpCode, errCode, nil, errorsMsg
    }

    role := models.Role{
        Name:           roleForm.Name,
        Description:    roleForm.Description,
        Permission:     roleForm.Permission,
        AppID:          appID,
    }

    dbRole, err := role.AddAppRole()
    if err != nil {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorCreateAppFail), errors)
        return http.StatusInternalServerError, e.ErrorCreateAppFail, nil, errors
    }

    roleInfo := RoleInfo{
        ID:             dbRole.ID,
        Name:           dbRole.Name,
        Description:    dbRole.Description,
        Permission:     dbRole.Permission,
        CreatedAt:      dbRole.CreatedAt,
        AppID:          appID,
    }
    return http.StatusOK, e.SuccessCreate, &roleInfo, nil
}

func DeleteAppRole(appG *utils.Gin) (int, int, map[string]interface{}, []string) {
    roleID := com.StrTo(appG.C.Param("role_id")).MustInt()
    role := models.Role{}
    err := role.DeleteAppRole(roleID)
    if err != nil {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.Error), errors)
        return http.StatusInternalServerError, e.Error, nil, errors
    }
    return http.StatusOK, e.SuccessResponse, nil, nil
}

func UpdateAppRole(appG *utils.Gin) (int, int, *RoleInfo, []string) {
    roleForm := AddUpdateRoleForm{}
    httpCode, errCode, errorsMsg := utils.BindAndValid(appG.C, &roleForm)
    if errCode != e.Success {
        return httpCode, errCode, nil, errorsMsg
    }

    appID := com.StrTo(appG.C.Param("id")).MustInt()
    roleID := com.StrTo(appG.C.Param("role_id")).MustInt()

    role := models.Role{
        //Name:           roleForm.Name,
        //Description:    roleForm.Description,
        //Permission:     roleForm.Permission,
        AppID:          appID,
    }

    maps := make(map[string]interface{})
    maps["name"] = roleForm.Name
    maps["description"] = roleForm.Description
    maps["permission"] = roleForm.Permission

    dbRole, err := role.EditAppRole(roleID, maps)
    if err != nil {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorUpdateAppRole), errors)
        return http.StatusInternalServerError, e.ErrorUpdateAppRole, nil, errors
    }

    roleInfo := RoleInfo{
        ID:             dbRole.ID,
        Name:           dbRole.Name,
        Description:    dbRole.Description,
        Permission:     dbRole.Permission,
        CreatedAt:      dbRole.CreatedAt,
        UpdatedAt:      dbRole.UpdatedAt,
        AppID:          appID,
    }
    return http.StatusOK, e.Success, &roleInfo, nil
}

func ListAppRoles(appG *utils.Gin) (int, int, map[string]interface{}, []string) {
    appID := com.StrTo(appG.C.Param("id")).MustInt()
    roleInfo := RoleInfo{
        AppID: appID,
    }
    role := models.Role{}
    totalAppRoles, err := role.CountAppRoles(roleInfo.getParamMaps())
    if err != nil  {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorCountAppRolesFail), errors)
        return http.StatusInternalServerError, e.ErrorCountAppRolesFail, nil, utils.ConvErrorToSlice(err, []string{})
    }

    pageNum := utils.GetPage(appG.C, totalAppRoles)
    pageSize := utils.GetPageSize(appG.C)

    roles, err := role.ListAppRoles(pageNum, pageSize, roleInfo.getParamMaps())
    if err != nil  {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorListAppRolesFail), errors)
        return http.StatusInternalServerError, e.ErrorListAppRolesFail, nil, utils.ConvErrorToSlice(err, []string{})
    }
    rolesInfo := formatData(roles)

    data := make(map[string]interface{})
    data["items"] = rolesInfo
    data["total"] = totalAppRoles
    data["page_number"] = utils.ShowPageNum(pageNum, pageSize, totalAppRoles)
    data["page_size"] = pageSize

    return http.StatusOK, e.Success, data, nil
}

func GetAppRoleByID(appG *utils.Gin) (int, int, []RoleInfo, []string) {
    appID := com.StrTo(appG.C.Param("id")).MustInt()
    roleID := com.StrTo(appG.C.Param("role_id")).MustInt()
    roleInfo := RoleInfo{
        AppID: appID,
        ID:    roleID,
    }
    role := models.Role{}

    roles, err := role.GetAppRoleByID(roleInfo.getParamMaps())
    if err != nil  {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(roleID, e.GetMsg(e.ErrorGetAppRole), errors)
        return http.StatusInternalServerError, e.ErrorGetAppRole, nil, utils.ConvErrorToSlice(err, []string{})
    }
    rolesInfo := formatData(roles)

    return http.StatusOK, e.Success, rolesInfo, nil
}

func (roleInfo *RoleInfo) getParamMaps() map[string]interface{} {
    maps := make(map[string]interface{})

    maps["app_id"] = roleInfo.AppID

    if roleInfo.Name != "" {
        maps["name"] = roleInfo.Name
    }
    if roleInfo.ID != 0 {
        maps["id"] = roleInfo.ID
    }

    return maps
}

// 方法需要重写，每个service里的formatData需要合并为一个方法
func formatData(obj interface{}) []RoleInfo {
    rolesInfo := []RoleInfo{}
    typ := reflect.ValueOf(obj)
    //val := reflect.ValueOf(obj)
    kd := typ.Kind() //获取kd对应的类别
    if kd == reflect.Slice {
        for _, v := range obj.([]models.Role) {
            rolesInfo = append(rolesInfo, *generateData(&v))
        }
    } else if kd == reflect.Ptr {
        rolesInfo = append(rolesInfo, *generateData(obj.(*models.Role)))
    }
    return rolesInfo
}

func generateData(data *models.Role) *RoleInfo{
    appInfo := RoleInfo{
        ID:           data.ID,
        Name:         data.Name,
        Description:  data.Description,
        Permission:   data.Permission,
        CreatedAt:    data.CreatedAt,
        UpdatedAt:    data.UpdatedAt,
        AppID:        data.AppID,
    }
    return &appInfo
}


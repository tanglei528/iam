package user_service

import (
    "github.com/Unknwon/com"
    "github.com/pkg/errors"
    "iam/app/models"
    e "iam/pkg/exception"
    "iam/pkg/logging"
    "iam/pkg/utils"
    "net/http"
    "reflect"
    "time"
)

type UserInfo struct {
    ID              int         `json:"id"`
    Name            string      `json:"name"`
    Email           string      `json:"email"`
    Phone           string      `json:"phone"`
    Address         string      `json:"address"`
    CreatedAt       time.Time   `json:"created_at"`
    UpdatedAt       time.Time   `json:"updated_at"`
}

type CreateOrUpdateUserForm struct {
    Name        string      `form:"name" json:"name" valid:"Required;MaxSize(100)"`
    Password    string      `form:"password" json:"password" valid:"Required;MaxSize(50)"`
    Email       string      `form:"email" json:"email" valid:"Required;Email;MaxSize(100)"`
    Phone       string      `form:"phone" json:"phone" valid:"Required;Mobile"`
    Address     string      `form:"address" json:"address" valid:"Required;MaxSize(255)"`
}

func CreateUser(appG *utils.Gin) (int, int, *UserInfo, []string) {
    // 验证表单属性
    userForm := new(CreateOrUpdateUserForm)
    httpCode, errCode, errorsMsg := utils.BindAndValid(appG.C, userForm)
    if errCode != e.Success {
        logging.Error(errorsMsg)
        return httpCode, errCode, nil, errorsMsg
    }

    // 绑定form属性到model上
    user := models.User{
        Name:       userForm.Name,
        Email:      userForm.Email,
        Phone:      userForm.Phone,
        Address:    userForm.Address,
    }
    // 密码加密存储
    pwd, err := utils.GeneratePassword(userForm.Password)
    if err != nil {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorGeneratePassword))
        return http.StatusInternalServerError, e.ErrorGeneratePassword, nil, errors
    }
    user.Password = pwd

    // 调用创建方法
    dbUser, err := user.AddUser()
    if err != nil {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorCreateUser), errors)
        return http.StatusInternalServerError, e.ErrorCreateUser, nil, errors
    }

    // 把数据库返回的数据绑定到用于页面显示的结构上
    userInfo := UserInfo{
        ID:         dbUser.ID,
        Name:       dbUser.Name,
        Email:      dbUser.Email,
        Phone:      dbUser.Phone,
        Address:    dbUser.Address,
        CreatedAt:  dbUser.CreatedAt,
        UpdatedAt:  dbUser.UpdatedAt,
    }
    return http.StatusOK, e.SuccessCreate, &userInfo, nil
}

func DeleteUser(appG *utils.Gin) (int, int, map[string]interface{}, []string) {
    userID := com.StrTo(appG.C.Param("id")).MustInt()
    user := &models.User{}
    err := user.DeleteUser(userID)
    if err != nil {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorDeleteUser), errors)
        return http.StatusInternalServerError, e.ErrorDeleteUser, nil, errors
    }
    return http.StatusOK, e.SuccessResponse, nil, nil
}

func UpdateUser(appG *utils.Gin) (int, int, *UserInfo, []string) {
    userForm := CreateOrUpdateUserForm{}
    httpCode, errCode, errorsMsg := utils.BindAndValid(appG.C, &userForm)
    if errCode != e.Success {
        return httpCode, errCode, nil, errorsMsg
    }

    userID := com.StrTo(appG.C.Param("id")).MustInt()

    user := models.User{}

    maps := make(map[string]interface{})
    maps["name"] = userForm.Name
    maps["password"] = userForm.Password
    maps["email"] = userForm.Email
    maps["phone"] = userForm.Phone
    maps["address"] = userForm.Address

    dbUser, err := user.EditUser(userID, maps)
    if err != nil {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorUpdateUser), errors)
        return http.StatusInternalServerError, e.ErrorUpdateUser, nil, errors
    }

    userInfo := UserInfo{
        ID:             dbUser.ID,
        Name:           dbUser.Name,
        Email:          dbUser.Email,
        Phone:          dbUser.Phone,
        Address:        dbUser.Address,
        CreatedAt:      dbUser.CreatedAt,
        UpdatedAt:      dbUser.UpdatedAt,
    }
    return http.StatusOK, e.Success, &userInfo, nil
}

func ListUsers(appG *utils.Gin) (int, int, map[string]interface{}, []string) {
    userInfo := UserInfo{}
    user := models.User{}
    totalUsers, err := user.CountUsers(userInfo.getParamMaps())
    if err != nil  {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorCountUser), errors)
        return http.StatusInternalServerError, e.ErrorCountUser, nil, utils.ConvErrorToSlice(err, []string{})
    }

    pageNum := utils.GetPage(appG.C, totalUsers)
    pageSize := utils.GetPageSize(appG.C)

    users, err := user.ListUsers(pageNum, pageSize, userInfo.getParamMaps())
    if err != nil  {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorListUsers), errors)
        return http.StatusInternalServerError, e.ErrorListUsers, nil, utils.ConvErrorToSlice(err, []string{})
    }
    rolesInfo := formatData(users)

    data := make(map[string]interface{})
    data["items"] = rolesInfo
    data["total"] = totalUsers
    data["page_number"] = utils.ShowPageNum(pageNum, pageSize, totalUsers)
    data["page_size"] = pageSize

    return http.StatusOK, e.Success, data, nil
}

// 根据email获取用户信息，用于登录验证
// 这里返回的是models.User，会把密码也返回
func GetUserByEmail(email string) (*models.User, error) {
    userInfo := UserInfo{
        Email:  email,
    }
    user := models.User{}

    users, err := user.ListUsers(0, 1, userInfo.getParamMaps())
    if err != nil  {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(e.GetMsg(e.ErrorListUsers), errors)
        return nil, err
    }
    if len(users) == 0 {
        logging.Error(email, "no record found")
        return nil, errors.New("no record found")
    }
    return &users[0], nil
}


func (userInfo *UserInfo) getParamMaps() map[string]interface{} {
    maps := make(map[string]interface{})

    if userInfo.Name != "" {
        maps["name"] = userInfo.Name
    }

    if userInfo.Email != "" {
        maps["email"] = userInfo.Email
    }

    return maps
}

func GetUserByID(appG *utils.Gin) (int, int, []UserInfo, []string) {
    userID := com.StrTo(appG.C.Param("id")).MustInt()
    user := models.User{}
    users, err := user.GetUserByID(userID)
    if err != nil  {
        errors := utils.ConvErrorToSlice(err, []string{})
        logging.Error(userID, e.GetMsg(e.ErrorGetUser), errors)
        return http.StatusInternalServerError, e.ErrorGetUser, nil, utils.ConvErrorToSlice(err, []string{})
    }
    usersInfo := formatData(users)

    return http.StatusOK, e.Success, usersInfo, nil
}

// 方法需要重写，每个service里的formatData需要合并为一个方法
func formatData(obj interface{}) []UserInfo {
    usersInfo := []UserInfo{}
    typ := reflect.ValueOf(obj)
    //val := reflect.ValueOf(obj)
    kd := typ.Kind() //获取kd对应的类别
    if kd == reflect.Slice {
        for _, v := range obj.([]models.User) {
            usersInfo = append(usersInfo, *generateData(&v))
        }
    } else if kd == reflect.Ptr {
        usersInfo = append(usersInfo, *generateData(obj.(*models.User)))
    }
    return usersInfo
}

func generateData(data *models.User) *UserInfo{
    appInfo := UserInfo{
        ID:             data.ID,
        Name:           data.Name,
        Email:          data.Email,
        Phone:          data.Phone,
        Address:        data.Address,
        CreatedAt:      data.CreatedAt,
        UpdatedAt:      data.UpdatedAt,

    }
    return &appInfo
}

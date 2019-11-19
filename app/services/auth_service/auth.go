package auth_service

import (
    "iam/app/services/user_service"
    e "iam/pkg/exception"
    "iam/pkg/logging"
    "iam/pkg/utils"
    "net/http"
)

type AuthToken struct {
    Token       string      `form:"token" json:"token" valid:"Required"`
}

type AuthForm struct {
    Email        string     `form:"email" json:"email" valid:"Required;Email;MaxSize(50)"`
    Password     string     `form:"password" json:"password" valid:"Required;MaxSize(50)"`
}

func Login(appG *utils.Gin) (int, int, *AuthToken, []string) {
    // 验证表单属性
    authForm := new(AuthForm)
    httpCode, errCode, errorsMsg := utils.BindAndValid(appG.C, authForm)
    if errCode != e.Success {
        logging.Error(errorsMsg)
        return httpCode, errCode, nil, errorsMsg
    }

    dbUser, err := user_service.GetUserByEmail(authForm.Email)
    if err != nil {
        return http.StatusInternalServerError, e.ErrorListUsers, nil, utils.ConvErrorToSlice(err, []string{})
    }

    err = utils.CheckPassword(dbUser.Password, authForm.Password)
    if err != nil {
        logging.Error(err)
        return http.StatusUnauthorized, e.ErrorCheckEmailAndPwd, nil, nil
    }

    token, err := GenerateToken(authForm.Email, authForm.Password)
    if err != nil {
        logging.Error(e.GetMsg(e.ErrorAuthToken), err)
        return http.StatusInternalServerError, e.ErrorAuthToken, nil, nil
    }
    authToken := &AuthToken{
        Token: token,
    }
    return http.StatusOK, e.Success, authToken, nil
}

// 验证token
func ValidateToken(appG *utils.Gin) (int, int, *AuthToken, []string) {
    // 验证表单属性
    authToken := new(AuthToken)
    httpCode, errCode, errorsMsg := utils.BindAndValid(appG.C, authToken)
    if errCode != e.Success {
        logging.Error(errorsMsg)
        return httpCode, errCode, nil, errorsMsg
    }

    code := utils.CheckToken(authToken.Token)
    if code != e.Success {
        return http.StatusUnauthorized, code, nil, nil
    }
    return http.StatusOK, e.Success, authToken, nil
}

func GenerateToken(username, password string) (string, error){
    token, err := utils.GenerateToken(username, password)
    return token, err
}

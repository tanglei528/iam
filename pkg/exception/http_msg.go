package exception

var MsgFlags = map[int]string {
    Success : "操作成功",
    SuccessCreate:  "创建成功",
    SuccessResponse: "请求成功",
    Error : "操作失败",
    InvalidParams : "请求参数错误",

    ErrorGetAppsFail : "获取多个应用失败",
    ErrorGetAppFail : "获取单个应用失败",
    ErrorCountAppFail  : "获取应用数量失败",
    ErrorCreateAppFail  : "创建应用失败",
    ErrorCountAppRolesFail: "获取应用角色数量失败",
    ErrorListAppRolesFail:  "获取应用角色列表失败",
    ErrorGetAppRole:  "获取应用单个角色失败",
    ErrorUpdateAppRole:  "修改应用角色失败",

    ErrorCreateUser:    "创建用户失败",
    ErrorDeleteUser:    "删除用户失败",
    ErrorUpdateUser:    "更新用户失败",
    ErrorListUsers:     "获取用户列表失败",
    ErrorGetUser:       "获取单个用户失败",
    ErrorCountUser:     "获取用户数量失败",

    ErrorCreateAppAction:  "创建行为失败",
    ErrorDeleteAppAction:  "删除行为失败",
    ErrorUpdateAppAction:  "跟新行为失败",
    ErrorListAppActions:   "获取行为列表失败",
    ErrorGetAppAction:     "获取单个行为失败",
    ErrorCountAppAction:   "获取行为数量失败",

    ErrorAuthCheckTokenFail : "Token鉴权失败",
    ErrorAuthCheckTokenTimeout : "Token过期",
    ErrorAuthToken : "Token生成失败",
    ErrorAuth : "Token错误",
    ErrorAuthTokenProvide:  "请求头没有token属性",
    ErrorGeneratePassword:  "密码加密失败",
    ErrorCheckEmailAndPwd: "Email密码验证失败",

    ErrorConvIntToBool : "整型转布尔型错误",
    ErrorAppHeader     : "Header必须包含APP_ID和APP_KEY",
    ErrorAppIDKey      : "应用ID和KEY不匹配",


    ErrorNoRedisKey:  "Redis中key不存在",
}

func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[Error]
}

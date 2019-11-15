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

    ErrorAuthCheckTokenFail : "Token鉴权失败",
    ErrorAuthCheckTokenTimeout : "Token已超时",
    ErrorAuthToken : "Token生成失败",
    ErrorAuth : "Token错误",

    ErrorConvIntToBool : "整型转布尔型错误",
}

func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[Error]
}
